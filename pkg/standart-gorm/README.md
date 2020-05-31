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
