package services

import (
	"errors"
	"go-api/domain"

	"gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// GetHotelByID obtiene los datos de un hotel por su ID
func GetHotelByID(id int) (*domain.Hotel, error) {
	// Implementar la lógica para obtener los datos del hotel por su ID

	// Realiza la consulta a la base de datos para obtener el hotel por su ID utilizando tu sistema de almacenamiento correspondiente
	// Por ejemplo, si estás utilizando GORM como ORM, podrías hacer algo como esto:
	var hotel domain.Hotel
	if err := database.DB.First(&hotel, id).Error; err != nil {
		// Ocurrió un error al obtener el hotel de la base de datos
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Si el hotel no existe, retornar nil en lugar de un error
			return nil, nil
		}
		return nil, err
	}

	// Retorna el hotel encontrado
	return &hotel, nil
}

// CreateHotel crea un nuevo hotel
func CreateHotel(newHotel domain.Hotel) (domain.Hotel, error) {
	// Realiza la validación de los datos del nuevo hotel según tus requisitos
	if len(newHotel.Name) == 0 {
		return domain.Hotel{}, errors.New("El nombre del hotel es obligatorio")
	}

	if len(newHotel.Name) > 255 {
		return domain.Hotel{}, errors.New("El nombre del hotel excede la longitud máxima permitida")
	}

	// Verifica la unicidad del nombre del hotel en la base de datos
	existingHotel, err := GetHotelByName(newHotel.Name)
	if err != nil {
		return domain.Hotel{}, err
	}

	if existingHotel != nil {
		return domain.Hotel{}, errors.New("Ya existe un hotel con el mismo nombre")
	}

	// Aquí se podría realizar la lógica de validación y persistencia del hotel en la base de datos

	// Persiste el nuevo hotel en la base de datos utilizando tu sistema de almacenamiento correspondiente
	err = database.CreateHotel(&newHotel)
	if err != nil {
		return domain.Hotel{}, err
	}

// Creación exitosa del hotel en la base de datos

	// Crea el objeto createdHotel a partir de los datos ingresados
	createdHotel := domain.Hotel{
		ID:          newHotel.ID,
		Name:        newHotel.Name,
		Photo:       newHotel.Photo,
		Description: newHotel.Description,
		Location:    newHotel.Location,
		Rooms:       newHotel.Rooms,
	}

	// Devuelve el hotel creado
	return createdHotel, nil
}

// UpdateHotel actualiza los datos de un hotel existente
func UpdateHotel(id int, updatedHotel domain.Hotel) (domain.Hotel, error) {
	// Aquí se podría realizar la lógica de validación y actualización del hotel en la base de datos

	// Realiza una consulta a la base de datos para obtener el hotel existente por su ID
	existingHotel, err := database.GetHotelByID(id)
	if err != nil {
		// Ocurrió un error al consultar la base de datos
		return domain.Hotel{}, err
	}

	// Si el hotel no existe, retorna un error
	if existingHotel == nil {
		return domain.Hotel{}, errors.New("El hotel no existe")
	}

	// Realiza la actualización de los campos necesarios del hotel existente con los valores proporcionados
	existingHotel.Name = updatedHotel.Name
	existingHotel.Photo = updatedHotel.Photo
	existingHotel.Description = updatedHotel.Description
	existingHotel.Location = updatedHotel.Location
	existingHotel.Rooms = updatedHotel.Rooms

	// Realiza la actualización del hotel en la base de datos utilizando tu sistema de almacenamiento correspondiente
	err = database.UpdateHotel(existingHotel)
	if err != nil {
		// Ocurrió un error al actualizar el hotel en la base de datos
		return domain.Hotel{}, err
	}

	// La actualización del hotel se realizó correctamente
	return existingHotel, nil
}

// DeleteHotel elimina un hotel existente
func DeleteHotel(id int) error {
	// Aquí se podría realizar la lógica de validación y eliminación del hotel en la base de datos

	// Realiza una consulta a la base de datos para verificar si el hotel existe
	existingHotel, err := database.GetHotelByID(id)
	if err != nil {
		// Ocurrió un error al consultar la base de datos
		return err
	}

	// Si el hotel no existe, retorna un error
	if existingHotel == nil {
		return errors.New("El hotel no existe")
	}

	// Si el hotel existe, procede con la eliminación en la base de datos
	err = database.DeleteHotel(id)
	if err != nil {
		// Ocurrió un error al eliminar el hotel en la base de datos
		return err
	}

	// La eliminación del hotel se realizó correctamente
	return nil
}
