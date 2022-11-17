default: build

.PHONY: generate-awx
generate-awx:
	rm -f internal/awx/gen_*.go resources/docs/*.md
	go run ./tools/generator/cmd/generator/main.go template resources/config.json resources/api/21.8.0.json
	goimports -w internal/awx/*.go
	gofmt -s -w internal/awx/*.go

.PHONY: generate-tfplugindocs
generate-tfplugindocs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

.PHONY: generate
generate: generate-awx generate-tfplugindocs

.PHONY: build
build:
	go build -trimpath -o ./build/terraform-provider-awx -ldflags "-s -w" .

test:
	go test ./internal/... -count=1 -parallel=4 -cover -coverprofile=build/coverage.out
	go tool cover -html=build/coverage.out -o build/coverage.html

testacc:
	TF_ACC=1 go test -count=1 -parallel=4 -timeout 10m -v ./...
