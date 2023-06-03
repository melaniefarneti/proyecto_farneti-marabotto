package services

import (
	"errors"
	"go-api/domain"
)

var users []domain.User

// GetUserByID obtiene los datos de un usuario por su ID
func GetUserByID(id int) (domain.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return domain.User{}, domain.ErrUserNotFound
}

// CreateUser crea un nuevo usuario
func CreateUser(user domain.User) (domain.User, error) {
	// Verificar si el usuario ya existe
	for _, u := range users {
		if u.Email == user.Email {
			return domain.User{}, errors.New("user already exists")
		}
	}

	// Asignar ID único al usuario
	user.ID = generateUserID()

	// Agregar usuario a la lista
	users = append(users, user)

	return user, nil
}

// generateUserID genera un ID único para un usuario
func generateUserID() int {
	// Implementar la lógica para generar un ID único para el usuario, por ejemplo, incrementando un contador o utilizando un generador de IDs único
	// generamos un ID incrementando un contador
	return len(users) + 1
}
