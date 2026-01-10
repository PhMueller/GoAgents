SHELL := /bin/bash

-include .env
export

up:
	docker-compose up

down:
	docker-compose down

show-tables:
	docker exec -it go-postgres psql $${DATABASE_DB} -U $${DATABASE_USER} -c "\dt"

migrate-up:
	migrate -database $${DATABASE_URL} -path db/migrations up

migrate-down:
	migrate -database $${DATABASE_URL} -path db/migrations down

# Force a specific migration version
# E.g. to force to version 5, run: make migrate-force:5
migrate-force\:%:
	migrate -database $${DATABASE_URL} -path db/migrations force $(*)

# Generate the SQLC repository code.
# It uses the sqlc.yaml configuration file and the migrations in the db folder.
generate-repository-code:
	sqlc generate


# Needs to be executed from the assistants directory
test:
	go test -v ./...

test-coverage:
	go test -coverprofile cover.out ./... \
	&& go tool cover -html cover.out -o cover.html

test-coverage-show:
	go test -coverprofile cover.out ./... \
	&& go tool cover -html cover.out -o cover.html \
	&& open cover.html