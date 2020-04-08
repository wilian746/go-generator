package rules

import (
	"github.com/jinzhu/gorm"
	EntitiesRepository "github.com/wilian746/gorm-crud-generator/pkg/repository/entities"
	"io"
)

type Interface interface {
	ConvertIoReaderToStruct(data io.Reader) (body interface{}, err error)
	GetMock() interface{}
	Migrate(connection *gorm.DB, toCreate bool, model EntitiesRepository.Interface) *gorm.DB
	Validate(model interface{}) error
}
