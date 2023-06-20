package controllers

import (
	"errors"
	"go-api/dao"
	"go-api/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{
		GetUserByIDFunc: func(userID int) (*dao.User, error) {
			// Simular la recuperación exitosa del usuario
			user := &dao.User{
				ID:    userID,
				Name:  "John Doe",
				Email: "john.doe@example.com",
				Role:  "admin",
			}
			return user, nil
		},
	}
	userController := NewUserController(userServiceMock)
	router.GET("/users/:userID", userController.GetUserByID)

	// Preparar una solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/users/1", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"id": 1, "name": "John Doe", "email": "john.doe@example.com", "role": "admin"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestGetUserByID_InvalidUserID(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{}
	userController := NewUserController(userServiceMock)
	router.GET("/users/:userID", userController.GetUserByID)

	// Preparar una solicitud HTTP de prueba con un ID de usuario inválido (no numérico)
	req, _ := http.NewRequest("GET", "/users/invalid", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"error": "invalid user ID"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestGetUserByEmail(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{
		GetUserByEmailFunc: func(email string) (*dao.User, error) {
			// Simular la recuperación exitosa del usuario
			user := &dao.User{
				ID:    1,
				Name:  "John Doe",
				Email: email,
				Role:  "admin",
			}
			return user, nil
		},
	}
	userController := NewUserController(userServiceMock)
	router.GET("/users/email/:email", userController.GetUserByEmail)

	// Preparar una solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/users/email/john.doe@example.com", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"id": 1, "name": "John Doe", "email": "john.doe@example.com", "role": "admin"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestGetUserByEmail_UserNotFound(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{
		GetUserByEmailFunc: func(email string) (*dao.User, error) {
			// Simular la no existencia del usuario
			return nil, dao.ErrUserNotFound
		},
	}
	userController := NewUserController(userServiceMock)
	router.GET("/users/email/:email", userController.GetUserByEmail)

	// Preparar una solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/users/email/nonexistent@example.com", nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"error": "error getting user"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestCreateUser(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{
		GetUserByEmailFunc: func(email string) (*dao.User, error) {
			// Simular la no existencia del usuario
			return nil, dao.ErrUserNotFound
		},
		CreateUserFunc: func(user *dao.User) (*dao.User, error) {
			// Simular la creación exitosa del usuario
			user.ID = 1
			return user, nil
		},
	}
	userController := NewUserController(userServiceMock)
	router.POST("/users", userController.CreateUser)

	// Preparar una solicitud HTTP de prueba
	jsonStr := `{
		"name": "John Doe",
		"email": "john.doe@example.com",
		"password": "password123",
		"role": "admin"
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"id": 1, "name": "John Doe", "email": "john.doe@example.com", "role": "admin"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestCreateUser_InvalidUserData(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{}
	userController := NewUserController(userServiceMock)
	router.POST("/users", userController.CreateUser)

	// Preparar una solicitud HTTP de prueba con datos de usuario no válidos (falta el campo "name")
	jsonStr := `{
		"email": "john.doe@example.com",
		"password": "password123",
		"role": "admin"
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"error": "invalid user data"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestCreateUser_UserAlreadyExists(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{
		GetUserByEmailFunc: func(email string) (*dao.User, error) {
			// Simular la existencia del usuario
			user := &dao.User{
				ID:    1,
				Name:  "John Doe",
				Email: email,
				Role:  "admin",
			}
			return user, nil
		},
	}
	userController := NewUserController(userServiceMock)
	router.POST("/users", userController.CreateUser)

	// Preparar una solicitud HTTP de prueba con un usuario que ya existe
	jsonStr := `{
		"name": "John Doe",
		"email": "john.doe@example.com",
		"password": "password123",
		"role": "admin"
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"error": "user already exists"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}

func TestCreateUser_ErrorCreatingUser(t *testing.T) {
	// Crear un enrutador Gin y configurar el controlador
	router := gin.Default()
	userServiceMock := &services.UserServiceMock{
		GetUserByEmailFunc: func(email string) (*dao.User, error) {
			// Simular la no existencia del usuario
			return nil, dao.ErrUserNotFound
		},
		CreateUserFunc: func(user *dao.User) (*dao.User, error) {
			// Simular un error al crear el usuario
			return nil, errors.New("failed to create user")
		},
	}
	userController := NewUserController(userServiceMock)
	router.POST("/users", userController.CreateUser)

	// Preparar una solicitud HTTP de prueba
	jsonStr := `{
		"name": "John Doe",
		"email": "john.doe@example.com",
		"password": "password123",
		"role": "admin"
	}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code, "Código de estado de respuesta incorrecto")

	// Verificar el cuerpo de la respuesta
	assert.JSONEq(t, `{"error": "error creating user"}`, resp.Body.String(), "Cuerpo de respuesta incorrecto")
}
