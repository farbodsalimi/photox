#!make

# Settings
.DEFAULT_GOAL := help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
GOGET=$(GOCMD) get
BINARY_DIR=./bin
BINARY_NAME=$(BINARY_DIR)/photox
BINARY_UNIX=$(BINARY_NAME)_unix
PHOTOX_CMD=cmd/photox/main.go


all: test build ## Build and test the binary

.PHONY: build
build: ## Build the binary
	$(GOBUILD) -o $(BINARY_NAME) -v $(PHOTOX_CMD)

.PHONY: test
test: ## Test all the test files recursively
	$(GOTEST) -v ./test/... -coverpkg=./...

test-cover: ## Test and generate the coverage report
	$(GOTEST)  -coverprofile=coverage.out ./test/... -coverpkg=./... && $(GOTOOL) cover -func=coverage.out

test-show-cover:
	$(GOTOOL) cover -html=coverage.out

clean: ## Clean the binaries
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

.PHONY: install
install: ## Install the binary
	$(GOINSTALL) $(PHOTOX_CMD)

.PHONY: run
run: ## Run the binary
	$(GOBUILD) -o $(BINARY_NAME) -v $(PHOTOX_CMD)
	./$(BINARY_NAME)

.PHONY: build-linux
build-linux: ## Cross compilation
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v $(PHOTOX_CMD)

.PHONY: build-scratch
build-scratch:
	CGO_ENABLED=0 GOOS=linux $(GOBUILD) -a -installsuffix cgo -o $(BINARY_UNIX) $(PHOTOX_CMD)

.PHONY: --help
--help: ##
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: help
help: --help
