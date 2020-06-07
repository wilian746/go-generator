# Standart Gorm

This project was generated from the [go-generator](https://github.com/wilian746/go-generator) to make the same use following the following steps

## Install dependence
Checkout in the directory installed and run command
```bash
go get -u ./...
```

## Run application
Run command to up application:
```bash
go run cmd/main.go
```

## Environments
This environments for use for you application:

| Name             | Default Value  | Type          |
|------------------|----------------|---------------|
| PORT             | 8080           | int           |
| TIMEOUT          | 30             | int           |
| DATABASE_DIALECT | sqlite3        | string        |
| DATABASE_URI     | :memory:       | string        |
| SWAGGER_HOST     | localhost:8080 | string        |

## Database
This project use the [GORM](https://gorm.io/) to manipulate a relational database if you do change dialect so change in environments.
So examples of connections for you use.

### Sqlite3
You can change uri to path of file database too. `Ex.: /path/dest/file/standartgorm.db`

| Environment      | Value          |
|------------------|----------------|
| DATABASE_DIALECT | sqlite3        |
| DATABASE_URI     | :memory:       |

### PostgreSQL
| Environment      | Value          |
|------------------|----------------|
| DATABASE_DIALECT | postgres       |
| DATABASE_URI     | host=localhost port=5432 user=admin dbname=standartgorm password=admin |

### MySQL
| Environment      | Value          |
|------------------|----------------|
| DATABASE_DIALECT | mysql          |
| DATABASE_URI     | admin:admin@localhost:3306/standartgorm?charset=utf8&parseTime=True&loc=Local |

### SQL Server
| Environment      | Value          |
|------------------|----------------|
| DATABASE_DIALECT | mssql          |
| DATABASE_URI     | sqlserver://admin:admin@localhost:1433?database=standartgorm |


## Swagger
To update swagger.json, you need run command into **root folder location** of the project
```bash
swag ini -g ./cmd/main.go
```
To more information you can see [docs of SWAG](https://github.com/swaggo/swag)

## Migrations
This project contains migrations to update database relational, install [go-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to use command-line and setup your database usage.
And to apply and usage it's using [go-migrate](https://github.com/golang-migrate/migrate)

#### Example of the Connection String
- [POSTGRESQL](https://github.com/golang-migrate/migrate/tree/master/database/postgres) Example:
```text
postgres://root:root@localhost:5432/postgres
```
- [MYSQL](https://github.com/golang-migrate/migrate/tree/master/database/mysql) Example:
```text
mysql://root:root@tcp(localhost:3306)/mysql
```
- [SQLSERVER](https://github.com/golang-migrate/migrate/tree/master/database/sqlserver) Example:
```text
sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=master
```

#### Install driver
For usage this lib your need install driver for usage.

- Example using `postgres`
```bash
go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
```
- Example using `mysql`
```bash
go get -tags 'mysql' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
```
- Example using `sqlserver`
```bash
go get -tags 'sqlserver' -u github.com/golang-migrate/migrate/v4/cmd/migrate/
```

#### Running

Migrate Create - Create new migration
```bash
migrate create -ext sql -dir ./migrations/{DRIVER} {MIGRATION_NAME}
```

Migrate Up - Up all migrations in sequence by date
```bash
migrate -path ./migrations/{DRIVER} -database {CONNECTION_STRING} up
```

Migrate Down - Down one migration in sequence by date
```bash
migrate -path ./migrations/{DRIVER} -database {CONNECTION_STRING} down
```

Migrate Steps - Up/Down NUMBER update/downgrade in sequence by date
```bash
migrate -path ./migrations/{DRIVER} -database {CONNECTION_STRING} up 1
```
```bash
migrate -path ./migrations/{DRIVER} -database {CONNECTION_STRING} down 1
```

Migrate Specific - Migrate to specific version of migrations in folder selected
```bash
migrate -path ./migrations/{DRIVER} -database {CONNECTION_STRING} goto {VERSION}
```

Migrate Version - See your actual version
```bash
migrate -path ./migrations/{DRIVER} -database {CONNECTION_STRING} version
```
