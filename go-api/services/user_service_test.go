package services

import (
	"errors"
	"go-api/clients"
	"go-api/dao"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	// Configurar el entorno de prueba
	userID := 1
	expectedUser := &dao.User{
		ID:   1,
		Name: "John Doe",
	}

	mockDBClient := &clients.DBClientMock{
		GetUserByIDFunc: func(userID int) (*dao.User, error) {
			// Simular la recuperación de un usuario por ID
			if userID == 1 {
				user := dao.User{
					ID:   1,
					Name: "John Doe",
				}
				return &user, nil
			}

			// Simular el caso en el que no se encuentre el usuario
			return nil, errors.New("user not found")
		},
	}

	service := UserService{
		DBClient: mockDBClient,
	}

	// Ejecutar la función a probar
	user, err := service.GetUserByID(userID)

	// Verificar el resultado esperado
	assert.NoError(t, err, "error getting user should be nil")
	assert.Equal(t, expectedUser, user, "retrieved user should match expected user")
}

func TestGetUserByEmail(t *testing.T) {
	// Create the UserService with the DBClientMock
	dbClientMock := &clients.DBClientMock{}
	userService := NewUserService(dbClientMock)

	// Call the GetUserByEmail function
	user, err := userService.GetUserByEmail("john.doe@example.com")

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 1, user.ID, "Incorrect user ID")
	assert.Equal(t, "John Doe", user.Name, "Incorrect user name")
	assert.Equal(t, "john.doe@example.com", user.Email, "Incorrect user email")
	assert.Equal(t, "admin", user.Role, "Incorrect user role")
}

func TestCreateUser(t *testing.T) {
	// Create the UserService with the DBClientMock
	dbClientMock := &clients.DBClientMock{}
	userService := NewUserService(dbClientMock)

	// Create a new user
	newUser := &dao.User{
		Name:     "Jane Smith",
		Email:    "jane.smith@example.com",
		Password: "password123",
	}

	// Call the CreateUser function
	createdUser, err := userService.CreateUser(newUser)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 1, createdUser.ID, "Incorrect user ID")
	assert.Equal(t, "Jane Smith", createdUser.Name, "Incorrect user name")
	assert.Equal(t, "jane.smith@example.com", createdUser.Email, "Incorrect user email")
	assert.Equal(t, "admin", createdUser.Role, "Incorrect user role")
}
