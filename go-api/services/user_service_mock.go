package services

import "go-api/dao"

// UserServiceMock es una implementación simulada del servicio de usuarios para pruebas
type UserServiceMock struct {
	GetUserByIDFunc func(userID int) (*dao.User, error)
}

// GetUserByID obtiene un usuario simulado por su ID
func (m *UserServiceMock) GetUserByID(userID int) (*dao.User, error) {
	return m.GetUserByIDFunc(userID)
}

// NewUserServiceMock crea una nueva instancia del servicio de usuarios simulado para pruebas
func NewUserServiceMock() *UserServiceMock {
	return &UserServiceMock{}
}
