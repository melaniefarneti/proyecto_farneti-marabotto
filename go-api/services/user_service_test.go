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
