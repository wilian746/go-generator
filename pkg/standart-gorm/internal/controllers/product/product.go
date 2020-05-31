package product

import (
	"github.com/google/uuid"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/product"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/adapter"
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
	return &Controller{repository: repository}
}

func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error) {
	response := c.repository.Find(nil, map[string]interface{}{"id": id}, &entity, entity.TableName())

	if err := response.Error(); err != nil {
		return entity, err
	}

	return entity, nil
}

func (c *Controller) ListAll() (entities []product.Product, err error) {
	var entity product.Product
	response := c.repository.Find(nil, map[string]interface{}{}, &entities, entity.TableName())

	if err := response.Error(); err != nil {
		return entities, err
	}

	return entities, nil
}

func (c *Controller) Create(entity *product.Product) (uuid.UUID, error) {
	entity.SetCreatedAt()
	response := c.repository.Create(nil, &entity, entity.TableName())
	if err := response.Error(); err != nil {
		return uuid.Nil, err
	}

	return entity.ID, nil
}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) error {
	entity.SetUpdatedAt()
	_, err := c.ListOne(id)
	if err != nil {
		return err
	}
	response := c.repository.Update(nil, map[string]interface{}{"id": id}, &entity, entity.TableName())
	return response.Error()
}

func (c *Controller) Remove(id uuid.UUID) error {
	var entity product.Product

	response := c.repository.Delete(nil, map[string]interface{}{"id": id}, &entity, entity.TableName())
	if response.Error() == nil && response.RowsAffected() == 0 {
		return adapter.ErrRecordNotFound
	}
	return response.Error()
}
