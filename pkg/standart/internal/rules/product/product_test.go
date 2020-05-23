package product

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/gorm-crud-generator/pkg/repository/instance"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities/product"
	"math"
	"testing"
	"time"
)

func TestNewRules(t *testing.T) {
	t.Run("Should return instance rules", func(t *testing.T) {
		assert.IsType(t, NewRules(), &Rules{})
	})
}

func TestRules_ConvertIoReaderToProduct(t *testing.T) {
	t.Run("Should parse ioRead to product", func(t *testing.T) {
		r := NewRules()
		data := &product.Product{
			Base: entities.Base{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: uuid.New().String(),
		}
		b, _ := data.Bytes()
		products, err := r.ConvertIoReaderToProduct(bytes.NewReader(b))
		assert.NoError(t, err)
		assert.NotEmpty(t, products)
	})
	t.Run("Should return err when parse data nil", func(t *testing.T) {
		r := NewRules()
		products, err := r.ConvertIoReaderToProduct(nil)
		assert.Error(t, err)
		assert.Nil(t, products)
	})
	t.Run("Should return err when parse data wrong", func(t *testing.T) {
		r := NewRules()
		b, _ := json.Marshal(math.NaN())
		products, err := r.ConvertIoReaderToProduct(bytes.NewReader(b))
		assert.Error(t, err)
		assert.Nil(t, products)
	})
}

func TestRules_GetMock(t *testing.T) {
	t.Run("should return mock correctly", func(t *testing.T) {
		r := NewRules()
		p := r.GetMock()
		assert.NotEqual(t, p.ID, uuid.Nil)
	})
}

func TestRules_Migrate(t *testing.T) {
	t.Run("should migrate new table correctly", func(t *testing.T) {
		r := NewRules()
		conn := instance.GetConnection("sqlite3", ":memory:")
		assert.NoError(t, conn.Error)
		conn = r.Migrate(conn, r.GetMock())
		assert.NoError(t, conn.Error)
	})
}

func TestRules_Validate(t *testing.T) {
	t.Run("Should return error if name is empty", func(t *testing.T) {
		r := NewRules()
		assert.Error(t, r.Validate(&product.Product{}))
	})
}
