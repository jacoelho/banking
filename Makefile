# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R
DATE  = $(shell date +%Y%m%d%H%M%S)
export GOBIN = $(CURDIR)/bin
SOURCE_ENCODING ?= auto

default: build

.PHONY: build
build: test
	go install -v ./...

.PHONY: generate
generate:
	@if [ -z "$(SOURCE_REGISTRY)" ]; then \
		echo "Error: SOURCE_REGISTRY variable is required. Usage: make generate SOURCE_REGISTRY=filename"; \
		exit 1; \
	fi
	go run ./cmd/banking-registry -registry-file "$(SOURCE_REGISTRY)" -encoding "$(SOURCE_ENCODING)" -dst-directory iban

.PHONY: update-registry
update-registry: generate fmt wasm

.PHONY: test
test:
	go test -race -shuffle=on -v ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem ./... | tee benchmarks/$(DATE).bench

.PHONY: fmt
fmt:
	gofmt -s -w $$(go list -f '{{.Dir}}' ./...)

.PHONY: check-registry-tools
check-registry-tools:
	go test ./internal/... ./cmd/banking-registry

.PHONY: vendor
vendor:
	go mod tidy && go mod vendor && go mod verify

.PHONY: ci-tidy
ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }

$(GOBIN)/staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: staticcheck
staticcheck: $(GOBIN)/staticcheck
	$(GOBIN)/staticcheck ./...

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o docs/iban.wasm ./cmd/wasmiban
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js docs/
