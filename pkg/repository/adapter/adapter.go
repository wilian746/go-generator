package adapter

import (
	"github.com/jinzhu/gorm"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/response"
)

type Database struct {
	connection *gorm.DB
	logMode    bool
}

type Interface interface {
	GetLogMode() bool
	SetLogMode(logMode bool) Interface

	Connection(tableName string) *gorm.DB

	ParseGormQueryToDefaultResponse(result *gorm.DB) *response.Response

	Health() bool
	Find(transaction *gorm.DB, condition, entity interface{}, tableName string) *response.Response
	Create(transaction *gorm.DB, entity interface{}, tableName string) *response.Response
	Update(transaction *gorm.DB, condition, entity interface{}, tableName string) *response.Response
	Delete(transaction *gorm.DB, condition, entity interface{}, tableName string) *response.Response
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

func (d *Database) ParseGormQueryToDefaultResponse(query *gorm.DB) *response.Response {
	return response.NewDefaultResponse(query)
}

func (d *Database) getDatabaseConnection(transaction *gorm.DB, tableName string) *gorm.DB {
	if transaction != nil {
		return transaction.Table(tableName).LogMode(d.logMode)
	}

	return d.Connection(tableName)
}

func (d *Database) Health() bool {
	return d.connection.DB().Ping() == nil
}

func (d *Database) Find(transaction *gorm.DB, condition, entity interface{}, tableName string) *response.Response {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Where(condition).Find(entity))
}

func (d *Database) Create(transaction *gorm.DB, entity interface{}, tableName string) *response.Response {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Create(entity))
}

func (d *Database) Update(transaction *gorm.DB, condition, entity interface{}, tableName string) *response.Response {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Where(condition).Updates(entity))
}

func (d *Database) Delete(transaction *gorm.DB, condition, entity interface{}, tableName string) *response.Response {
	connection := d.getDatabaseConnection(transaction, tableName)

	return d.ParseGormQueryToDefaultResponse(connection.Where(condition).Delete(entity))
}
