package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockDB struct{}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	return nil
}

func (m *MockDB) Model(value interface{}) *gorm.DB {
	return nil
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return nil
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return nil
}

func (m *MockDB) Count(count *int64, conds ...interface{}) *gorm.DB {
	*count = 5 // Valor de ejemplo
	return nil
}

func (m *MockDB) Session(dbSession *gorm.Session) *gorm.DB {
	return nil
}

func (m *MockDB) Preload(column string, conditions ...interface{}) *gorm.DB {
	return nil
}

func (m *MockDB) Order(value interface{}, reorder ...bool) *gorm.DB {
	return nil
}

func (m *MockDB) Limit(limit int) *gorm.DB {
	return nil
}

func (m *MockDB) Offset(offset int) *gorm.DB {
	return nil
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return nil
}

func TestCreateReservation(t *testing.T) {
	// Configurar el entorno de prueba
	db := &MockDB{}
	token := "valid-token"
	hotelID := 1
	checkin := "2023-06-01"
	checkout := "2023-06-03"
	clientName := "John Doe"

	// Ejecutar la función a probar
	err := CreateReservation(db, hotelID, checkin, checkout, token, clientName)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error creating reservation should be nil")
}

func TestCalculateAvailableRooms(t *testing.T) {
	// Configurar el entorno de prueba
	db := &MockDB{}
	hotelID := 1
	checkin := "2023-06-01"
	checkout := "2023-06-03"

	// Ejecutar la función a probar
	availableRooms, err := calculateAvailableRooms(db, hotelID, checkin, checkout)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error calculating available rooms should be nil")
	assert.Equal(t, 5, availableRooms, "available rooms should be 5")
}

func TestGetTotalRoomsFromDB(t *testing.T) {
	// Configurar el entorno de prueba
	db := &MockDB{}
	hotelID := 1

	// Ejecutar la función a probar
	totalRooms, err := getTotalRoomsFromDB(db, hotelID)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error getting total rooms should be nil")
	assert.Equal(t, 5, totalRooms, "total rooms should be 5")
}

func TestGetReservedRoomsFromDB(t *testing.T) {
	// Configurar el entorno de prueba
	db := &MockDB{}
	hotelID := 1
	checkin := "2023-06-01"
	checkout := "2023-06-03"

	// Ejecutar la función a probar
	reservedRooms, err := getReservedRoomsFromDB(db, hotelID, checkin, checkout)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error getting reserved rooms should be nil")
	assert.Equal(t, 5, reservedRooms, "reserved rooms should be 5")
}
