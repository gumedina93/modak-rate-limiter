.PHONY: install
install:
	@echo "=> Installing dependencies"
	@go mod tidy

.PHONY: run
run:
	@echo "=> Running rate limiter application"
	@go run ./...

.PHONY: test
test:
	@echo "=> Running tests"
	@go test ./...