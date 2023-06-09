package app

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	pahtGetHotels   = "/hotels/gethotels"
	pathCreateHotel = "/hotels/createhotel"
	pathDeleteHotel = "/hotels/deletehotel/:hotelID"
	//pathGetHotel          = "/hotels/:hotelID"

	pathCreateReservation = "/reservations"
	//pathGetAmenity     = "/amenities/:amenityID"
	pathGetUser = "/users/:userID"
)

// mapRoutes mapea las rutas de la aplicación
func mapRoutes(router *gin.Engine) {
	router.GET(pahtGetHotels, func(ctx *gin.Context) {
		controllers.GetHotels(ctx)
	})
	router.POST(pathCreateHotel, controllers.CreateHotel)
	router.DELETE(pathDeleteHotel, controllers.DeleteHotel)
	//router.GET(pathGetHotel, controllers.GetHotels)
	router.POST(pathCreateReservation, controllers.CreateReservation)
	//router.GET(pathGetAmenity, controllers.GetAmenity)
	router.GET(pathGetUser, controllers.GetUser)
}
