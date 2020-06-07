package entities

import (
	"github.com/google/uuid"
	"time"
)

type Student struct {
	Interface
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	FirstName string
	LastName  string
	ContactID uuid.UUID `sql:"type:uuid REFERENCES contacts(id) ON DELETE CASCADE"`
	Contact   Contact   `gorm:"foreignkey:ContactID;association_foreignkey:ID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Student) TableName() string {
	return "students"
}

func (s *Student) SetCreatedAt() {
	s.CreatedAt = time.Now()
}

func (s *Student) SetUpdatedAt() {
	s.UpdatedAt = time.Now()
}

type Contact struct {
	Interface
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	City      string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Contact) TableName() string {
	return "contacts"
}
func (c *Contact) SetCreatedAt() {
	c.CreatedAt = time.Now()
}
func (c *Contact) SetUpdatedAt() {
	c.UpdatedAt = time.Now()
}
