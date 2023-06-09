package controllers

import (
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
