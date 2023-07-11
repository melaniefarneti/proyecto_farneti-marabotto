package controllers

import (
	"go-api/dao"
	"go-api/dto"
	"go-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHotels(ctx *gin.Context) {
	hotelService := services.NewHotelService() // Crear una instancia del servicio de hoteles

	// Llama al servicio para obtener el listado de hoteles
	hotels, err := hotelService.GetHotels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting hotels",
		})
		return
	}

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

	ctx.JSON(http.StatusOK, hotels)
}

func CreateHotel(ctx *gin.Context) {
	// Parsear los datos del cuerpo de la solicitud JSON en una estructura HotelRequest
	var hotelRequest dto.HotelRequest
	if err := ctx.ShouldBindJSON(&hotelRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	// Verificar si algún campo requerido está en blanco
	if hotelRequest.Name == "" || hotelRequest.Photo == "" || hotelRequest.Description == "" || hotelRequest.Location == "" || hotelRequest.Rooms == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing required data",
		})
		return
	}

	// Crear una instancia del servicio de hoteles
	hotelService := services.NewHotelService()

	// Crear un objeto Hotel basado en los datos del HotelRequest
	hotel := dao.Hotel{
		Name:        hotelRequest.Name,
		Photo:       hotelRequest.Photo,
		Description: hotelRequest.Description,
		Location:    hotelRequest.Location,
		Rooms:       hotelRequest.Rooms,
	}

	// Llamar al servicio para crear el hotel
	createdHotel, err := hotelService.CreateHotel(&hotel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating hotel",
		})
		return
	}

	// Retornar el objeto creado como respuesta
	ctx.JSON(http.StatusOK, createdHotel)
}

func DeleteHotel(ctx *gin.Context) {
	hotelIDStr := ctx.Param("hotelID")

	// Convertir hotelIDStr a un valor entero
	hotelID, err := strconv.Atoi(hotelIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel ID",
		})
		return
	}

	hotelService := services.NewHotelService()
	err = hotelService.DeleteHotel(hotelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error deleting hotel",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "hotel deleted successfully",
	})
}

func UploadHotelPhoto(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read photo"})
		return
	}
	defer file.Close()

	// Obtener el ID del hotel del parámetro de la ruta
	hotelIDStr := ctx.Param("hotelID")
	hotelID, err := strconv.Atoi(hotelIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid hotel ID"})
		return
	}

	// Crear una instancia del servicio de hoteles
	hotelService := services.NewHotelService()

	// Llamar al servicio para cargar la foto del hotel
	err = hotelService.UploadHotelPhoto(dto.HotelPhoto{HotelID: hotelID}, file, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error uploading hotel photo"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "hotel photo uploaded"})
}

func GetHotelPhotos(ctx *gin.Context) {
	hotelIDStr := ctx.Param("hotelID")

	// Convertir hotelIDStr a un valor entero
	hotelID, err := strconv.Atoi(hotelIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid hotel ID",
		})
		return
	}

	hotelService := services.NewHotelService()
	photos, err := hotelService.GetHotelPhotos(hotelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error getting hotel photos",
		})
		return
	}

	baseURL := "./uploads/" // Ruta relativa a la carpeta "uploads"
	for i := range photos {
		photos[i].Filename = baseURL + photos[i].Filename
	}

	ctx.JSON(http.StatusOK, photos)
}
