OUT_BIN_DIR=./bin
SRC_DIR= .
TEST_OPTS=--race
PROJECT_NAME="invastion simulator"
$(info BUILD_TAGS=$(BUILD_TAGS))

.PHONY: build
build: install build-cli

.PHONY: install
install:
	@go mod download

.PHONY: build-cli
build-cli:
	@echo "Building $(PROJECT_NAME) cli..."
	@go build -tags "$(BUILD_TAGS)" -mod=mod -o "$(OUT_BIN_DIR)/" $(SRC_DIR)/cmd/cli

.PHONY: generate
generate: install-tools
	go generate ./...

.PHONY: test
test: generate
	go test $(TEST_OPTS) ./...

.PHONY: install-tools
install-tools:
	@go mod download -modfile=tools.mod

.PHONY: fmt
fmt: install-tools
	# Fixup modules
	go mod tidy
	# Format the Go sources:
	gofmt -s -w .

.PHONY: lint
lint: install-tools
	@golangci-lint run
