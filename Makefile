# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R
GOLINT_VERSION = v1.24.0
DATE = $(shell date +%Y%m%d%H%M%S)

export GOBIN=$(CURDIR)/bin

default: build

.PHONY: build
build: test
	go install -v ./...

.PHONY: generate
generate:
	go generate -v ./...

.PHONY: test
test:
	go test -race -v ./...

.PHONY: bench
bench:
	go test -bench=. -benchmem ./... | tee benchmarks/$(DATE).bench

.PHONY: validation
validation:
	go test -race -v -tags validation ./...

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

.PHONY: lint
lint:
	docker run -t --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:$(GOLINT_VERSION) golangci-lint run
	docker run -t --rm -v $(CURDIR):/app -w /app/registry golangci/golangci-lint:$(GOLINT_VERSION) golangci-lint run

