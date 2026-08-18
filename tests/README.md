# VCR-backed Integration Tests

End-to-end integration tests that exercise example Terraform configs
against recorded HTTP cassettes. CI and day-to-day runs need no live AWX;
only re-recording does. All test scaffolding is self-contained under
`tests/`.

## Layout

```
tests/
  README.md
  bootstrap/main.tf           # creates a test token on a fresh AWX
  .bootstrap-token            # token written by bootstrap (gitignored)
  examples/
    vcr.go                    # *http.Client wrapped in go-vcr (helper)
    <name>_test.go            # one Go test file per example (TestIntegration_*)
    testdata/
      <name>/main.tf          # trimmed Terraform fixture
      cassettes/<name>.yaml   # committed recordings
```

## Replay (default — no AWX needed)

```sh
make test-integration
```

Replays cassettes under `tests/examples/testdata/cassettes/`. Fails loudly
on missing cassettes or unmatched interactions.

## Re-record against a local AWX

The local AWX is expected at `awx.local`. Tests always configure the
provider with `http://awx.local`; in record mode the VCR transport
forwards on-the-wire requests to whatever `TOWER_HOST`/`AWX_HOST` you set,
but cassettes only ever contain `awx.local` so they stay portable.

1. Spin up a fresh, disposable AWX. Set:
   ```sh
   export TOWER_HOST=http://awx.local
   export TOWER_USERNAME=admin
   export TOWER_PASSWORD=admin
   ```

2. Bootstrap a token:
   ```sh
   make bootstrap-awx
   ```
   Builds the provider, generates `.terraformrc` (dev_overrides → `build/`)
   in the repo root, then runs `tests/bootstrap` to create a personal access
   token in AWX. Writes the token to `tests/.bootstrap-token`.

4. Re-record cassettes:
   ```sh
   make test-integration-record
   ```
   Overwrites `tests/examples/testdata/cassettes/*.yaml`.

5. Inspect the diff and commit cassettes.

## Adding a new integration test

1. Drop a trimmed `main.tf` (no `terraform { required_providers }`,
   no `provider "awx" {}` block) under
   `tests/examples/testdata/<name>/main.tf`.
2. Add `tests/examples/<name>_test.go` with a `TestIntegration_<Name>`
   function mirroring `inventory_test.go`.
3. Re-record (`make test-integration-record`) to generate the cassette.

## Running against OpenTofu

```sh
make test-integration TF=tofu
```

`TF=tofu` sets `TF_ACC_TERRAFORM_PATH` and `TF_ACC_PROVIDER_HOST` together.
Both matter. The harness registers its reattach providers under the legacy `-`
namespace, and OpenTofu accepts that namespace only beneath
`registry.opentofu.org`. Point `TF_ACC_TERRAFORM_PATH` at `tofu` on its own and
`init` fails with `Invalid provider namespace`. The same cassettes replay under
either CLI.

CI runs the suite as one job per CLI through `make test-integration-cover`,
which writes coverage into `COVERDIR` for the merge step. The same target works
locally:

```sh
make test-integration-cover TF=tofu COVERDIR=build/covdata-tofu
```

## Notes

- VCR tests live behind the `integration` build tag so plain `make test`
  never invokes them. Use `make test-integration` (or pass
  `-tags=integration` manually).
- Replay skips the latency recorded for each interaction, so a run is not
  as slow as the original AWX round-trips were. Set `AWX_VCR_LATENCY=1` to
  replay in real time instead.
- `Authorization`, `Cookie`, `Set-Cookie`, and `X-Csrftoken` headers are
  redacted before save.
- AWX assigns sequential IDs. Record from a clean instance to keep
  cassettes deterministic.
