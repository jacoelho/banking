BIN = $(CURDIR)/bin

export GOBIN=$(BIN)

.PHONY: build
build:
	go install -v ./...

.PHONY: test
test:
	go test -race -v ./...

.PHONY: vendor
vendor:
	go mod tidy && go mod vendor && go mod verify

.PHONY: ci-tidy
ci-tidy:
	go mod tidy
	git status --porcelain go.mod go.sum || { echo "Please run 'go mod tidy'."; exit 1; }
