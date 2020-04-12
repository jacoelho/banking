# disable default rules
.SUFFIXES:
MAKEFLAGS+=-r -R

export GOBIN=$(CURDIR)/bin

default: build

.PHONY: build
build: generate test
	go install -v ./...

.PHONY: generate
generate:
	go generate -v ./...

.PHONY: test
test:
	go test -race -v ./...

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

.PHONY: tools
tools:
	go get -tags tools ./...
