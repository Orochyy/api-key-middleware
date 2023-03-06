DATABASE_URL := "mysql://root1:root1@tcp(go_api_database:3306)/go_api_database"
MIGRATION_SOURCE := "file://./database/migrations/"

.PHONY: migrate-up
migrate-up:
	migrate -path $(MIGRATION_SOURCE) -database $(DATABASE_URL) up

.PHONY: migrate-down
migrate-down:
	migrate -path $(MIGRATION_SOURCE) -database $(DATABASE_URL) down

#migrate -path "file://./database/migrations/" -database "mysql://root1:root1@tcp(go_api_database:3306)/go_api_database" force 1

