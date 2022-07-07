# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R
GOBIN = $(shell go env GOPATH)/bin
DATE  = $(shell date +%Y%m%d%H%M%S)

default: build

.PHONY: build
build: test
	go install -v ./...

.PHONY: generate
generate:
	go generate -v ./...

.PHONY: test
test:
	go test -race -shuffle=on -v ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem ./... | tee benchmarks/$(DATE).bench

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vendor
vendor:
	go mod tidy && go mod vendor && go mod verify

.PHONY: ci-tidy
ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }

$(GOBIN)/staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

$(GOBIN)/gcassert:
	go install github.com/jordanlewis/gcassert/cmd/gcassert@latest

.PHONY: staticcheck
staticcheck: $(GOBIN)/staticcheck
	staticcheck ./...

.PHONY: gcassert
gcassert: $(GOBIN)/gcassert
	gcassert ./...

