default: build

VERSION=23.8.1

.PHONY: generate-config
generate-config:
	node tools/config-merge.js $(shell pwd)/resources/config $(shell pwd)/resources/api/$(VERSION)

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
	go build -cover -o ./build/terraform-provider-awx ./cmd/provider

.PHONY: build
build:
	go build -trimpath -o ./build/terraform-provider-awx -ldflags "-s -w" ./cmd/provider

test:
	go test ./internal/... -count=1 -parallel=4 -cover -coverprofile=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
