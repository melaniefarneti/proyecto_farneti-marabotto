package app

import (
	"go-api/clients"
	"go-api/controllers"
	"go-api/services"

	"github.com/gin-gonic/gin"
)

const (
	pahtGetHotels   = "/hotels/gethotels"
	pathCreateHotel = "/hotels/createhotel"
	pathDeleteHotel = "/hotels/deletehotel/:hotelID"
	//pathGetHotel          = "/hotels/:hotelID"

	pathCreateReservation = "/reservations"
	//pathGetAmenity     = "/amenities/:amenityID"
	pathGetUserByID             = "/users/:userID"
	pathGetUserByEmail          = "/users/emailuser/:email"
	pathCreateUser              = "/users/createuser"
	pathLoginAdmin              = "/login/admin"
	pathLogin                   = "/login"
	pathGetReservation          = "/reservations/getreservations"
	pathGetReservationByUserId  = "/reservations/getreservationsbyuserid/:userID"
	pathGetReservationByHotelId = "/reservations/getreservationsbyhotelid/:hotelID"
)

// mapRoutes mapea las rutas de la aplicación
func mapRoutes(router *gin.Engine) {
	router.GET(pahtGetHotels, func(ctx *gin.Context) {
		controllers.GetHotels(ctx)
	})
	router.POST(pathCreateHotel, controllers.CreateHotel)
	router.DELETE(pathDeleteHotel, controllers.DeleteHotel)
	//router.GET(pathGetHotel, controllers.GetHotels)
	router.POST(pathCreateReservation, controllers.NewReservationController(clients.NewDBClient()).CreateReservation)
	//router.GET(pathGetAmenity, controllers.GetAmenity)
	router.GET(pathGetUserByID, controllers.NewUserController(services.NewUserService(clients.NewDBClient())).GetUserByID)
	router.GET(pathGetUserByEmail, controllers.NewUserController(services.NewUserService(clients.NewDBClient())).GetUserByEmail)
	router.POST(pathCreateUser, func(ctx *gin.Context) {
		userService := services.NewUserService(clients.NewDBClient())
		userController := controllers.NewUserController(userService)
		userController.CreateUser(ctx)
	})
	router.POST(pathLoginAdmin, controllers.NewUserController(services.NewUserService(clients.NewDBClient())).LoginAdmin)
	router.POST(pathLogin, controllers.NewUserController(services.NewUserService(clients.NewDBClient())).Login)
	router.GET(pathGetReservation, controllers.NewReservationController(clients.NewDBClient()).GetReservations)
	router.GET(pathGetReservationByUserId, controllers.NewReservationController(clients.NewDBClient()).GetReservationsByUserID)
	router.GET(pathGetReservationByHotelId, controllers.NewReservationController(clients.NewDBClient()).GetReservationsByHotelID)
}
