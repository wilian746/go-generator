package product

import (
	"encoding/json"
	"errors"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	EntitiesRepository "github.com/wilian746/gorm-crud-generator/pkg/repository/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"io"
	"time"
)

type ProductRules struct{}

func NewRules() *ProductRules {
	return &ProductRules{}
}

func (p *ProductRules) ConvertIoReaderToStruct(data io.Reader) (body interface{}, err error) {
	if data == nil {
		return body, errors.New("body is invalid")
	}
	err = json.NewDecoder(data).Decode(&body)
	return body, err
}

func (p *ProductRules) GetMock() interface{} {
	return product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.New().String(),
	}
}

func (p *ProductRules) Migrate(connection *gorm.DB, toCreate bool, model EntitiesRepository.Interface) *gorm.DB {
	connection = connection.Table(model.TableName())
	connection.AutoMigrate(model)
	connection.New().Not(map[string]interface{}{"id": uuid.Nil}).Delete(model)
	if toCreate {
		connection.Create(model)
	}
	return connection
}

func (p *ProductRules) Validate(model interface{}) error {
	productModel, err := product.InterfaceToModel(model)
	if err != nil {
		return err
	}

	return Validation.ValidateStruct(productModel,
		Validation.Field(&productModel.ID, Validation.Required, is.UUIDv4),
		Validation.Field(&productModel.Name, Validation.Required, Validation.Length(3, 50)),
	)
}
