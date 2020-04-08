package product

import (
	"github.com/google/uuid"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"time"
)

type Controller struct {
	repository adapter.Interface
}

type Interface interface {
	ListOne(ID uuid.UUID) (entity product.Product, err error)
	ListAll() (entities []product.Product, err error)
	Create(entity *product.Product) (uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error
}

func NewController(repository adapter.Interface) Interface {
	return &Controller{repository:repository}
}

func (c *Controller) ListOne(ID uuid.UUID) (entity product.Product, err error) {
	response := c.repository.Find(map[string]interface{}{"id": ID}, &entity, entity.TableName())

	if err := response.Error(); err != nil {
		return entity, err
	}

	return entity, nil
}

func (c *Controller) ListAll() (entities []product.Product, err error) {
	var entity product.Product
	response := c.repository.Find(map[string]interface{}{}, &entities, entity.TableName())

	if err := response.Error(); err != nil {
		return entities, err
	}

	return entities, nil
}

func (c *Controller) Create(entity *product.Product) (uuid.UUID, error) {
	entity.CreatedAt = time.Now()
	response := c.repository.Create(nil, &entity, entity.TableName())

	if err := response.Error(); err != nil {
		return entity.ID, err
	}

	return entity.ID, nil
}

func (c *Controller) Update(ID uuid.UUID, entity *product.Product) error {
	entity.UpdatedAt = time.Now()

	response := c.repository.Update(nil, map[string]interface{}{"id": ID}, &entity, entity.TableName())

	return response.Error()
}

func (c *Controller) Remove(ID uuid.UUID) error {
	var entity product.Product

	response := c.repository.Delete(nil, map[string]interface{}{"id": ID}, &entity, entity.TableName())

	return response.Error()
}
