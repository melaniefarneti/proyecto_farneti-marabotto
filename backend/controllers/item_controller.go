package controllers

import (
	"fmt"
	"go-api/services" //contiene los servicios
	"go-api/services/clients" //contiene los clientes utilizados
	"net/http" //para trabajar con HTTP
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	paramItemID = "itemID"
)

func init() {
	// Aquí definimos que para la aplicación real utilizaremos HTTPClient como MLClient
	services.MLClient = clients.HTTPClient{}
}

func GetItem(ctx *gin.Context) {
	// Obtiene el parámetro de id de la URL como una cadena
	idString := ctx.Param(paramItemID)

	// Convierte el ID de tipo cadena a un ID de tipo entero
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error parsing item ID: %w", err))
		return
	}

	// Llama al servicio con el ID de tipo entero
	item, err := services.GetItem(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error getting item: %w", err))
		return
	}

	// Caso exitoso
	ctx.JSON(http.StatusOK, item)
}