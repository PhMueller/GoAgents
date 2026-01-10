# Agent Service

## TODO List

- [ ] Threads
    - [X] Create Thread
    - [X] Get Thread by ID
    - [ ] Get Threads Info: 
      - [ ] Add pagination
      - [ ] Add filter on user related data (can the user access these threads)
    - [ ] Delete Thread

- [ ] Messages
  - [X] Create Message
  - [X] Get Message by ID
  - [X] Get Messages by Thread ID

- [ ] Error handling improvement
  - [ ] Implement error handling middleware (https://gin-gonic.com/en/docs/examples/error-handling-middleware/)
- [ ] How to prevent extra parameters in request body?
- [ ] Authentication and Authorization
  - [ ] Create a login endpoint (receives a client_id and client_secret, returns jwt as bearer token)
  - [ ] Define claims: How can do what? (scopes)
  - [ ] For now create register endpoint (only with admin key reachable) to create a new client (client_id, client_secret, scopes)
  - [ ] Add authorization middleware to check for valid jwt token with correct scopes.
- [ ] Rate limiting
- [ ] Logging
- [ ] Swagger UI
- [ ] First try to configure llm 

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

The data layer (repository) is responsible for interacting with the PostgreSQL database using `sqlc` for type-safe SQL queries. 
The application layer (services) is the only layer that interacts with the data layer.
The data layer returns data transfer objects (DTOs) to the application layer. The application layer maps them to domain entities.

We do not implement an additional repository pattern, as `sqlc` already provides a clear interface for data access.

Nullable values in `sqlc` are implemented with pointer types. (The application layer is responsible for handling nil 
values appropriately.). `sqlc` often maps the types in our case to `pgtype` types, e.g. `pgtype.Timestamptz`.
This interferes with the nullable pointer handling. We manually adjust the sqlc.yml to override pgtypes to standard go 
types.
Issue on github explaining this for timestampz. https://github.com/sqlc-dev/sqlc/issues/814#issuecomment-3042290683
We like to highlight that you have to override for the case that the type is not nullable and for the case that it is 
nullable.

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

## Authorization and Authentication
We plan to implement a client-based authentication and authorization system using JWT tokens.

This feature consists of the following components: 

### Admin: Client Management
- CRUD operations on /admin/clients.
- Only accessible with an admin key in header "X-Admin-Key: <admin_key>"
- Here, we manage the clients with specific scopes.

### Admin Client Session Management
- CRUD operations on /admin/clients/:client_id/sessions.
- Only accessible with an admin key in header "X-Admin-Key: <admin_key>"
- This is an internal endpoint that is developed to fulfill the "API first" principle.

#### Workflow: 
- The admin has created a client with specific scopes. 
- The client now calls the /auth/token endpoint. 
- When the client's request is valid, we create a session for the client.
- The session contains information about the client, the issued token, expiration date, and scopes.

### Client: Token Management
- The user calls the /auth/token endpoint with their client_id and client_secret to obtain a JWT token.
- This token is used to authenticate subsequent requests to the API.
- A middleware checks the validity of the token, its scopes, and expiration date for each request.

### Architecture and implementation details

Create two tables in db. 
--> Clients ( client_id, client_secret, scopes, ... )
--> ClientSessions ( session_id, client_id, issued_at, expires_at, scopes, ... )

Only admin app can create / modify new clients as well as the client sessions. 
The client app can only request to get a new session by calling /auth/token with valid client_id and client_secret.
We dont take care of the case that a user creates to many open sessions yet. 

### Authorization Workflow: 

Middleware checks for valid token, scopes, and expiration date. Some endpoints have different scopes(?)

- Prerequisite: There exists a client_id, client_secret in the db. 
- Client calls /auth/token with client_id and client_secret in body.
- Server checks for valid client_id and client_secret in db.
- Server creates a new session in client_sessions table with issued_at, expires_at, scopes.
- Server creates a JWT token with claims: client_id, issued_at, expires_at, scopes
- JWT is signed with server's private key.
- Server returns token to client.
- Client sends in header "Authorization: Bearer <token>"
- Server extracts token from header.
- Server verifies token signature with server's public key.
- Server Middleware checks for expiration date and scopes.
- Request is processed if token is valid.

Open questions: 
- how to allow different clients to interact with data? 
- E.g. client A can only access threads it created. Client B can access all threads.
- --> Define claims in token accordingly

## Validation

We use the `go-playground/validator` package for validation of incoming requests.
Validation tags are added to the request structs in the `presentation` layer.


## Testing

### Integration tests
For that we would need to mock our db! An example using `mockgen` and `sqlc`. 
https://medium.com/@emmyvera01/enhancing-http-api-testing-in-golang-with-a-mock-db-harnessing-the-power-of-mockgen-b509dc6b7e75


## Further reading

Very nice tutorial on Golang:
https://golang.howtos.io/

Split the route into request parsing, validation and response handling:
https://medium.com/@rluders/improving-request-validation-and-response-handling-in-go-microservices-cc54208123f2

Good `sqlc` example
https://conroy.org/introducing-sqlc

Tips on robust error handling, custom errors, error wrapping, error-groups (err-groups), ...
https://leapcell.io/blog/robust-go-best-practices-for-error-handling