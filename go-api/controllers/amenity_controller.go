package controllers

import (
	"go-api/dao"
	"go-api/dto"
	"go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAmenity(ctx *gin.Context) {
	// Parsear los datos del cuerpo de la solicitud JSON en una estructura AmenityRequest
	var amenityRequest dto.AmenityRequest
	if err := ctx.ShouldBindJSON(&amenityRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	// Verificar si algún campo requerido está en blanco
	if amenityRequest.Nombre == "" || amenityRequest.HotelID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required data",
		})
		return
	}

	// Crear una instancia del servicio de amenities
	amenityService := services.NewAmenityService()

	// Crear un objeto Amenity basado en los datos del AmenityRequest
	amenity := dao.Amenity{
		Nombre:  amenityRequest.Nombre,
		HotelID: amenityRequest.HotelID,
	}

	// Llamar al servicio para crear el amenity
	err := amenityService.CreateAmenity(&amenity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating amenity",
		})
		return
	}

	// Retornar el objeto creado como respuesta
	ctx.JSON(http.StatusOK, amenity)
}
