package services

import (
	"go-api/dao"
	"go-api/services/clients"
)

// UserServiceInterface define la interfaz para el servicio de usuarios
type UserServiceInterface interface {
	GetUserByID(userID int) (*dao.User, error)
}

// UserService es una implementación del servicio de usuarios
type UserService struct {
	DBClient clients.DBClientInterface
}

// NewUserService crea una nueva instancia del servicio de usuarios
func NewUserService(dbClient clients.DBClientInterface) UserServiceInterface {
	return &UserService{
		DBClient: dbClient,
	}
}

// GetUserByID obtiene un usuario por su ID
func (s *UserService) GetUserByID(userID int) (*dao.User, error) {
	user, err := s.DBClient.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
