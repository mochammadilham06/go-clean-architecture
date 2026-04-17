run : 
	@echo "Run the application"
	go run ./cmd/main.go

build : 
	@echo "Build the application"
	go build -o bin/app main.go

start : build
	@echo "Run the application"
	./bin/app

swag :
	@echo "Generate Swagger documentation"
	swag init --generalInfo main.go

wire :
	@echo "Wire generate"
	@wire ./server/api

	.PHONY: lint audit

tidy:
	go mod tidy
	go mod verify

lint:
	@echo "Running Linter..."
	golangci-lint run --timeout=5m

audit: tidy lint
	@echo "All checks passed! Ready to push."