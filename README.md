# Agent Service


## Set up

```shell
brew install sqlc
brew install golang-migrate
```

## Architecture
The service is structured following the principles of Clean Architecture.
We split into the following layers: `presentation`, `application`, `data`, and `domain`.

The presentation layer is handling HTTP requests using the Gin framework.
The application layer contains the business logic and orchestrates the flow of data between the presentation layer and 
the data layer.

The data layer is responsible for interacting with the PostgreSQL database using `sqlc` for type-safe SQL queries. 
The application layer is the only layer that interacts with the data layer.
The data layer returns data transfer objects (DTOs) to the application layer, which then maps them to domain entities.

We do not implement an additional repository pattern, as `sqlc` already provides a clear interface for data access.
Nullable values in `sqlc` are implemented with pointer types. (The application layer is responsible for handling nil 
values appropriately.)

We use migration files to manage the database schema.
The domain layer defines the core entities and their behaviors.


## SQLC 
To generate the database models and queries, run the following command from the /assistants directory:
```shell
make generate-repository-code
```

We use `sqlc` as tool to avoid ORM code. 
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

There is no package in golang like `alembic` in python that automatically scans the project for required new migrations. 
You need to write them manually. 

`golang-migrate` only helps to execute them in a controlled manner and keeps track of which migrations have already 
been applied to the database.

## Validation
We use the `go-playground/validator` package for validation of incoming requests.
Validation tags are added to the request structs in the `presentation` layer.

## Further reading
Very nice tutorial on Golang:
https://golang.howtos.io/

Split the route into request parsing, validation and response handling:
https://medium.com/@rluders/improving-request-validation-and-response-handling-in-go-microservices-cc54208123f2

Good `sqlc` example
https://conroy.org/introducing-sqlc

Tips on robust error handling, custom errors, error wrapping, error-groups (err-groups), ...
https://leapcell.io/blog/robust-go-best-practices-for-error-handling