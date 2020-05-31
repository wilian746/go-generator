package product

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities"
	"github.com/wilian746/go-generator/pkg/standart-gorm/internal/entities/product"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/database"
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
		ID := uuid.New()
		data := &product.Product{
			Base: entities.Base{
				ID:        ID,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Name: uuid.New().String(),
		}
		products, err := r.ConvertIoReaderToProduct(bytes.NewReader(data.Bytes()), ID)
		assert.NoError(t, err)
		assert.NotEmpty(t, products)
	})
	t.Run("Should return err when parse data nil", func(t *testing.T) {
		r := NewRules()
		products, err := r.ConvertIoReaderToProduct(nil, uuid.New())
		assert.Error(t, err)
		assert.Nil(t, products)
	})
	t.Run("Should return err when parse data wrong", func(t *testing.T) {
		r := NewRules()
		b, _ := json.Marshal(math.NaN())
		products, err := r.ConvertIoReaderToProduct(bytes.NewReader(b), uuid.New())
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
		conn := database.GetConnection("sqlite3", ":memory:")
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
