package adapter

import (
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
	defaultresponse "github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/response"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) GetLogMode() bool {
	args := m.MethodCalled("GetLogMode")
	return args.Get(0).(bool)
}
func (m *Mock) SetLogMode(logMode bool) Interface {
	args := m.MethodCalled("SetLogMode")
	return args.Get(0).(*Mock)
}
func (m *Mock) Health() bool {
	args := m.MethodCalled("Health")
	return args.Get(0).(bool)
}
func (m *Mock) Connection(tableName string) *gorm.DB {
	args := m.MethodCalled("Connection")
	return args.Get(0).(*gorm.DB)
}
func (m *Mock) ParseGormQueryToDefaultResponse(result *gorm.DB) *defaultresponse.Response {
	args := m.MethodCalled("ParseGormQueryToDefaultResponse")
	return args.Get(0).(*defaultresponse.Response)
}
func (m *Mock) Find(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.Response {
	args := m.MethodCalled("Find")
	return args.Get(0).(*defaultresponse.Response)
}
func (m *Mock) Create(transaction *gorm.DB, entity interface{}, tableName string) *defaultresponse.Response {
	args := m.MethodCalled("Create")
	return args.Get(0).(*defaultresponse.Response)
}
func (m *Mock) Update(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.Response {
	args := m.MethodCalled("Update")
	return args.Get(0).(*defaultresponse.Response)
}
func (m *Mock) Delete(transaction *gorm.DB, condition, entity interface{}, tableName string) *defaultresponse.Response {
	args := m.MethodCalled("Delete")
	return args.Get(0).(*defaultresponse.Response)
}
