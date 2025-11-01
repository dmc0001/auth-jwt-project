MIGRATIONS_DIR=./migrations

migrate-up:
	source .env && goose -dir $(MIGRATIONS_DIR) mysql "$$DB_DSN" up

migrate-down:
	source .env && goose -dir $(MIGRATIONS_DIR) mysql "$$DB_DSN" down

migrate-status:
	source .env && goose -dir $(MIGRATIONS_DIR) mysql "$$DB_DSN" status