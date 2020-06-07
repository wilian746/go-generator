package adapter

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/response"
)

var ErrRecordNotFound = errors.New("record not found")

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

	StartTransaction() Interface
	CommitTransaction() *response.Response
	RollbackTransaction() *response.Response

	Find(query *gorm.DB, entity interface{}, tableName string) *response.Response
	Create(entity interface{}, tableName string) *response.Response
	Update(condition, entity interface{}, tableName string) *response.Response
	Delete(condition interface{}, tableName string) *response.Response
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

func (d *Database) Health() bool {
	return d.connection.DB().Ping() == nil
}

func (d *Database) StartTransaction() Interface {
	d.connection = d.Connection("").Begin()
	return d
}

func (d *Database) CommitTransaction() *response.Response {
	return d.ParseGormQueryToDefaultResponse(d.connection.Commit())
}

func (d *Database) RollbackTransaction() *response.Response {
	return d.ParseGormQueryToDefaultResponse(d.connection.Rollback())
}

func (d *Database) Find(query *gorm.DB, entity interface{}, tableName string) *response.Response {
	queryFind := query.Table(tableName).Find(entity)
	return d.ParseGormQueryToDefaultResponse(queryFind)
}

func (d *Database) Create(entity interface{}, tableName string) *response.Response {
	queryCreate := d.Connection(tableName).Create(entity)
	return d.ParseGormQueryToDefaultResponse(queryCreate)
}

func (d *Database) Update(condition, entity interface{}, tableName string) *response.Response {
	queryUpdate := d.Connection(tableName).Where(condition).Update(entity)
	return d.ParseGormQueryToDefaultResponse(queryUpdate)
}

func (d *Database) Delete(condition interface{}, tableName string) *response.Response {
	queryDelete := d.Connection(tableName).Where(condition).Delete(nil)
	return d.ParseGormQueryToDefaultResponse(queryDelete)
}
