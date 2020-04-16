package product

import (
	"encoding/json"
	"errors"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	entities2 "github.com/wilian746/gorm-crud-generator/pkg/repository/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"io"
	"time"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (p *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model)
}

func (p *Rules) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.New().String(),
	}
}

func (p *Rules) Migrate(connection *gorm.DB, toCreate bool, model entities2.Interface) *gorm.DB {
	connection = connection.Table(model.TableName())
	connection.AutoMigrate(model)
	connection.New().Not(map[string]interface{}{"id": uuid.Nil}).Delete(model)
	if toCreate {
		connection.Create(model)
	}
	return connection
}

func (p *Rules) Validate(model interface{}) error {
	productModel, err := product.InterfaceToModel(model)
	if err != nil {
		return err
	}

	return Validation.ValidateStruct(productModel,
		Validation.Field(&productModel.ID, Validation.Required, is.UUIDv4),
		Validation.Field(&productModel.Name, Validation.Required, Validation.Length(3, 50)),
	)
}
