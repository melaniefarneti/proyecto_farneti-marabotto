package services

/*
import (
	"errors"
	"go-api/domain"
)

var amenities []domain.Amenity

// GetAmenityByID obtiene los datos de una amenidad por su ID
func GetAmenityByID(id int) (domain.Amenity, error) {
	for _, amenity := range amenities {
		if amenity.ID == id {
			return amenity, nil
		}
	}
	return domain.Amenity{}, domain.ErrAmenityNotFound
}


// CreateAmenity crea una nueva amenidad
func CreateAmenity(amenity domain.Amenity) (domain.Amenity, error) {
	for _, existingAmenity := range amenities {
		if existingAmenity.Name == amenity.Name {
			return domain.Amenity{}, errors.New("amenity already exists")
		}
	}

	amenity.ID = generateAmenityID()
	amenities = append(amenities, amenity)

	return amenity, nil
}

// UpdateAmenity actualiza una amenidad existente
func UpdateAmenity(amenity domain.Amenity) error {
	for i, existingAmenity := range amenities {
		if existingAmenity.ID == amenity.ID {
			amenities[i] = amenity
			return nil
		}
	}
	return domain.ErrAmenityNotFound
}

// DeleteAmenity elimina una amenidad existente
func DeleteAmenity(id int) error {
	for i, amenity := range amenities {
		if amenity.ID == id {
			amenities = append(amenities[:i], amenities[i+1:]...)
			return nil
		}
	}
	return domain.ErrAmenityNotFound
}

// generateAmenityID genera un ID único para una amenidad
func generateAmenityID() int {
	// Implementar la lógica para generar un ID único para la amenidad
	// incrementando un contador
	return len(amenities) + 1
}
*/
