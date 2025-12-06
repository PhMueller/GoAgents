# Agent Service


## Set up

```shell
brew install golang-migrate
```

## Migration

Install the dependency, we use golang-migrate
```shell
brew install golang-migrate
```


From the /assistants directory, create a first migration file
```shell 
migrate create -ext sql -dir db/migrations -seq <NAME_OF_MIGRATION>
```

Execute the migration
```shell
make migrate-up 
make migrate-down
```

If an exception happens, you might need to force down to a version: 
```shell
migrate -database ${DATABASE_URL} -path db/migrations force <MIGRATION_NUMBER>
migrate-down 
migrate-up
```
