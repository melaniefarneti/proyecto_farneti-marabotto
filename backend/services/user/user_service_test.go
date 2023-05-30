package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-api/domain"
)

func TestGetUserByID(t *testing.T) {
	// Prepare
	users = []domain.User{
		{ID: 1, Name: "John", Email: "john@example.com"},
		{ID: 2, Name: "Jane", Email: "jane@example.com"},
	}

	// Test existing user
	user, err := GetUserByID(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John", user.Name)
	assert.Equal(t, "john@example.com", user.Email)

	// Test non-existing user
	_, err = GetUserByID(3)
	assert.NotNil(t, err)
	assert.Equal(t, domain.ErrUserNotFound, err)
}

func TestCreateUser(t *testing.T) {
	// Prepare
	users = []domain.User{
		{ID: 1, Name: "John", Email: "john@example.com"},
	}

	// Test new user
	newUser := domain.User{Name: "Jane", Email: "jane@example.com"}
	createdUser, err := CreateUser(newUser)
	assert.Nil(t, err)
	assert.Equal(t, 2, createdUser.ID)
	assert.Equal(t, "Jane", createdUser.Name)
	assert.Equal(t, "jane@example.com", createdUser.Email)

	// Verify user added to the list
	assert.Equal(t, 2, len(users))
	assert.Equal(t, newUser, users[1])

	// Test duplicate email
	duplicateUser := domain.User{Name: "Mark", Email: "john@example.com"}
	_, err = CreateUser(duplicateUser)
	assert.NotNil(t, err)
	assert.Equal(t, "user already exists", err.Error())

	// Verify user not added to the list
	assert.Equal(t, 2, len(users))
}

