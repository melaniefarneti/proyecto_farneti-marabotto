package controllers

import (
	"encoding/json"
	"go-api/dao"
	"go-api/dto"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetHotels(t *testing.T) {
	// Configurar el enrutador Gin y el contexto de la solicitud
	router := gin.Default()
	router.GET("/hotels", GetHotels)
	request, _ := http.NewRequest("GET", "/hotels", nil)
	response := httptest.NewRecorder()

	// Ejecutar la solicitud HTTP
	router.ServeHTTP(response, request)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, response.Code)

	var hotels []dto.HotelRequest
	err := json.Unmarshal(response.Body.Bytes(), &hotels)
	assert.NoError(t, err)
}

func TestCreateHotel(t *testing.T) {
	// Configurar el enrutador Gin y el contexto de la solicitud
	router := gin.Default()
	router.POST("/hotels", CreateHotel)
	hotelRequest := dao.Hotel{
		Name:        "Hotel Test",
		Photo:       "test.jpg",
		Description: "Test hotel description",
		Location:    "Test location",
		Rooms:       10,
	}
	requestBody, _ := json.Marshal(hotelRequest)
	request, _ := http.NewRequest("POST", "/hotels", strings.NewReader(string(requestBody)))
	response := httptest.NewRecorder()

	// Ejecutar la solicitud HTTP
	router.ServeHTTP(response, request)

	// Verificar el código de estado y el cuerpo de la respuesta
	assert.Equal(t, http.StatusOK, response.Code)

	var createdHotel dto.HotelRequest
	err := json.Unmarshal(response.Body.Bytes(), &createdHotel)
	assert.NoError(t, err)
}
