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
