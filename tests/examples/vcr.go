// Package examples runs end-to-end Terraform configs from examples/ against
// recorded HTTP cassettes. The provider's HTTP client is injectable via
// provider.NewFuncProvider; tests here drop in a go-vcr-wrapped client so
// no live AWX is required for replay.
package examples

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"gopkg.in/dnaeon/go-vcr.v4/pkg/cassette"
	"gopkg.in/dnaeon/go-vcr.v4/pkg/recorder"
)

// StableHost is the placeholder hostname used in cassettes. Tests configure
// the provider with this URL; in record mode the wrapper transport rewrites
// outgoing requests to TOWER_HOST/AWX_HOST before they hit the wire, but the
// cassette only ever sees StableHost. This keeps cassettes portable.
const StableHost = "http://awx.local"

// FakeToken is the token written into provider config during replay. Cassettes
// are matched on method/URL/body (not auth headers), so any non-empty value
// works.
const FakeToken = "replay-token"

// RecordEnvVar toggles record mode. Any non-empty value enables recording.
const RecordEnvVar = "AWX_VCR_RECORD"

// IsRecording reports whether the test should record new interactions.
func IsRecording() bool {
	return os.Getenv(RecordEnvVar) != ""
}

func realHost(t *testing.T) string {
	t.Helper()
	for _, k := range []string{"TOWER_HOST", "AWX_HOST"} {
		if v := os.Getenv(k); v != "" {
			return strings.TrimRight(v, "/")
		}
	}
	t.Fatalf("recording mode requires TOWER_HOST or AWX_HOST to be set")
	return ""
}

// NewVCRClient returns an *http.Client wired to a cassette named
// tests/examples/testdata/cassettes/<cassetteName>.yaml. In replay mode
// (default) the cassette must exist. In record mode (AWX_VCR_RECORD=1)
// requests are forwarded to the real AWX and the cassette is overwritten on
// Stop.
func NewVCRClient(t *testing.T, cassetteName string) *http.Client {
	t.Helper()

	cassettePath := cassettePath(t, cassetteName)

	mode := recorder.ModeReplayOnly
	if IsRecording() {
		mode = recorder.ModeRecordOnly
	}

	opts := []recorder.Option{
		recorder.WithSkipRequestLatency(false),
		recorder.WithMode(mode),
		recorder.WithMatcher(matcher),
		recorder.WithHook(redactHook, recorder.AfterCaptureHook),
	}

	if mode == recorder.ModeRecordOnly {
		opts = append(opts, recorder.WithRealTransport(newRewriteTransport(realHost(t))))
	}

	rec, err := recorder.New(cassettePath, opts...)
	if err != nil {
		t.Fatalf("vcr: failed to start recorder for %s: %v", cassettePath, err)
	}
	t.Cleanup(func() {
		if err := rec.Stop(); err != nil {
			t.Errorf("vcr: failed to stop recorder: %v", err)
		}
	})

	return rec.GetDefaultClient()
}

// cassettePath resolves the cassette location relative to this source file
// so tests work regardless of the package they run from.
func cassettePath(t *testing.T, name string) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("vcr: cannot resolve caller for cassette path")
	}
	pkgDir := filepath.Dir(thisFile)
	return filepath.Join(pkgDir, "testdata", "cassettes", name)
}

// FixturePath resolves the path to a fixture .tf under
// tests/examples/testdata/<rel>.
func FixturePath(t *testing.T, rel string) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("cannot resolve caller for fixture path")
	}
	return filepath.Join(filepath.Dir(thisFile), "testdata", rel)
}

// ReadFixture reads and returns the contents of a fixture .tf file.
func ReadFixture(t *testing.T, rel string) string {
	t.Helper()
	b, err := os.ReadFile(FixturePath(t, rel))
	if err != nil {
		t.Fatalf("read fixture %s: %v", rel, err)
	}
	return string(b)
}

// LoadBootstrapToken reads the token file produced by tests/bootstrap.
func LoadBootstrapToken(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("cannot resolve caller for bootstrap token path")
	}
	// tests/examples/vcr.go -> tests/.bootstrap-token
	path := filepath.Join(filepath.Dir(thisFile), "..", ".bootstrap-token")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("recording mode requires %s — run `make bootstrap-awx` first: %v", path, err)
	}
	return strings.TrimSpace(string(b))
}

// matcher matches by method + URL path/query + body. The default matcher
// also checks Host, headers, RequestURI, Form, etc., which is too strict
// (User-Agent changes with provider version, Authorization carries the
// token, RequestURI is empty for client requests).
func matcher(r *http.Request, i cassette.Request) bool {
	if r.Method != i.Method {
		return false
	}
	iURL, err := url.Parse(i.URL)
	if err != nil {
		return false
	}
	if r.URL.Path != iURL.Path {
		return false
	}
	if r.URL.RawQuery != iURL.RawQuery {
		return false
	}
	return bodyMatches(r, i.Body)
}

func bodyMatches(r *http.Request, want string) bool {
	if r.Body == nil || r.Body == http.NoBody {
		return want == ""
	}
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	r.Body = io.NopCloser(bytes.NewBuffer(buf))
	return string(buf) == want
}

// redactHook strips secrets from interactions before they reach the cassette
// in-memory store.
func redactHook(i *cassette.Interaction) error {
	scrub := []string{"Authorization", "Cookie", "Set-Cookie", "X-Csrftoken"}
	for _, h := range scrub {
		if _, ok := i.Request.Headers[h]; ok {
			i.Request.Headers[h] = []string{"REDACTED"}
		}
		if _, ok := i.Response.Headers[h]; ok {
			i.Response.Headers[h] = []string{"REDACTED"}
		}
	}
	return nil
}

// rewriteTransport rewrites outgoing requests from StableHost to the real
// AWX host. Used only in record mode.
type rewriteTransport struct {
	target *url.URL
	base   http.RoundTripper
}

func newRewriteTransport(realHostURL string) http.RoundTripper {
	u, err := url.Parse(realHostURL)
	if err != nil {
		return http.DefaultTransport
	}
	return &rewriteTransport{target: u, base: http.DefaultTransport}
}

func (t *rewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	stable, _ := url.Parse(StableHost)
	if req.URL.Host == stable.Host {
		req.URL.Scheme = t.target.Scheme
		req.URL.Host = t.target.Host
		req.Host = t.target.Host
	}
	return t.base.RoundTrip(req)
}
