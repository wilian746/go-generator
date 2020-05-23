package product

import (
	"encoding/json"
	"github.com/wilian746/gorm-crud-generator/pkg/standart/internal/entities"
)

type Product struct {
	entities.Base
	Name string `json:"name"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}
