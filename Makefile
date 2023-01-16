GOTEST             =gotest
ifeq (, $(shell which gotest))
    GOTEST=go test
endif

.PHONY: lint test lint-fix

test:
	$(GOTEST) ./...

lint:
	golangci-lint run --timeout 10m0s --allow-parallel-runners $(param) ./...

lint-fix: param=--fix
lint-fix: lint

tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1