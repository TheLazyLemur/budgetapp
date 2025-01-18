.PHONY: migrate/create migrate/up generate build run

REQUIRED_BINS := go goose sqlite3 templ sqlc
CHECK_BINS := $(foreach bin,$(REQUIRED_BINS),$(if $(shell command -v $(bin) 2>/dev/null),,$(bin)))

ifeq ($(strip $(CHECK_BINS)),)
$(info All required binaries are available.)
else
$(error Missing required binaries: $(CHECK_BINS))
endif

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
