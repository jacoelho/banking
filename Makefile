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
	go mod tidy && go mod vendor

ci-tidy:
	go mod tidy
	go list all > /dev/null
	git diff --exit-code --quiet || (echo "Please run 'go mod tidy' to clean up the 'go.mod' and 'go.sum' files."; false)
