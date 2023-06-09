package controllers

import (
	"errors"
	"go-api/dao"
	"go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController es el controlador de usuarios
type UserController struct {
	UserService services.UserServiceInterface
}

// NewUserController crea una nueva instancia del controlador de usuarios
func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// GetUserByID obtiene un usuario por su ID
func (c *UserController) GetUserByID(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")

	// Convertir userIDStr a un valor entero
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	// Llamar al servicio para obtener el usuario por su ID
	user, err := c.UserService.GetUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting user",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GetUserByEmail obtiene un usuario por su dirección de correo electrónico
func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	// Llamar al servicio para obtener el usuario por su dirección de correo electrónico
	user, err := c.UserService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting user",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	// Obtener los datos del usuario del cuerpo de la solicitud
	var user dao.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user data",
		})
		return
	}

	// Verificar si el usuario ya existe en la base de datos
	existingUser, err := c.UserService.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, dao.ErrUserNotFound) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error checking user existence",
		})
		return
	}
	if existingUser != nil && existingUser.Email == user.Email {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "user already exists",
		})
		return
	}

	// Encriptar la contraseña antes de almacenarla
	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error hashing password",
		})
		return
	}

	// Actualizar la contraseña hasheada en el objeto de usuario
	user.Password = hashedPassword

	// Llamar al servicio de usuarios para crear el usuario
	createdUser, err := c.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// Login realiza el proceso de autenticación y devuelve un token de acceso si las credenciales son válidas
func (c *UserController) Login(ctx *gin.Context) {
	// Obtener las credenciales del cuerpo de la solicitud
	var credentials dao.User
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Llamar al servicio para obtener el usuario por su dirección de correo electrónico
	user, err := c.UserService.GetUserByEmail(credentials.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Verificar la contraseña
	err = services.CheckPassword(credentials.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	// Generar un token de acceso
	token, err := services.GenerateAccessToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate access token",
		})
		return
	}

	// Devolver el token de acceso en la respuesta
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
