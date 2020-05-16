
default: fmt deps build

fmt: ## Format code
	@go fmt ./...

## fetch dependencies
deps:
	@echo "=> Running go mod tidy..."
	@go mod tidy
	@echo "=> Running go mod vendor..."
	@GO111MODULE=on go mod vendor

# build loganalyzer
build: deps
	@echo "==> Building ..."
	@go build -o bin/log-analyzer

