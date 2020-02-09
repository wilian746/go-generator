package instance

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wilian746/gorm-crud-generator/internal/utils/logger"
)

var ErrDialectInvalid = errors.New("dialect not valid")
var ErrURIInvalid = errors.New("uri not valid")

func GetConnection(dialectName, uri string) *gorm.DB {
	if !dialectIsAllowed(dialectName) {
		logger.PANIC("Error", ErrDialectInvalid)
	}
	if uri == "" {
		logger.PANIC("Error", ErrURIInvalid)
	}

	connection, err := gorm.Open(dialectName, uri)
	if err != nil {
		logger.PANIC("Error on open connection", err)
	}

	return connection
}

func dialectIsAllowed(name string) bool {
	for _, value := range typesValid() {
		if value == name {
			return true
		}
	}
	return false
}

func typesValid() []string {
	return []string{"mysql", "postgres", "sqlite3", "mssql"}
}
