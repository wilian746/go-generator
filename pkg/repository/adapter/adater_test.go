package adapter

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/instance"
	defaultresponse "github.com/wilian746/gorm-crud-generator/pkg/repository/response"
	"testing"
	"time"
)

type Product struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (p *Product) tableName() string {
	return "products"
}

func GetMemoryDatabase(isTransaction bool) (*Product, *gorm.DB) {
	product := &Product{}
	connection := instance.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(product)
	connection.Where("id != 0").Delete(product)

	if !isTransaction {
		return product, connection
	}

	return product, connection.Begin()
}

func TestMock(t *testing.T) {
	mock := &Mock{}
	product := &Product{}

	t.Run("should return expected value in mock of GetLogMode", func(t *testing.T) {
		mock.On("GetLogMode").Return(true)
		assert.Equal(t, mock.GetLogMode(), true)
	})
	t.Run("should return expected value in mock of SetLogMode", func(t *testing.T) {
		mock.On("SetLogMode").Return(mock)
		assert.Equal(t, mock.SetLogMode(true), mock)
	})
	t.Run("should return expected value in mock of Connection", func(t *testing.T) {
		mock.On("Connection").Return(&gorm.DB{})
		assert.Equal(t, mock.Connection("product"), &gorm.DB{})
	})
	t.Run("should return expected value in mock of ParseGormQueryToDefaultResponse", func(t *testing.T) {
		mock.On("ParseGormQueryToDefaultResponse").Return(&defaultresponse.Response{})
		assert.Equal(t, mock.ParseGormQueryToDefaultResponse(&gorm.DB{}), &defaultresponse.Response{})
	})
	t.Run("should return expected value in mock of Find", func(t *testing.T) {
		mock.On("Find").Return(&defaultresponse.Response{})
		assert.Equal(t, mock.Find(nil, map[string]interface{}{}, product, product.tableName()), &defaultresponse.Response{})
	})
	t.Run("should return expected value in mock of Create", func(t *testing.T) {
		mock.On("Create").Return(&defaultresponse.Response{})
		assert.Equal(t, mock.Create(nil, product, product.tableName()), &defaultresponse.Response{})
	})
	t.Run("should return expected value in mock of Update", func(t *testing.T) {
		mock.On("Update").Return(&defaultresponse.Response{})
		assert.Equal(t, mock.Update(nil, map[string]interface{}{}, product, product.tableName()), &defaultresponse.Response{})
	})
	t.Run("should return expected value in mock of Delete", func(t *testing.T) {
		mock.On("Delete").Return(&defaultresponse.Response{})
		assert.Equal(t, mock.Delete(nil, map[string]interface{}{}, product, product.tableName()), &defaultresponse.Response{})
	})
}

func TestNewAdapter(t *testing.T) {
	t.Run("should return type of Database", func(t *testing.T) {
		assert.IsType(t, &Database{}, NewAdapter(&gorm.DB{}))
	})
}

func TestDatabase_SetLogMode_GetLogMode(t *testing.T) {
	t.Run("should setLog correctly", func(t *testing.T) {
		connection := instance.GetConnection("sqlite3", ":memory:")
		adapter := NewAdapter(connection)
		adapter.SetLogMode(true)
		assert.True(t, adapter.GetLogMode())
		adapter.SetLogMode(false)
		assert.False(t, adapter.GetLogMode())
	})
}

func TestDatabase_Connection(t *testing.T) {
	product := &Product{}
	connection := instance.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(&Product{})

	t.Run("should return connection of type *gorm.DB", func(t *testing.T) {
		adapter := NewAdapter(connection)
		assert.IsType(t, &gorm.DB{}, adapter.Connection(product.tableName()))
	})
}

func TestDatabase_ParseGormQueryToDefaultResponse(t *testing.T) {
	product := &Product{}
	connection := instance.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(&Product{})

	t.Run("should return response of type *defaultresponse.Response", func(t *testing.T) {
		adapter := NewAdapter(connection)
		connection := adapter.Connection(product.tableName())

		assert.IsType(t, &defaultresponse.Response{}, adapter.ParseGormQueryToDefaultResponse(connection))
	})
}

func TestDatabase_Create_Find(t *testing.T) {
	t.Run("should not exist product after create product exist in database", func(t *testing.T) {
		var productList []Product
		product, connection := GetMemoryDatabase(false)
		adapter := NewAdapter(connection)

		response := adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(0))
		assert.Len(t, productList, 0)

		product.Name = uuid.New().String()
		product.CreatedAt = time.Now()
		response = adapter.Create(nil, product, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))

		response = adapter.Find(nil, map[string]interface{}{"id": 1}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))
		assert.Len(t, productList, 1)
		assert.NotEmpty(t, productList[0].Name)
	})
	t.Run("should not exist product after create product exist in database with transaction", func(t *testing.T) {
		var productList []Product
		product, transaction := GetMemoryDatabase(true)
		adapter := NewAdapter(transaction)

		response := adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(0))
		assert.Len(t, productList, 0)

		product.Name = uuid.New().String()
		product.CreatedAt = time.Now()
		response = adapter.Create(transaction, product, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))

		response = adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))
		assert.Len(t, productList, 1)
		assert.NotEmpty(t, productList[0].Name)

		transaction.Commit()
	})
}

func TestDatabase_Update_Find(t *testing.T) {
	t.Run("should create product, update product e check if then not is equals", func(t *testing.T) {
		product, connection := GetMemoryDatabase(false)
		nameCreated := uuid.New().String()
		product.Name = nameCreated
		product.CreatedAt = time.Now()

		var productList []Product

		adapter := NewAdapter(connection)
		_ = adapter.Create(nil, product, product.tableName())

		nameUpdated := uuid.New().String()
		response := adapter.Update(nil, map[string]interface{}{
			"id": 1,
		}, &Product{
			UpdatedAt: time.Now(),
			Name:      nameUpdated,
		}, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))

		response = adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))
		assert.Len(t, productList, 1)
		assert.NotEmpty(t, productList[0].Name)
		assert.NotEqual(t, nameCreated, nameUpdated)
	})
	t.Run("should create product, update product e check if then not is equals with transaction", func(t *testing.T) {
		product, transaction := GetMemoryDatabase(true)
		nameCreated := uuid.New().String()
		product.Name = nameCreated
		product.CreatedAt = time.Now()

		var productList []Product

		adapter := NewAdapter(transaction)
		_ = adapter.Create(transaction, product, product.tableName())

		nameUpdated := uuid.New().String()
		response := adapter.Update(transaction, map[string]interface{}{
			"id": 1,
		}, &Product{
			UpdatedAt: time.Now(),
			Name:      nameUpdated,
		}, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))

		response = adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))
		assert.Len(t, productList, 1)
		assert.NotEmpty(t, productList[0].Name)
		assert.NotEqual(t, nameCreated, nameUpdated)

		transaction.Commit()
	})
}

func TestDatabase_Delete_Find(t *testing.T) {
	t.Run("should create product, update product e check if then not is equals", func(t *testing.T) {
		product, connection := GetMemoryDatabase(false)
		product.Name = uuid.New().String()
		product.CreatedAt = time.Now()

		var productList []Product

		adapter := NewAdapter(connection)
		_ = adapter.Create(nil, product, product.tableName())

		response := adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))
		assert.Len(t, productList, 1)

		response = adapter.Delete(nil, map[string]interface{}{"id": 1}, product, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))

		response = adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(0))
		assert.Len(t, productList, 0)
	})
	t.Run("should create product, update product e check if then not is equals with transaction", func(t *testing.T) {
		product, transaction := GetMemoryDatabase(true)
		product.Name = uuid.New().String()
		product.CreatedAt = time.Now()

		var productList []Product

		adapter := NewAdapter(transaction)
		_ = adapter.Create(transaction, product, product.tableName())

		response := adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))
		assert.Len(t, productList, 1)

		response = adapter.Delete(transaction, map[string]interface{}{"id": 1}, product, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(1))

		response = adapter.Find(nil, map[string]interface{}{}, &productList, product.tableName())
		assert.NoError(t, response.Error())
		assert.Equal(t, response.RowsAffected(), int64(0))
		assert.Len(t, productList, 0)

		transaction.Commit()
	})
}
