package services

/*
import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-api/domain"
	"testing"
)


import (
	"errors"
	"go-api/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAmenityByID(t *testing.T) {
	// Preparación
	amenities = []domain.Amenity{
		{ID: 1, Name: "Amenity 1"},
		{ID: 2, Name: "Amenity 2"},
		{ID: 3, Name: "Amenity 3"},
	}

	// Prueba exitosa
	amenity, err := GetAmenityByID(2)
	assert.Nil(t, err)
	assert.Equal(t, "Amenity 2", amenity.Name)

	// Prueba de amenidad no encontrada
	_, err = GetAmenityByID(4)
	assert.Equal(t, domain.ErrAmenityNotFound, err)
}

func TestCreateAmenity(t *testing.T) {
	// Preparación
	amenities = []domain.Amenity{
		{ID: 1, Name: "Amenity 1"},
		{ID: 2, Name: "Amenity 2"},
	}

	// Prueba exitosa
	newAmenity := domain.Amenity{Name: "Amenity 3"}
	createdAmenity, err := CreateAmenity(newAmenity)
	assert.Nil(t, err)
	assert.Equal(t, "Amenity 3", createdAmenity.Name)
	assert.Equal(t, 3, createdAmenity.ID)
	assert.Equal(t, 3, len(amenities))

	// Prueba de amenidad duplicada
	duplicateAmenity := domain.Amenity{Name: "Amenity 2"}
	_, err = CreateAmenity(duplicateAmenity)
	assert.Equal(t, errors.New("amenity already exists"), err)
}

func TestUpdateAmenity(t *testing.T) {
	// Preparación
	amenities = []domain.Amenity{
		{ID: 1, Name: "Amenity 1"},
		{ID: 2, Name: "Amenity 2"},
	}

	// Prueba exitosa
	updatedAmenity := domain.Amenity{ID: 2, Name: "Updated Amenity"}
	err := UpdateAmenity(updatedAmenity)
	assert.Nil(t, err)

	// Verificar que la amenidad haya sido actualizada
	amenity, _ := GetAmenityByID(2)
	assert.Equal(t, "Updated Amenity", amenity.Name)

	// Prueba de amenidad no encontrada
	err = UpdateAmenity(domain.Amenity{ID: 3, Name: "New Amenity"})
	assert.Equal(t, domain.ErrAmenityNotFound, err)
}

func TestDeleteAmenity(t *testing.T) {
	// Preparación
	amenities = []domain.Amenity{
		{ID: 1, Name: "Amenity 1"},
		{ID: 2, Name: "Amenity 2"},
	}

	// Prueba exitosa
	err := DeleteAmenity(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(amenities))

	// Verificar que la amenidad haya sido eliminada
	_, err = GetAmenityByID(1)
	assert.Equal(t, domain.ErrAmenityNotFound, err)

	// Prueba de amenidad no encontrada
	err = DeleteAmenity(3)
	assert.Equal(t, domain.ErrAmenityNotFound, err)
}*/
