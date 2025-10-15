.PHONY: dep lint start stop

## Ensure all dependencies are up to date
deps:
	@go mod tidy && go mod download

## Run linters
lint:
	@golangci-lint run ./...

## Start the application
start:
	@docker compose up -d

## Stop the application
stop:
	@docker compose down

