package rules

import (
	"github.com/jinzhu/gorm"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/entities"
	"io"
)

type Interface interface {
	ConvertIoReaderToStruct(data io.Reader, model interface{}) (body interface{}, err error)
	GetMock() interface{}
	Migrate(connection *gorm.DB, toCreate bool, model entities.Interface) *gorm.DB
	Validate(model interface{}) error
}
