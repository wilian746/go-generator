package adapter

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/database"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/entities"
	"github.com/wilian746/go-generator/pkg/standart-gorm/pkg/repository/response"
	"testing"
	"time"
)

type Product struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (p *Product) tableName() string {
	return "products"
}

func GetMemoryDatabase(isTransaction bool) (*Product, *gorm.DB) {
	product := &Product{}
	connection := database.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(product)
	connection.Where("id != 0").Delete(product)

	if !isTransaction {
		return product, connection
	}

	return product, connection.Begin()
}

func TestMock(t *testing.T) {
	mock := &Mock{}
	product := &Product{}

	t.Run("should return expected value in mock of GetLogMode", func(t *testing.T) {
		mock.On("GetLogMode").Return(true)
		assert.Equal(t, mock.GetLogMode(), true)
	})
	t.Run("should return expected value in mock of SetLogMode", func(t *testing.T) {
		mock.On("SetLogMode").Return(mock)
		assert.Equal(t, mock.SetLogMode(true), mock)
	})
	t.Run("should return expected value in mock of Connection", func(t *testing.T) {
		mock.On("Connection").Return(&gorm.DB{})
		assert.Equal(t, mock.Connection("product"), &gorm.DB{})
	})
	t.Run("should return expected value in mock of ParseGormQueryToDefaultResponse", func(t *testing.T) {
		mock.On("ParseGormQueryToDefaultResponse").Return(&response.Response{})
		assert.Equal(t, mock.ParseGormQueryToDefaultResponse(&gorm.DB{}), &response.Response{})
	})
	t.Run("should return expected value in mock of StartTransaction", func(t *testing.T) {
		mock.On("StartTransaction").Return(mock)
		assert.Equal(t, mock.StartTransaction(), mock)
	})
	t.Run("should return expected value in mock of CommitTransaction", func(t *testing.T) {
		mock.On("CommitTransaction").Return(&response.Response{})
		assert.Equal(t, mock.CommitTransaction(), &response.Response{})
	})
	t.Run("should return expected value in mock of RollbackTransaction", func(t *testing.T) {
		mock.On("RollbackTransaction").Return(&response.Response{})
		assert.Equal(t, mock.RollbackTransaction(), &response.Response{})
	})
	t.Run("should return expected value in mock of Find", func(t *testing.T) {
		mock.On("Find").Return(&response.Response{})
		assert.Equal(t, mock.Find(&gorm.DB{}, product, product.tableName()), &response.Response{})
	})
	t.Run("should return expected value in mock of Create", func(t *testing.T) {
		mock.On("Create").Return(&response.Response{})
		assert.Equal(t, mock.Create(product, product.tableName()), &response.Response{})
	})
	t.Run("should return expected value in mock of Update", func(t *testing.T) {
		mock.On("Update").Return(&response.Response{})
		assert.Equal(t, mock.Update(map[string]interface{}{}, product, product.tableName()), &response.Response{})
	})
	t.Run("should return expected value in mock of Delete", func(t *testing.T) {
		mock.On("Delete").Return(&response.Response{})
		assert.Equal(t, mock.Delete(map[string]interface{}{}, product.tableName()), &response.Response{})
	})
	t.Run("should return expected value in mock of Delete", func(t *testing.T) {
		mock.On("Health").Return(true)
		assert.Equal(t, mock.Health(), true)
	})
}

func TestNewAdapter(t *testing.T) {
	t.Run("should return type of Database", func(t *testing.T) {
		assert.IsType(t, &Database{}, NewAdapter(&gorm.DB{}))
	})
}

func TestDatabase_SetLogMode_GetLogMode(t *testing.T) {
	t.Run("should setLog correctly", func(t *testing.T) {
		connection := database.GetConnection("sqlite3", ":memory:")
		adapter := NewAdapter(connection)
		adapter.SetLogMode(true)
		assert.True(t, adapter.GetLogMode())
		adapter.SetLogMode(false)
		assert.False(t, adapter.GetLogMode())
	})
}

func TestDatabase_Connection(t *testing.T) {
	product := &Product{}
	connection := database.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(&Product{})

	t.Run("should return connection of type *gorm.DB", func(t *testing.T) {
		adapter := NewAdapter(connection)
		assert.IsType(t, &gorm.DB{}, adapter.Connection(product.tableName()))
	})
}

func TestDatabase_ParseGormQueryToDefaultResponse(t *testing.T) {
	product := &Product{}
	connection := database.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(&Product{})

	t.Run("should return response of type *response.Response", func(t *testing.T) {
		adapter := NewAdapter(connection)
		connection := adapter.Connection(product.tableName())

		assert.IsType(t, &response.Response{}, adapter.ParseGormQueryToDefaultResponse(connection))
	})
}

func TestDatabase_Health(t *testing.T) {
	product := &Product{}
	connection := database.GetConnection("sqlite3", ":memory:")
	connection.Table(product.tableName()).AutoMigrate(&Product{})

	t.Run("should return response of type *response.Response", func(t *testing.T) {
		adapter := NewAdapter(connection)
		adapter.Connection(product.tableName())
		assert.True(t, adapter.Health())
	})
}

func Test_One_to_One(t *testing.T) {
	t.Run("Should run tests to execute Create, Read, Update, Delete with One To One", func(t *testing.T) {
		connection := database.GetConnection("sqlite3", ":memory:").LogMode(true)
		adapter := NewAdapter(connection).SetLogMode(true)
		firstNameToCreateStudent := uuid.New().String()
		contact := &entities.Contact{
			ID:      uuid.New(),
			City:    uuid.New().String(),
			Phone:   uuid.New().String(),
			Address: uuid.New().String(),
		}
		student := &entities.Student{
			ID:        uuid.New(),
			FirstName: uuid.New().String(),
			LastName:  firstNameToCreateStudent,
			ContactID: contact.ID,
		}
		assert.NoError(t, connection.Table(student.TableName()).AutoMigrate(&entities.Student{}).Error)
		assert.NoError(t, connection.Table(contact.TableName()).AutoMigrate(&entities.Contact{}).Error)

		contact.SetCreatedAt()
		student.SetUpdatedAt()
		student.SetCreatedAt()
		contact.SetUpdatedAt()

		responseCreate := adapter.Create(contact, contact.TableName())
		assert.NoError(t, responseCreate.Error())
		responseCreate = adapter.Create(student, student.TableName())
		assert.NoError(t, responseCreate.Error())

		queryFindAll := adapter.Connection(student.TableName())
		entityToCheckFindAll := []entities.Student{}
		responseFindAll := adapter.Find(queryFindAll, &entityToCheckFindAll, student.TableName())
		assert.NoError(t, responseFindAll.Error())
		assert.NotEqual(t, len(entityToCheckFindAll), 0)

		queryFindOne := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1).
			Preload("Contact")
		entityToCheckFindOne := entities.Student{}
		responseFindOne := adapter.Find(queryFindOne, &entityToCheckFindOne, student.TableName())
		assert.NoError(t, responseFindOne.Error())
		assert.NotEqual(t, entityToCheckFindOne.ID, uuid.Nil)
		assert.NotEqual(t, entityToCheckFindOne.ContactID, uuid.Nil)
		assert.NotEqual(t, entityToCheckFindOne.Contact.ID, uuid.Nil)

		student.LastName = uuid.New().String()
		responseUpdate := adapter.Update(map[string]interface{}{"id": student.ID}, student, student.TableName())
		assert.NoError(t, responseUpdate.Error())

		queryFindOneUpdate := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneUpdate.Error)
		entityToCheckUpdate := entities.Student{}
		responseFindOneUpdate := adapter.Find(queryFindOneUpdate, &entityToCheckUpdate, student.TableName())
		assert.NoError(t, responseFindOneUpdate.Error())
		assert.NotEqual(t, entityToCheckUpdate.LastName, entityToCheckFindOne.LastName)

		responseDelete := adapter.Delete(map[string]interface{}{"id": student.ID}, student.TableName())
		assert.NoError(t, responseDelete.Error())
		responseDelete = adapter.Delete(map[string]interface{}{"id": contact.ID}, contact.TableName())
		assert.NoError(t, responseDelete.Error())

		queryFindOneDelete := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneDelete.Error)
		entityToDelete := entities.Student{}
		responseFindOneDelete := adapter.Find(queryFindOneDelete, &entityToDelete, student.TableName())
		assert.Error(t, responseFindOneDelete.Error())
		assert.Equal(t, responseFindOneDelete.Error().Error(), ErrRecordNotFound.Error())

	})
}

func Test_One_to_One_Transaction(t *testing.T) {
	t.Run("Should run tests to execute Create, Read, Update, Delete with One To One with transaction", func(t *testing.T) {
		connection := database.GetConnection("sqlite3", ":memory:").LogMode(true)
		adapter := NewAdapter(connection).SetLogMode(true)
		firstNameToCreateStudent := uuid.New().String()
		contact := &entities.Contact{
			ID:      uuid.New(),
			City:    uuid.New().String(),
			Phone:   uuid.New().String(),
			Address: uuid.New().String(),
		}
		student := &entities.Student{
			ID:        uuid.New(),
			FirstName: uuid.New().String(),
			LastName:  firstNameToCreateStudent,
			ContactID: contact.ID,
		}
		assert.NoError(t, connection.Table(student.TableName()).AutoMigrate(&entities.Student{}).Error)
		assert.NoError(t, connection.Table(contact.TableName()).AutoMigrate(&entities.Contact{}).Error)

		contact.SetCreatedAt()
		student.SetCreatedAt()

		adapter = adapter.StartTransaction()

		responseCreate := adapter.Create(contact, contact.TableName())
		assert.NoError(t, responseCreate.Error())
		if responseCreate.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		responseCreate = adapter.Create(student, student.TableName())
		assert.NoError(t, responseCreate.Error())
		if responseCreate.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}

		queryFindAll := adapter.Connection(student.TableName())
		entityToCheckFindAll := []entities.Student{}
		responseFindAll := adapter.Find(queryFindAll, &entityToCheckFindAll, student.TableName())
		assert.NoError(t, responseFindAll.Error())
		if responseFindAll.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		assert.NotEqual(t, len(entityToCheckFindAll), 0)

		queryFindOne := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1).
			Preload("Contact")
		assert.NoError(t, queryFindOne.Error)
		if queryFindOne.Error != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		entityToCheckFindOne := entities.Student{}
		responseFindOne := adapter.Find(queryFindOne, &entityToCheckFindOne, student.TableName())
		assert.NoError(t, responseFindOne.Error())
		if responseFindOne.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		assert.NotEqual(t, entityToCheckFindOne.ID, uuid.Nil)
		assert.NotEqual(t, entityToCheckFindOne.ContactID, uuid.Nil)
		assert.NotEqual(t, entityToCheckFindOne.Contact.ID, uuid.Nil)

		student.LastName = uuid.New().String()
		responseUpdate := adapter.Update(map[string]interface{}{"id": student.ID}, student, student.TableName())
		assert.NoError(t, responseUpdate.Error())
		if responseUpdate.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}

		queryFindOneUpdate := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneUpdate.Error)
		if queryFindOneUpdate.Error != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		entityToCheckUpdate := entities.Student{}
		responseFindOneUpdate := adapter.Find(queryFindOneUpdate, &entityToCheckUpdate, student.TableName())
		assert.NoError(t, responseFindOneUpdate.Error())
		if responseFindOneUpdate.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		assert.NotEqual(t, entityToCheckUpdate.LastName, entityToCheckFindOne.LastName)

		responseDelete := adapter.Delete(map[string]interface{}{"id": student.ID}, student.TableName())
		assert.NoError(t, responseDelete.Error())
		if responseDelete.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		responseDelete = adapter.Delete(map[string]interface{}{"id": contact.ID}, contact.TableName())
		assert.NoError(t, responseDelete.Error())
		if responseDelete.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}

		queryFindOneDelete := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneDelete.Error)
		if queryFindOneDelete.Error != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		entityToDelete := entities.Student{}
		responseFindOneDelete := adapter.Find(queryFindOneDelete, &entityToDelete, student.TableName())
		assert.Error(t, responseFindOneDelete.Error())
		assert.Equal(t, responseFindOneDelete.Error().Error(), ErrRecordNotFound.Error())
		if responseFindOneDelete.Error() == nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		} else {
			assert.NoError(t, adapter.CommitTransaction().Error())
		}
	})
	t.Run("Should run tests to execute Create, Read, Update, Delete with One To One with transaction and retur error on preloads", func(t *testing.T) {
		connection := database.GetConnection("sqlite3", ":memory:").LogMode(true)
		adapter := NewAdapter(connection).SetLogMode(true)
		firstNameToCreateStudent := uuid.New().String()
		contact := &entities.Contact{
			ID:      uuid.New(),
			City:    uuid.New().String(),
			Phone:   uuid.New().String(),
			Address: uuid.New().String(),
		}
		student := &entities.Student{
			ID:        uuid.New(),
			FirstName: uuid.New().String(),
			LastName:  firstNameToCreateStudent,
			ContactID: contact.ID,
		}
		assert.NoError(t, connection.Table(student.TableName()).AutoMigrate(&entities.Student{}).Error)
		assert.NoError(t, connection.Table(contact.TableName()).AutoMigrate(&entities.Contact{}).Error)

		contact.SetCreatedAt()
		student.SetCreatedAt()

		adapter = adapter.StartTransaction()

		responseCreate := adapter.Create(contact, contact.TableName())
		assert.NoError(t, responseCreate.Error())
		if responseCreate.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		responseCreate = adapter.Create(student, student.TableName())
		assert.NoError(t, responseCreate.Error())
		if responseCreate.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}

		queryFindOne := adapter.
			Connection(student.TableName()).
			Where(map[string]interface{}{"id": student.ID}).
			Limit(1).
			Preload("Contact")
		assert.NoError(t, queryFindOne.Error)
		if queryFindOne.Error != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
		entityToCheckFindOne := entities.Student{}
		responseFindOne := adapter.Find(queryFindOne, &entityToCheckFindOne, "Not existing table")
		assert.Error(t, responseFindOne.Error())
		if responseFindOne.Error() != nil {
			assert.NoError(t, adapter.RollbackTransaction().Error())
		}
	})
}

func Test_One_to_Many(t *testing.T) {
	t.Run("Should run tests to execute Create, Read, Update, Delete with One To Many", func(t *testing.T) {
		connection := database.GetConnection("sqlite3", ":memory:").LogMode(true)
		adapter := NewAdapter(connection).SetLogMode(true)
		nameToCreate := uuid.New().String()
		restaurant := &entities.Restaurant{
			ID:   uuid.New(),
			Name: nameToCreate,
		}
		order := &entities.Order{
			ID:           uuid.New(),
			Description:  uuid.New().String(),
			RestaurantID: restaurant.ID,
		}
		assert.NoError(t, connection.Table(restaurant.TableName()).AutoMigrate(&entities.Restaurant{}).Error)
		assert.NoError(t, connection.Table(order.TableName()).AutoMigrate(&entities.Order{}).Error)

		restaurant.SetCreatedAt()
		order.SetUpdatedAt()
		order.SetCreatedAt()
		restaurant.SetUpdatedAt()

		responseCreate := adapter.Create(restaurant, restaurant.TableName())
		assert.NoError(t, responseCreate.Error())
		responseCreate = adapter.Create(order, order.TableName())
		assert.NoError(t, responseCreate.Error())

		queryFindAll := adapter.Connection(restaurant.TableName())
		entityToCheckFindAll := []entities.Restaurant{}
		responseFindAll := adapter.Find(queryFindAll, &entityToCheckFindAll, restaurant.TableName())
		assert.NoError(t, responseFindAll.Error())
		assert.NotEqual(t, len(entityToCheckFindAll), 0)

		queryFindOne := adapter.
			Connection(restaurant.TableName()).
			Where(map[string]interface{}{"id": restaurant.ID}).
			Limit(1).
			Preload("Orders")
		entityToCheckFindOne := entities.Restaurant{}
		responseFindOne := adapter.Find(queryFindOne, &entityToCheckFindOne, restaurant.TableName())
		assert.NoError(t, responseFindOne.Error())
		assert.NotEqual(t, entityToCheckFindOne.ID, uuid.Nil)
		assert.NotEqual(t, len(entityToCheckFindOne.Orders), 0)

		restaurant.Name = uuid.New().String()
		responseUpdate := adapter.Update(map[string]interface{}{"id": restaurant.ID}, restaurant, restaurant.TableName())
		assert.NoError(t, responseUpdate.Error())

		queryFindOneUpdate := adapter.
			Connection(restaurant.TableName()).
			Where(map[string]interface{}{"id": restaurant.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneUpdate.Error)
		entityToCheckUpdate := entities.Restaurant{}
		responseFindOneUpdate := adapter.Find(queryFindOneUpdate, &entityToCheckUpdate, restaurant.TableName())
		assert.NoError(t, responseFindOneUpdate.Error())
		assert.NotEqual(t, entityToCheckUpdate.Name, entityToCheckFindOne.Name)

		responseDelete := adapter.Delete(map[string]interface{}{"id": restaurant.ID}, restaurant.TableName())
		assert.NoError(t, responseDelete.Error())
		responseDelete = adapter.Delete(map[string]interface{}{"id": order.ID}, order.TableName())
		assert.NoError(t, responseDelete.Error())

		queryFindOneDelete := adapter.
			Connection(restaurant.TableName()).
			Where(map[string]interface{}{"id": restaurant.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneDelete.Error)
		entityToDelete := entities.Restaurant{}
		responseFindOneDelete := adapter.Find(queryFindOneDelete, &entityToDelete, restaurant.TableName())
		assert.Error(t, responseFindOneDelete.Error())
		assert.Equal(t, responseFindOneDelete.Error().Error(), ErrRecordNotFound.Error())

	})
}

func Test_Many_to_Many(t *testing.T) {
	t.Run("Should run tests to execute Create, Read, Update, Delete with Many To Many", func(t *testing.T) {
		connection := database.GetConnection("sqlite3", ":memory:").LogMode(true)
		adapter := NewAdapter(connection).SetLogMode(true)
		nameToCreateDoctor := uuid.New().String()
		nameToCreatePatient := uuid.New().String()
		doctor := &entities.Doctor{
			ID:   uuid.New(),
			Name: nameToCreateDoctor,
		}
		patient := &entities.Patient{
			ID:   uuid.New(),
			Name: nameToCreatePatient,
		}
		doctorPatient := &entities.DoctorPatient{
			ID:        uuid.New(),
			PatientID: patient.ID,
			DoctorID:  doctor.ID,
		}
		assert.NoError(t, connection.Table(doctor.TableName()).AutoMigrate(&entities.Doctor{}).Error)
		assert.NoError(t, connection.Table(patient.TableName()).AutoMigrate(&entities.Patient{}).Error)
		assert.NoError(t, connection.Table(doctorPatient.TableName()).AutoMigrate(&entities.DoctorPatient{}).Error)

		doctor.SetCreatedAt()
		doctor.SetUpdatedAt()
		patient.SetCreatedAt()
		patient.SetUpdatedAt()
		doctorPatient.SetCreatedAt()
		doctorPatient.SetUpdatedAt()

		responseCreate := adapter.Create(doctor, doctor.TableName())
		assert.NoError(t, responseCreate.Error())
		responseCreate = adapter.Create(patient, patient.TableName())
		assert.NoError(t, responseCreate.Error())
		responseCreate = adapter.Create(doctorPatient, doctorPatient.TableName())
		assert.NoError(t, responseCreate.Error())

		queryFindAll := adapter.Connection(doctor.TableName())
		entityToCheckFindAll := []entities.Doctor{}
		responseFindAll := adapter.Find(queryFindAll, &entityToCheckFindAll, doctor.TableName())
		assert.NoError(t, responseFindAll.Error())
		assert.NotEqual(t, len(entityToCheckFindAll), 0)

		queryFindOne := adapter.
			Connection(doctor.TableName()).
			Where(map[string]interface{}{"id": doctor.ID}).
			Limit(1).
			Preload("DoctorsPatients").
			Preload("DoctorsPatients.Patient").
			Preload("DoctorsPatients.Doctor")
		entityToCheckFindOne := entities.Doctor{}
		responseFindOne := adapter.Find(queryFindOne, &entityToCheckFindOne, doctor.TableName())
		assert.NoError(t, responseFindOne.Error())
		assert.NotEqual(t, entityToCheckFindOne.ID, uuid.Nil)
		assert.NotEqual(t, len(entityToCheckFindOne.DoctorsPatients), 0)
		if len(entityToCheckFindOne.DoctorsPatients) > 0 {
			assert.NotEqual(t, entityToCheckFindOne.DoctorsPatients[0].Doctor.ID, uuid.Nil)
		}

		doctor.Name = uuid.New().String()
		responseUpdate := adapter.Update(map[string]interface{}{"id": doctor.ID}, doctor, doctor.TableName())
		assert.NoError(t, responseUpdate.Error())

		queryFindOneUpdate := adapter.
			Connection(doctor.TableName()).
			Where(map[string]interface{}{"id": doctor.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneUpdate.Error)
		entityToCheckUpdate := entities.Doctor{}
		responseFindOneUpdate := adapter.Find(queryFindOneUpdate, &entityToCheckUpdate, doctor.TableName())
		assert.NoError(t, responseFindOneUpdate.Error())
		assert.NotEqual(t, entityToCheckUpdate.Name, entityToCheckFindOne.Name)

		responseDelete := adapter.Delete(map[string]interface{}{"id": doctor.ID}, doctor.TableName())
		assert.NoError(t, responseDelete.Error())
		responseDelete = adapter.Delete(map[string]interface{}{"id": doctor.ID}, doctor.TableName())
		assert.NoError(t, responseDelete.Error())

		queryFindOneDelete := adapter.
			Connection(doctor.TableName()).
			Where(map[string]interface{}{"id": doctor.ID}).
			Limit(1)
		assert.NoError(t, queryFindOneDelete.Error)
		entityToDelete := entities.Doctor{}
		responseFindOneDelete := adapter.Find(queryFindOneDelete, &entityToDelete, doctor.TableName())
		assert.Error(t, responseFindOneDelete.Error())
		assert.Equal(t, responseFindOneDelete.Error().Error(), ErrRecordNotFound.Error())

	})
}
