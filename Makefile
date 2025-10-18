.PHONY: build install test clean release

# Build binary
build:
	go build -o womba main.go

# Install locally
install:
	go install

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f womba womba-*

# Build for multiple platforms
release:
	GOOS=darwin GOARCH=amd64 go build -o womba-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o womba-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o womba-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o womba-windows-amd64.exe main.go
	@echo "✅ Built binaries for all platforms"
	@ls -lh womba-*

# Run locally
run:
	go run main.go

# Format code
fmt:
	go fmt ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

