.PHONY: migrate/create migrate/up generate build run

migrate/create:
	@goose create -dir db/migrations $(NAME) sqlite3

migrate/up:
	@go run ./src/cmd/migrate

generate:
	@sqlc generate
	@templ generate ./... -lazy

build: generate
	@go build -o bin/budgetapp ./src/cmd/budgetapp/...

run: build
	@./bin/budgetapp
