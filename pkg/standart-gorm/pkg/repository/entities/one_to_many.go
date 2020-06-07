package entities

import (
	"github.com/google/uuid"
	"time"
)

type Restaurant struct {
	Interface
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Orders    []Order `gorm:"foreignkey:RestaurantID;association_foreignkey:ID"`
}

func (r *Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) SetCreatedAt() {
	r.CreatedAt = time.Now()
}

func (r *Restaurant) SetUpdatedAt() {
	r.UpdatedAt = time.Now()
}

type Order struct {
	Interface
	ID           uuid.UUID `gorm:"type:uuid;primary_key;"`
	Description  string
	RestaurantID uuid.UUID `sql:"type:uuid REFERENCES restaurants(id) ON DELETE CASCADE"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) SetCreatedAt() {
	o.CreatedAt = time.Now()
}

func (o *Order) SetUpdatedAt() {
	o.UpdatedAt = time.Now()
}
