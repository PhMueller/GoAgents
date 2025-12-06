# Agent Service


## Set up

```shell
brew install sqlc
brew install golang-migrate
```

## SQLC 
To generate the database models and queries, run the following command from the /assistants directory:
```shell
make generate-repository-code
```

We use sqlc as tool to avoid ORM code. 
It generates type-safe code from SQL queries and schema definition.
The configuration file is located in `sqlc.yml`.

We refrain from having a separate schema.sql file but use the migration files in `db/migrations` as the 
source of truth for the database schema.

For each model, we have a corresponding SQL file, e.g. `db/queries/threads.sql`.
This file contains the standard CRUD operations, as well as any custom queries we need.

### Notes
- We override the default uuid type to be `github.com/google/uuid.UUID` in `sqlc.yml`.
- Do not edit the generated files directly.
- Do not alter already existing migration files. Create new ones for any schema changes.

## Migration

Install the dependency, we use golang-migrate
```shell
brew install golang-migrate
```

From the /assistants directory, create a first migration file
```shell 
migrate create -ext sql -dir db/migrations -seq <NAME_OF_MIGRATION>
```
Naming convention for <NAME_OF_MIGRATION> is:
`<TABLE_AFFECTED_IN_ALL_CAPS>_<OPERATION_OR_DESCRIPTION_IN_LOWERCASE>`

Execute the migration
```shell
make migrate-up 
make migrate-down
```

If an exception happens, you might need to force down to a version: 
```shell
# MIGRATION_NUMBER is the number of the migration you want to force down to
make migrate-force:<MIGRATION_NUMBER>
```
