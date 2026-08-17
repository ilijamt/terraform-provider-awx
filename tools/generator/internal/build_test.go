package internal

import (
	"testing"

	"github.com/Masterminds/semver/v3"
)

func TestBuildConfigVersionGetBuildVersion(t *testing.T) {
	for _, tc := range []struct {
		name    string
		version string
		build   uint8
		want    string
		wantErr bool
	}{
		{name: "first build of a patch release", version: "24.6.1", build: 0, want: "24.6.100"},
		{name: "current versions.yaml state", version: "24.6.1", build: 3, want: "24.6.103"},
		{name: "highest build in range", version: "24.6.1", build: 99, want: "24.6.199"},
		{name: "second patch release", version: "24.6.2", build: 0, want: "24.6.200"},
		{name: "zero patch degenerates", version: "24.2.0", build: 3, want: "24.2.3"},
		{name: "build collides with next patch", version: "24.6.1", build: 100, wantErr: true},
		{name: "invalid version", version: "not-a-version", build: 0, wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := &BuildConfigVersion{Version: tc.version, Build: tc.build}
			got, err := c.GetBuildVersion()
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got %q", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}

// The registry cannot unpublish, so new tags must sort above published ones.
func TestBuildConfigVersionOrderingAgainstPublishedTags(t *testing.T) {
	next, err := (&BuildConfigVersion{Version: "24.6.1", Build: 3}).GetBuildVersion()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := semver.MustParse(next)

	for _, published := range []string{"23.8.1", "24.2.0-3", "24.6.1-0", "24.6.1-2"} {
		if !got.GreaterThan(semver.MustParse(published)) {
			t.Errorf("%s must sort above already-published %s", next, published)
		}
	}

	for _, constraint := range []string{"~> 24.6", ">= 24.0.0"} {
		c, err := semver.NewConstraint(constraint)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !c.Check(got) {
			t.Errorf("%s must satisfy %q", next, constraint)
		}
	}
}
