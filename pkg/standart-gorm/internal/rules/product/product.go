package product

import (
	"encoding/json"
	"errors"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/product"
	RepositoryEntity "github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/entities"
	"io"
	"time"
)

type Rules struct{}

func NewRules() *Rules {
	return &Rules{}
}

func (p *Rules) ConvertIoReaderToProduct(data io.Reader, id uuid.UUID) (model *product.Product, err error) {
	if data == nil {
		return model, errors.New("body is invalid")
	}
	err = json.NewDecoder(data).Decode(&model)
	if err != nil {
		return model, err
	}
	if id == uuid.Nil {
		model.GenerateID()
	} else {
		model.ID = id
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

func (p *Rules) Validate(productEntity *product.Product) error {
	return Validation.ValidateStruct(productEntity,
		Validation.Field(&productEntity.ID, Validation.Required, is.UUIDv4),
		Validation.Field(&productEntity.Name, Validation.Required, Validation.Length(3, 50)),
	)
}
