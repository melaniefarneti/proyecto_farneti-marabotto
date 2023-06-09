package services

import (
	"fmt"
	"go-api/dao"
)

// UserServiceMock es una implementación simulada del servicio de usuarios para pruebas
type UserServiceMock struct {
	GetUserByIDFunc    func(userID int) (*dao.User, error)
	GetUserByEmailFunc func(email string) (*dao.User, error)
	CreateUserFunc     func(user *dao.User) (*dao.User, error)
}

// GetUserByID obtiene un usuario simulado por su ID
func (m *UserServiceMock) GetUserByID(userID int) (*dao.User, error) {
	return m.GetUserByIDFunc(userID)
}

// NewUserServiceMock crea una nueva instancia del servicio de usuarios simulado para pruebas
func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}

func (m *UserServiceMock) GetUserByEmail(email string) (*dao.User, error) {
	if m.GetUserByEmail != nil {
		return m.GetUserByEmail(email)
	}
	return nil, fmt.Errorf("GetUserByEmail not implemented")
}

func (m *UserServiceMock) CreateUser(user *dao.User) (*dao.User, error) {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(user)
	}
	return nil, fmt.Errorf("CreateUserFunc not implemented")
}
