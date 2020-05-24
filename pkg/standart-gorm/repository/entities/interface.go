package entities

type Interface interface {
	TableName() string
	GenerateID()
	Bytes() []byte
	SetCreatedAt()
	SetUpdatedAt()
}
