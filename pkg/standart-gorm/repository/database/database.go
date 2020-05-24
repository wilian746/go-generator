package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"    // gorm dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"    // gorm dialect
	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm dialect
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // gorm dialect
	"log"
)

var ErrDialectInvalid = errors.New("dialect not valid")
var ErrURIInvalid = errors.New("uri not valid")

func GetConnection(dialectName, uri string) *gorm.DB {
	if !dialectIsAllowed(dialectName) && ErrDialectInvalid != nil {
		log.Panic("Error", ErrDialectInvalid)
	}
	if uri == "" && ErrURIInvalid != nil {
		log.Panic("Error", ErrURIInvalid)
	}

	connection, err := gorm.Open(dialectName, uri)
	if err != nil {
		log.Panic("Error on open connection", err)
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
