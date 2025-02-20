.PHONY: build run test clean migrate docker-up docker-down

# Build the application
build:
	go build -o bin/app cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out

# Database migrations
migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/echo_store?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/echo_store?sslmode=disable" down

# Docker commands
docker-up:
	docker-compose -f deployments/docker-compose.yml up --build -d

docker-down:
	docker-compose -f deployments/docker-compose.yml down

# Generate swagger documentation
swagger:
	swag init -g cmd/main.go

# Lint the code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Install dependencies
deps:
	go mod download
	go mod tidy 