package services

import (
	"testing"

	"go-api/dao"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// Configurar el entorno de prueba
	users = []dao.User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Smith"},
		{ID: 3, Name: "Alice Johnson"},
	}

	// Caso de prueba: Usuario existente
	expectedUser := dao.User{ID: 2, Name: "Jane Smith"}
	user, err := GetUser(2)
	assert.NoError(t, err, "error should be nil")
	assert.Equal(t, expectedUser, user, "user data does not match")

	// Caso de prueba: Usuario no encontrado
	_, err = GetUser(4)
	assert.ErrorIs(t, err, dao.ErrUserNotFound, "error should be ErrUserNotFound")
}
