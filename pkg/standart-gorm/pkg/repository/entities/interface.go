package entities

import (
	"github.com/google/uuid"
	"time"
)

type Interface interface {
	TableName() string
	GenerateID()
	Bytes() []byte
	SetCreatedAt()
	SetUpdatedAt()
}

/*
======================== MANY_TO_MANY ========================
*/

type Doctor struct {
	Interface
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DoctorsPatients []DoctorPatient `gorm:"foreignkey:DoctorID;association_foreignkey:ID"`
}

func (d *Doctor) TableName() string {
	return "doctors"
}

func (d *Doctor) SetCreatedAt() {
	d.CreatedAt = time.Now()
}

func (d *Doctor) SetUpdatedAt() {
	d.UpdatedAt = time.Now()
}

type DoctorPatient struct {
	Interface
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	PatientID uuid.UUID `sql:"type:uuid REFERENCES patients(id) ON DELETE CASCADE"`
	DoctorID  uuid.UUID `sql:"type:uuid REFERENCES doctors(id) ON DELETE CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Patient   Patient `gorm:"foreignkey:PatientID;association_foreignkey:ID"`
	Doctor    Doctor  `gorm:"foreignkey:DoctorID;association_foreignkey:ID"`
}

func (dp *DoctorPatient) TableName() string {
	return "doctors_patients"
}

func (dp *DoctorPatient) SetCreatedAt() {
	dp.CreatedAt = time.Now()
}

func (dp *DoctorPatient) SetUpdatedAt() {
	dp.UpdatedAt = time.Now()
}

type Patient struct {
	Interface
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DoctorsPatients []DoctorPatient `gorm:"foreignkey:PatientID;association_foreignkey:ID"`
}

func (p *Patient) TableName() string {
	return "patients"
}

func (p *Patient) SetCreatedAt() {
	p.CreatedAt = time.Now()
}

func (p *Patient) SetUpdatedAt() {
	p.UpdatedAt = time.Now()
}

/*
======================== ONE_TO_MANY ========================
*/

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

/*
======================== ONE_TO_ONE ========================
*/

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
