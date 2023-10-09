default: build

.PHONY: generate-awx
generate-awx:
	rm -f internal/awx_21_8_0/gen_*.go
	rm -rf cmd/provider_21_8_0/docs/*
	go run ./tools/generator/cmd/generator/main.go template resources/config.json resources/api/21.8.0 internal/awx_21_8_0
	goimports -w internal/awx_21_8_0/*.go
	gofmt -s -w internal/awx_21_8_0/*.go

.PHONY: generate-tfplugindocs
generate-tfplugindocs:
	mkdir -p docs/21.8.0
	tfplugindocs generate --examples-dir examples --provider-name awx --provider-dir ./cmd/provider_21_8_0

.PHONY: generate
generate: generate-awx generate-tfplugindocs

.PHONY: build
build:
	go build -trimpath -o ./build/terraform-provider-awx -ldflags "-s -w" ./cmd/provider_21_8_0

test:
	go test ./internal/... -count=1 -parallel=4 -cover -coverprofile=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
