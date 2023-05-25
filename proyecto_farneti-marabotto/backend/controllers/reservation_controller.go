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
	paramReservationID = "reservationID"
)

func init() {
	services.MLClient = clients.HTTPClient{}
}

func GetReservation(ctx *gin.Context) {
	idString := ctx.Param(paramReservationID)
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("error parsing reservation ID: %w", err))
		return
	}

	reservation, err := services.GetReservation(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Errorf("error getting reservation: %w", err))
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

