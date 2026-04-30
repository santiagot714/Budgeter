# Makefile for Budgeter project
include .env
export
# Variables
DB_URL=$(DATABASE_URL)
MIGRATIONS_PATH=db/migrations

.PHONY: migrate migrate-down migrate-force migrate-create migrate-version

migrate:
	@echo "Running migrations..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	@echo "Running migrations down..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down 1

migrate-force:
	@echo "Running migrations force..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" force $(version)

migrate-create:
	@echo "Creating migration..."
	@migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(name)

migrate-version:
	@echo "Showing migration version..."
	@migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" version