package controllers

import (
	"fmt"
	"go-api/services"
	"go-api/services/clients"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	paramUserID = "userID"
)

func init() {
	services.MLClient = clients.HTTPClient{}
}

func GetUser(ctx *gin.Context) {
	idString := ctx.Param(paramUserID)
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error parsing user ID: %w", err))
		return
	}

	user, err := services.GetUser(int(id)) // Convertir id de int64 a int
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error getting user: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
