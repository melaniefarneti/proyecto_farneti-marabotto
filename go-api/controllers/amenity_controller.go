package controllers

import (
	"go-api/dao"
	"go-api/dto"
	"go-api/services"
	"net/http"
	"strconv"

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

func GetAmenityByHotelID(ctx *gin.Context) {
	hotelIDStr := ctx.Param("hotelID")

	// Convertir hotelIDStr a un valor entero
	hotelID, err := strconv.ParseInt(hotelIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel ID",
		})
		return
	}

	// Crear una instancia del servicio de amenidades
	amenityService := services.NewAmenityService()

	// Llamar al servicio para obtener las amenidades por ID de hotel
	amenities, err := amenityService.GetAmenityByHotelID(hotelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting amenities",
		})
		return
	}

	ctx.JSON(http.StatusOK, amenities)
}
