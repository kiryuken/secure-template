.PHONY: help build run-api run-worker run-cli test lint migrate-up migrate-down docker-build docker-up docker-down clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build all binaries
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/api ./cmd/api
	@go build -o bin/worker ./cmd/worker
	@go build -o bin/cli ./cmd/cli
	@echo "Build complete!"

run-api: ## Run API server
	@go run ./cmd/api/main.go

run-worker: ## Run background worker
	@go run ./cmd/worker/main.go

run-cli: ## Run CLI
	@go run ./cmd/cli/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

lint: ## Run linters and security scans
	@echo "Running linters..."
	@golangci-lint run ./...
	@gosec ./...
	@staticcheck ./...
	@govulncheck ./...

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	@migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/max_secure?sslmode=disable" up

migrate-down: ## Rollback last migration
	@echo "Rolling back migration..."
	@migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/max_secure?sslmode=disable" down 1

migrate-create: ## Create new migration (use: make migrate-create name=migration_name)
	@migrate create -ext sql -dir migrations -seq $(name)

docker-build: ## Build Docker image
	@docker build -t max-secure:latest .

docker-up: ## Start all services with Docker Compose
	@docker-compose up -d

docker-down: ## Stop all services
	@docker-compose down

docker-logs: ## Show Docker logs
	@docker-compose logs -f

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "Clean complete!"

install-tools: ## Install development tools
	@echo "Installing tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Tools installed!"

swag: ## Generate Swagger documentation
	@echo "Generating Swagger docs..."
	@swag init -g cmd/api/main.go -o docs/api

mod-tidy: ## Tidy go modules
	@go mod tidy

mod-download: ## Download go modules
	@go mod download

mod-vendor: ## Vendor dependencies
	@go mod vendor
