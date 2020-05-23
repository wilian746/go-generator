package product

import (
	"encoding/json"
	"errors"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	RepositoryEntity "github.com/wilian746/gorm-crud-generator/pkg/repository/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"io"
	"time"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (p *Rules) ConvertIoReaderToProduct(data io.Reader) (model *product.Product, err error) {
	if data == nil {
		return model, errors.New("body is invalid")
	}
	err = json.NewDecoder(data).Decode(&model)
	if err != nil {
		return model, err
	}

	return model, p.Validate(model)
}

func (p *Rules) GetMock() *product.Product {
	return &product.Product{
		Base: entities.Base{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: uuid.New().String(),
	}
}

func (p *Rules) Migrate(connection *gorm.DB, model RepositoryEntity.Interface) *gorm.DB {
	connection = connection.Table(model.TableName())
	connection.AutoMigrate(model)
	return connection
}

func (p *Rules) Validate(product *product.Product) error {
	return Validation.ValidateStruct(product,
		Validation.Field(&product.ID, Validation.Required, is.UUIDv4),
		Validation.Field(&product.Name, Validation.Required, Validation.Length(3, 50)),
	)
}
