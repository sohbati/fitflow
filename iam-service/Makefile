.PHONY: help build run test clean swag docker-up docker-down

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building iam-service..."
	@go build -o iam-service main.go
	@echo "Build completed!"

run: ## Run the application
	@echo "Starting iam-service..."
	@go run main.go

run-cmd: ## Run the application from cmd directory
	@echo "Starting iam-service from cmd..."
	@go run cmd/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -f iam-service
	@go clean

swag: ## Generate Swagger documentation
	@echo "Generating Swagger docs..."
	@swag init
	@echo "Swagger docs generated!"

docker-up: ## Start MySQL with Docker Compose
	@echo "Starting MySQL..."
	@docker-compose up -d mysql
	@echo "MySQL started!"

docker-up-postgres: ## Start PostgreSQL with Docker Compose
	@echo "Starting PostgreSQL..."
	@docker-compose --profile postgres up -d postgres
	@echo "PostgreSQL started!"

docker-down: ## Stop all services
	@echo "Stopping services..."
	@docker-compose down
	@echo "Services stopped!"

dev: docker-up swag run ## Start development environment (MySQL + Swagger + App)

dev-postgres: docker-up-postgres swag run ## Start development environment (PostgreSQL + Swagger + App)

install-deps: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod tidy
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "Dependencies installed!"

setup: install-deps docker-up swag ## Complete setup for development
	@echo "Setup completed! Run 'make run' to start the application."
