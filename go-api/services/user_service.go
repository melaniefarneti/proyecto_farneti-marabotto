package services

import (
	"go-api/dao"
	"go-api/services/clients"

	"golang.org/x/crypto/bcrypt"
)

// UserServiceInterface define la interfaz para el servicio de usuarios
type UserServiceInterface interface {
	GetUserByID(userID int) (*dao.User, error)
	GetUserByEmail(email string) (*dao.User, error)
	CreateUser(user *dao.User) (*dao.User, error)
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

func (s *UserService) GetUserByEmail(email string) (*dao.User, error) {
	user, err := s.DBClient.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) CreateUser(user *dao.User) (*dao.User, error) {
	// Realizar cualquier validación o lógica de negocio adicional aquí antes de crear el usuario en la base de datos

	// Hashear la contraseña del usuario antes de almacenarla
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// Actualizar la contraseña hasheada en el objeto de usuario
	user.Password = hashedPassword

	// Llamar al cliente de la base de datos para crear el usuario
	createdUser, err := s.DBClient.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
