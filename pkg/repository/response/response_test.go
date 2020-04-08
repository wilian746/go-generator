package response

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefaultResponse(t *testing.T) {
	t.Run("should return equals type when create new default response", func(t *testing.T) {
		assert.IsType(t, &Response{}, NewDefaultResponse(&gorm.DB{}))
	})
}

func TestDefaultResponse_RowsAffected(t *testing.T) {
	t.Run("should return equals number of rowsAffected", func(t *testing.T) {
		query := &gorm.DB{RowsAffected: 1}

		assert.Equal(t, NewDefaultResponse(query).RowsAffected(), int64(1))
	})
}

func TestDefaultResponse_Error(t *testing.T) {
	t.Run("should return equals error", func(t *testing.T) {
		query := &gorm.DB{Error: errors.New("error when find data")}

		assert.Equal(t, NewDefaultResponse(query).Error(), errors.New("error when find data"))
	})
}
