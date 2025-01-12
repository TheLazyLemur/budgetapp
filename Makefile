.PHONY: migrate build run

migrate:
	@go run ./src/cmd/migrate

build:
	@templ generate ./...
	@go build -o bin/budgetapp ./src/cmd/budgetapp/...

run: build
	@./bin/budgetapp
