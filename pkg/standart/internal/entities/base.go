package entities

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) GenerateID () {
	b.ID = uuid.New()
}

func (b *Base) SetCreatedAt () {
	b.CreatedAt = time.Now()
}

func (b *Base) SetUpdatedAt () {
	b.UpdatedAt = time.Now()
}
