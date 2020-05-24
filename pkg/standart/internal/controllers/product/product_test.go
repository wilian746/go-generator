package product

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	EntitiesProduct "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	RulesProduct "github.com/wilian746/gorm-crud-generator/pkg/standart/internal/rules/product"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/repository/adapter"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/repository/database"
	"testing"
)

func TestNewController(t *testing.T) {
	t.Run("Should create instance controller", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		assert.NotEmpty(t, NewController(adapter.NewAdapter(conn)))
	})
}

func TestController_Create(t *testing.T) {
	t.Run("Should create product on database", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		id, err := controller.Create(productMock)
		assert.NoError(t, err)
		assert.NotEqual(t, id, uuid.Nil)
	})
	t.Run("Should not create product on database because not migrate table", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		id, err := controller.Create(productMock)
		assert.Error(t, err)
		assert.Equal(t, id, uuid.Nil)
	})
}

func TestController_ListAll(t *testing.T) {
	t.Run("Should return empty list of product on database", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		list, err := controller.ListAll()
		assert.NoError(t, err)
		assert.Equal(t, list, []EntitiesProduct.Product{})
	})
	t.Run("Should not return empty list of product on database", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		id, err := controller.Create(productMock)
		assert.NoError(t, err)
		assert.NotEqual(t, id, uuid.Nil)
		list, err := controller.ListAll()
		assert.NoError(t, err)
		assert.NotEqual(t, list, []EntitiesProduct.Product{})
	})
	t.Run("Should return error on list of product on database because not migrate", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		list, err := controller.ListAll()
		assert.Error(t, err)
		assert.Equal(t, list, []EntitiesProduct.Product{})
	})
}

func TestController_ListOne(t *testing.T) {
	t.Run("Should return error record not found product on database", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		_, err := controller.ListOne(uuid.New())
		assert.Error(t, err)
		assert.Equal(t, err, adapter.RecordNotFound)
	})
	t.Run("Should return error product on database because not migrate", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		_, err := controller.ListOne(uuid.New())
		assert.Error(t, err)
		assert.NotEqual(t, err, adapter.RecordNotFound)
	})
	t.Run("Should not return empty product on database", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		id, err := controller.Create(productMock)
		assert.NoError(t, err)
		assert.NotEqual(t, id, uuid.Nil)
		founded, err := controller.ListOne(productMock.ID)
		assert.NoError(t, err)
		assert.Equal(t, founded.ID, productMock.ID)
	})
}

func TestController_Update(t *testing.T) {
	t.Run("Should return error record not found product on database when update", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		ID := uuid.New()
		_, err := controller.ListOne(ID)
		assert.Error(t, err)
		err = controller.Update(ID, productMock)
		assert.Error(t, err)
		assert.Equal(t, err, adapter.RecordNotFound)
	})
	t.Run("Should return error product on database because not migrate when update", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		err := controller.Update(uuid.New(), productMock)
		assert.Error(t, err)
		assert.NotEqual(t, err, adapter.RecordNotFound)
	})
	t.Run("Should udpdate product without error and check if names is different", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		id, err := controller.Create(productMock)
		assert.NoError(t, err)
		assert.NotEqual(t, id, uuid.Nil)
		founded, err := controller.ListOne(productMock.ID)
		assert.NoError(t, err)
		assert.Equal(t, founded.ID, productMock.ID)
		productMockUpdate := rules.GetMock()
		productMockUpdate.ID = productMock.ID
		err = controller.Update(productMockUpdate.ID, productMockUpdate)
		assert.NoError(t, err)
		foundedUpdated, err := controller.ListOne(productMockUpdate.ID)
		assert.NoError(t, err)
		assert.NotEqual(t, foundedUpdated.Name, productMock.Name)
	})
}

func TestController_Delete(t *testing.T) {
	t.Run("Should return error record not found product on database when delete", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		ID := uuid.New()
		_, err := controller.ListOne(ID)
		assert.Error(t, err)
		err = controller.Remove(ID)
		assert.Error(t, err)
		assert.Equal(t, err, adapter.RecordNotFound)
	})
	t.Run("Should return error product on database because not migrate when delete", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		err := controller.Remove(uuid.New())
		assert.Error(t, err)
		assert.NotEqual(t, err, adapter.RecordNotFound)
	})
	t.Run("Should delete product without error and check if product not exists", func(t *testing.T) {
		conn := database.GetConnection("sqlite3", ":memory:")
		controller := NewController(adapter.NewAdapter(conn))
		rules := RulesProduct.NewRules()
		productMock := rules.GetMock()
		rules.Migrate(conn, productMock)
		id, err := controller.Create(productMock)
		assert.NoError(t, err)
		assert.NotEqual(t, id, uuid.Nil)
		founded, err := controller.ListOne(productMock.ID)
		assert.NoError(t, err)
		assert.Equal(t, founded.ID, productMock.ID)
		err = controller.Remove(founded.ID)
		assert.NoError(t, err)
		_, err = controller.ListOne(founded.ID)
		assert.Error(t, err)
		assert.Equal(t, err, adapter.RecordNotFound)
	})
}