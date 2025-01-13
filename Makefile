.PHONY: migrate generate build run

migrate:
	@Pgo run ./src/cmd/migrate

generate:
	@templ generate ./... -lazy

build: generate
	@go build -o bin/budgetapp ./src/cmd/budgetapp/...

run: build
	@./bin/budgetapp
