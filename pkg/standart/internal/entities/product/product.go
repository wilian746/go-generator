package product

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func ConvertIoReaderToStruct(data io.Reader) (body Product, err error) {
	if data == nil {
		return body, errors.New("body is invalid")
	}
	err = json.NewDecoder(data).Decode(&body)
	return body, err
}

func (p *Product) TableName() string {
	return "products"
}

