.PHONY: build run test clean

# Build the application
build:
	@echo "Building the application..."
	@go build -o bin/gobank

# Run the application
run: build
	@echo "Running the application..."
	@./bin/gobank

# Test the application
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin
