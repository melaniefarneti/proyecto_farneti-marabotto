package controllers

/*
import (
	"fmt"
	"go-api/services"
	"go-api/services/clients"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	paramAmenityID = "amenityID"
)

func init() {
	services.MLClient = clients.HTTPClient{}
}

func GetAmenity(ctx *gin.Context) {
	idString := ctx.Param(paramAmenityID)
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error parsing amenity ID: %w", err))
		return
	}

	amenity, err := services.GetAmenity(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error getting amenity: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, amenity)
}*/
