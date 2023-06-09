package services

import (
	"errors"
	"go-api/dao"
	"go-api/services/clients"

	"time"

	"github.com/dgrijalva/jwt-go"
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

// CreateUser registra un nuevo usuario
func (s *UserService) CreateUser(user *dao.User) (*dao.User, error) {
	// Realizar cualquier validación o lógica de negocio adicional aquí antes de crear el usuario en la base de datos

	// Llamar al cliente de la base de datos para crear el usuario
	createdUser, err := s.DBClient.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// HashPassword encripta la contraseña utilizando bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword verifica si la contraseña ingresada coincide con la contraseña encriptada almacenada
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Login realiza el proceso de autenticación y devuelve un token de acceso si las credenciales son válidas
func (s *UserService) Login(email, password string) (string, error) {
	// Llamar al cliente de la base de datos para obtener el usuario por su dirección de correo electrónico
	user, err := s.DBClient.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verificar el tipo de usuario (administrador o cliente)
	if user.Role != "admin" {
		return "", errors.New("only administrators can log in")
	}

	// Verificar la contraseña
	err = CheckPassword(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generar y devolver un token de acceso
	token, err := GenerateAccessToken(user)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return token, nil
}

// GenerateAccessToken genera un token de acceso JWT para el usuario dado
func GenerateAccessToken(user *dao.User) (string, error) {
	// Crear un nuevo token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Agregar los claims (datos) al token
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expira en 24 horas

	// Firmar el token con una clave secreta
	// Debes proporcionar tu propia clave secreta en lugar de "mysecretkey"
	tokenString, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
