package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go-api/domain"
	"go-api/services"
)

// TestGetHotelByID testea la función GetHotelByID
func TestGetHotelByID(t *testing.T) {
	// Crear un hotel de prueba en la base de datos para obtener su ID
	testHotel := domain.Hotel{
		ID:          1,
		Name:        "Hotel de Prueba",
		Photo:       "photo.png",
		Description: "Descripción del hotel",
		Location:    "Ubicación del hotel",
		Rooms:       "Habitaciones disponibles",
	}
	createdHotel, err := services.CreateHotel(testHotel)
	assert.Nil(t, err)

	// Prueba
	retrievedHotel, err := services.GetHotelByID(createdHotel.ID)
	assert.Nil(t, err)
	assert.Equal(t, createdHotel, retrievedHotel)
}

// TestCreateHotel testea la función CreateHotel
func TestCreateHotel(t *testing.T) {
	// Preparación
	testHotel := domain.Hotel{
		ID:          2,
		Name:        "Nuevo Hotel",
		Photo:       "photo.png",
		Description: "Descripción del nuevo hotel",
		Location:    "Ubicación del nuevo hotel",
		Rooms:       "Habitaciones del nuevo hotel",
	}

	// Prueba
	createdHotel, err := services.CreateHotel(testHotel)
	assert.Nil(t, err)
	assert.Equal(t, testHotel, createdHotel)
}

// TestUpdateHotel testea la función UpdateHotel
func TestUpdateHotel(t *testing.T) {
	// Preparación
	testHotel := domain.Hotel{
		ID:          3,
		Name:        "Hotel a actualizar",
		Photo:       "photo.png",
		Description: "Descripción del hotel a actualizar",
		Location:    "Ubicación del hotel a actualizar",
		Rooms:       "Habitaciones del hotel a actualizar",
	}
	createdHotel, err := services.CreateHotel(testHotel)
	assert.Nil(t, err)

	// Actualización de los datos del hotel
	createdHotel.Name = "Hotel actualizado"
	createdHotel.Description = "Descripción actualizada"
	createdHotel.Location = "Ubicación actualizada"
	createdHotel.Rooms = "Habitaciones actualizadas"

	// Prueba
	updatedHotel, err := services.UpdateHotel(createdHotel)
	assert.Nil(t, err)
	assert.Equal(t, createdHotel, updatedHotel)
}

// TestDeleteHotel testea la función DeleteHotel
func TestDeleteHotel(t *testing.T) {
	// Preparación
	testHotel := domain.Hotel{
		ID:          4,
		Name:        "Hotel a eliminar",
		Photo:       "photo.png",
		Description: "Descripción del hotel a eliminar",
		Location:    "Ubicación del hotel a eliminar",
		Rooms:       "Habitaciones del hotel a eliminar",
	}
	createdHotel, err := services.CreateHotel(testHotel)
	assert.Nil(t, err)

	// Prueba
	err = services.DeleteHotel(createdHotel.ID)
	assert.Nil(t, err)

	// Verificar que el hotel haya sido eliminado
	_, err = services.GetHotelByID(createdHotel.ID)
	assert.NotNil(t, err)
}
