package entities

type Interface interface {
	TableName() string
	GenerateID()
	SetCreatedAt()
	SetUpdatedAt()
}
