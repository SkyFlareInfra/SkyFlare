GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Use `make dev` for development with live reloads
dev:
	air

# Use `make build` to create a production build
build:
	$(GO_BUILD_ENV) go build -o bin/app main.go
	docker build -t skyflareinfra/skyflare:latest .
	docker push skyflareinfra/skyflare:latest

# Use `make run` to run the application locally
run:
	$(GO_BUILD_ENV) go run main.go

# Use `make install-test-tools` to install test generation tools
install-test-tools:
	go install github.com/cweill/gotests/gotests@latest
	go install github.com/rakyll/gotest@latest

# Use `make unit-test` to generate unit tests with actual test cases
unit-test: install-test-tools
	@echo "Generating tests for all packages..."
	@for dir in $$(go list ./...); do \
		pkg_path=$$(echo $$dir | sed 's|github.com/SkyFlareInfra/SkyFlare/||'); \
		if [ -d "$$pkg_path" ]; then \
			echo "Generating tests for package: $$pkg_path"; \
			find "$$pkg_path" -name "*.go" -type f | grep -v "_test.go" | while read file; do \
				gotests -all -w "$$file"; \
			done; \
		fi; \
	done

# Use `make test` to run all tests with ENVIRONMENT=TEST
test:
	ENVIRONMENT=TEST go test ./...

# Use `make test-verbose` to run all tests with verbose output
test-verbose:
	ENVIRONMENT=TEST go test -v ./...

# Use `make test-coverage` to run tests with coverage
test-coverage:
	ENVIRONMENT=TEST go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Use `make test-bench` to run benchmarks
test-bench:
	ENVIRONMENT=TEST go test -bench=. -benchmem ./...

# Use `make remove-test-artifacts` to remove test artifacts
remove-test-artifacts:
	@echo "Removing test artifacts..."
	@find . -type f -name '*_test.go' -exec echo "Removing: {}" \; -exec rm -f {} \;
	@echo "Test artifacts removed."

# Use `make lint` to run linters
lint:
	golangci-lint run ./...

# Use `make migrate` to run database migrations
migrate:
	$(GO_BUILD_ENV) go run db.go --migrate

# Use `make clean` to clean up build artifacts
clean:
	rm -rf bin/ coverage.out coverage.html

# Use `make help` to display this help message
help:
	@echo "Makefile commands:"
	@echo "  dev               	- Start development server with live reloads"
	@echo "  build             	- Build the application for production"
	@echo "  run               	- Run the application locally"
	@echo "  install-test-tools	- Install test generation tools"
	@echo "  unit-test         	- Generate unit tests with actual test cases"
	@echo "  test              	- Run all tests with ENVIRONMENT=TEST"
	@echo "  test-verbose      	- Run all tests with verbose output"
	@echo "  test-coverage     	- Run tests with coverage report"
	@echo "  test-bench        	- Run benchmarks"
	@echo "  lint              	- Run linters"
	@echo "  migrate           	- Run database migrations"
	@echo "  clean             	- Clean up build artifacts"
	@echo "  help              	- Display this help message"

.PHONY: dev build run install-test-tools unit-test test test-verbose test-coverage test-bench lint migrate clean help
