package adapter

import (
	"github.com/jinzhu/gorm"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/default_response"
)

type Database struct {
	connection *gorm.DB
	logMode    bool
}

type Interface interface {
	GetLogMode() bool
	SetLogMode(logMode bool) Interface

	Connection(tableName string) *gorm.DB

	ParseGormQueryToDefaultResponse(result *gorm.DB) *defaultresponse.DefaultResponse

	Find(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.DefaultResponse
	Create(transaction *gorm.DB, entity interface{}, tableName string) *defaultresponse.DefaultResponse
	Update(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.DefaultResponse
	Delete(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.DefaultResponse
}

func NewAdapter(connection *gorm.DB) Interface {
	return &Database{
		connection: connection,
	}
}

func (d *Database) GetLogMode() bool {
	return d.logMode
}

func (d *Database) SetLogMode(logMode bool) Interface {
	d.logMode = logMode
	return d
}

func (d *Database) Connection(tableName string) *gorm.DB {
	return d.connection.New().Table(tableName).LogMode(d.logMode)
}

func (d *Database) ParseGormQueryToDefaultResponse(query *gorm.DB) *defaultresponse.DefaultResponse {
	return defaultresponse.NewDefaultResponse(query)
}

func (d *Database) Find(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.DefaultResponse {
	return d.ParseGormQueryToDefaultResponse(d.Connection(tableName).Where(condition).Find(entity))
}
func (d *Database) Create(transaction *gorm.DB, entity interface{}, tableName string) *defaultresponse.DefaultResponse {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Create(entity))
}
func (d *Database) Update(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.DefaultResponse {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Where(condition).Updates(entity))
}
func (d *Database) Delete(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.DefaultResponse {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Where(condition).Delete(entity))
}

func (d *Database) getDatabaseConnection(transaction *gorm.DB, tableName string) *gorm.DB {
	if transaction != nil {
		return transaction.Table(tableName).LogMode(d.logMode)
	}

	return d.Connection(tableName)
}
