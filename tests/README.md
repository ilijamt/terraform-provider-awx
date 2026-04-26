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

## Notes

- VCR tests live behind the `integration` build tag so plain `make test`
  never invokes them. Use `make test-integration` (or pass
  `-tags=integration` manually).
- `Authorization`, `Cookie`, `Set-Cookie`, and `X-Csrftoken` headers are
  redacted before save.
- AWX assigns sequential IDs. Record from a clean instance to keep
  cassettes deterministic.
