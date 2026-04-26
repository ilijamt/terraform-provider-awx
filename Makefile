default: generate build

VERSION=24.6.1

# Local AWX endpoint and credentials used by bootstrap-awx and
# test-integration-record. Override on the command line if your local AWX
# differs (e.g. `make bootstrap-awx TOWER_PASSWORD=...`).
TOWER_HOST ?= http://awx.local
TOWER_USERNAME ?= admin
TOWER_PASSWORD ?= admin

.PHONY: generate-config
generate-config:
	node tools/config-merge.js $(shell pwd)/resources/config $(shell pwd)/resources/api/$(VERSION)

.PHONY: download-api
download-api:
	mkdir -p resources/api/$(VERSION)/config resources/api/$(VERSION)/gen-data
	go run ./tools/generator/cmd/generator/main.go fetch-api-resources resources/api/$(VERSION) \
		--host $(TOWER_HOST) --password $(TOWER_PASSWORD) --username $(TOWER_USERNAME) --insecure-skip-verify

.PHONY: generate-configs
generate-configs: resources/api/*
	@for file in $^ ; do \
		node tools/config-merge.js $(shell pwd)/resources/config $(shell pwd)/$${file} ; \
	done

.PHONY: generate-awx
generate-awx: generate-config
	rm -f internal/awx/gen_*.go
	rm -rf cmd/provider/docs/*
	go run ./tools/generator/cmd/generator/main.go template resources/api/$(VERSION) internal/awx
	goimports -w internal/awx/*.go
	gofmt -s -w internal/awx/*.go

.PHONY: generate-tfplugindocs
generate-tfplugindocs:
	rm -rf docs
	mkdir -p cmd/provider/docs
	tfplugindocs generate --examples-dir examples --provider-name awx --provider-dir ./cmd/provider
	mv cmd/provider/docs .

.PHONY: generate
generate: generate-awx generate-tfplugindocs

.PHONY: build-cover
build-cover:
	go build -cover -trimpath -o ./build/terraform-provider-awx ./cmd/provider

.PHONY: build
build:
	go build -trimpath -o ./build/terraform-provider-awx -ldflags "-s -w" ./cmd/provider

.PHONY: build-debug
build-debug:
	go build -cover -covermode=atomic -trimpath -o ./build/terraform-provider-awx ./cmd/provider

.PHONY: test
test:
	go test ./internal/... -count=1 -parallel=4 -cover -coverprofile=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

.PHONY: test-all
test-all:
	rm -rf build/covdata-internal build/covdata-tests build/covdata-merged
	mkdir -p build/covdata-internal build/covdata-tests build/covdata-merged
	go test ./internal/... -count=1 -parallel=4 -cover -coverpkg=./internal/... \
		-args -test.gocoverdir=$(shell pwd)/build/covdata-internal
	TF_ACC=1 go test -tags=integration ./tests/... -run '^TestIntegration_' -count=1 \
		-cover -coverpkg=./internal/... \
		-args -test.gocoverdir=$(shell pwd)/build/covdata-tests
	go tool covdata merge \
		-i=build/covdata-internal,build/covdata-tests \
		-o=build/covdata-merged
	go tool covdata textfmt -i=build/covdata-merged -o=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

.PHONY: testacc
testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...

# Generate .terraformrc with dev_overrides pointing at the local provider
# binary. Path must be absolute, so we derive it from $(shell pwd).
.PHONY: terraformrc
terraformrc:
	@printf 'provider_installation {\n  dev_overrides {\n    "ilijamt/awx" = "%s/build/"\n  }\n\n  direct {}\n}\n' "$(shell pwd)" > .terraformrc
	@echo "wrote $(shell pwd)/.terraformrc"

# Bootstrap a local AWX for VCR recording. Builds the provider, generates a
# dev_overrides .terraformrc, and runs tests/bootstrap to create a token.
# Requires TOWER_USERNAME/TOWER_PASSWORD env vars (TOWER_HOST defaults above).
.PHONY: bootstrap-awx
bootstrap-awx: build terraformrc
	cd tests/bootstrap && \
		TF_CLI_CONFIG_FILE=$(shell pwd)/.terraformrc terraform init && \
		TF_ACC=1 \
		TF_CLI_CONFIG_FILE=$(shell pwd)/.terraformrc \
		TOWER_HOST=$(TOWER_HOST) \
		TOWER_USERNAME=$(TOWER_USERNAME) \
		TOWER_PASSWORD=$(TOWER_PASSWORD) \
		terraform apply -auto-approve

# Replay VCR cassettes — no AWX needed.
.PHONY: test-integration
test-integration:
	TF_ACC=1 go test -tags=integration ./tests/examples/... -run '^TestIntegration_' -count=1 -v

# Re-record VCR cassettes against the local AWX. Requires bootstrap-awx first.
.PHONY: test-integration-record
test-integration-record:
	TF_ACC=1 AWX_VCR_RECORD=1 TOWER_HOST=$(TOWER_HOST) go test -tags=integration ./tests/examples/... -run '^TestIntegration_' -count=1 -v
