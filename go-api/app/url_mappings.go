package app

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	pahtGetHotels = "/hotels/all"
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
	//router.GET(pathGetHotel, controllers.GetHotels)
	router.POST(pathCreateReservation, controllers.CreateReservation)
	//router.GET(pathGetAmenity, controllers.GetAmenity)
	router.GET(pathGetUser, controllers.GetUser)
}

// SetupRoutes configura y mapea las rutas de la aplicación
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Mapear las rutas
	mapRoutes(router)

	return router
}
