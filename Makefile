# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R
DATE  = $(shell date +%Y%m%d%H%M%S)
export GOBIN = $(CURDIR)/bin

default: build

.PHONY: build
build: test
	go install -v ./...

$(GOBIN)/banking-registry:
	cd registry; go install ./cmd/banking-registry

$(GOBIN)/tsv-to-yaml:
	cd registry; go install ./cmd/tsv-to-yaml

.PHONY: generate
generate: $(GOBIN)/banking-registry
	go generate -v ./...

.PHONY: update-registry
update-registry: $(GOBIN)/tsv-to-yaml
	@if [ -z "$(REGISTRY)" ]; then \
		echo "Error: REGISTRY variable is required. Usage: make update-registry REGISTRY=filename"; \
		exit 1; \
	fi
	$(GOBIN)/tsv-to-yaml -input $(REGISTRY) -output docs/registry.yml
	$(MAKE) generate fmt wasm

.PHONY: test
test:
	go test -race -shuffle=on -v ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem ./... | tee benchmarks/$(DATE).bench

.PHONY: fmt
fmt:
	gofmt -s -w $$(go list -f '{{.Dir}}' ./... | grep -v vendor)

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

