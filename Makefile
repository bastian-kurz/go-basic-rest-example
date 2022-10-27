.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: fmt
fmt: dep ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: dep ## Run go vet against code.
	go vet ./...

.PHONY: dep
dep: ## Get the dependencies.
	go get -v -d ./...

.PHONY: test-e2e
test-e2e: fmt vet  ## Run unit-e2e tests.
	go test $(go list ./... | grep -v /test/) -coverprofile cover.out

.PHONY: test
test: fmt vet ## Run unit tests.
	go test ./... -coverprofile cover-unittest.out

.PHONY: build
build: fmt vet ## Build CLI-Tool binary.
	go build -o bin/server cmd/server/main.go && chmod +x bin/server
