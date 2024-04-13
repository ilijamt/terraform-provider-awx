default: generate build

VERSION=24.1.0

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

.PHONY: test
test:
	go test ./internal/... -count=1 -parallel=4 -cover -coverprofile=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

.PHONY: testacc
testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
