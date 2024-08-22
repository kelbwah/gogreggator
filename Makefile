# --- Loading .env vars ---
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# --- Make targets ---
run:
	@./bin/gogreggator

build:
	@go build -o ./bin/gogreggator ./cmd/app/main.go

test:
	@go test -v ./...

db-up:
	@cd internal/sql/schema && goose postgres $(DATABASE_URL) up

db-down:
	@cd internal/sql/schema && goose postgres $(DATABASE_URL) down

