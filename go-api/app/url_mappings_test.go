package app

import (
	"go-api/clients"
	"go-api/controllers"
	"go-api/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMapRoutes(t *testing.T) {
	router := gin.Default() // Crea un enrutador Gin para el test
	mapRoutes(router)       // Mapea las rutas en el enrutador

	// Verifica que las rutas estén correctamente mapeadas
	assertRouteExists(t, router, "GET", "/hotels/gethotels", controllers.GetHotels)
	assertRouteExists(t, router, "POST", "/hotels/createhotel", controllers.CreateHotel)
	assertRouteExists(t, router, "DELETE", "/hotels/deletehotel/:hotelID", controllers.DeleteHotel)
	//assertRouteExists(t, router, "GET", "/hotels/:hotelID", controllers.GetHotels)
	assertRouteExists(t, router, "POST", "/reservations", controllers.NewReservationController(clients.NewDBClient()).CreateReservation)
	//assertRouteExists(t, router, "GET", "/amenities/:amenityID", controllers.GetAmenity)
	assertRouteExists(t, router, "GET", "/users/:userID", controllers.NewUserController(services.NewUserService(clients.NewDBClient())).GetUserByID)
	assertRouteExists(t, router, "GET", "/users/emailuser/:email", controllers.NewUserController(services.NewUserService(clients.NewDBClient())).GetUserByEmail)
	assertRouteExists(t, router, "POST", "/users/createuser", func(ctx *gin.Context) {
		userService := services.NewUserService(clients.NewDBClient())
		userController := controllers.NewUserController(userService)
		userController.CreateUser(ctx)
	})
	assertRouteExists(t, router, "POST", "/login/admin", controllers.NewUserController(services.NewUserService(clients.NewDBClient())).LoginAdmin)
	assertRouteExists(t, router, "POST", "/login", controllers.NewUserController(services.NewUserService(clients.NewDBClient())).Login)
	assertRouteExists(t, router, "GET", "/reservations/getreservations", controllers.NewReservationController(clients.NewDBClient()).GetReservations)
}

func assertRouteExists(t *testing.T, router *gin.Engine, method, path string, handler gin.HandlerFunc) {
	routeInfo := findRouteInfo(router, method, path)
	assert.NotNilf(t, routeInfo, "Route %s %s not found", method, path)
	assert.Equalf(t, handler, routeInfo.Handler, "Incorrect handler for route %s %s", method, path)
}

func findRouteInfo(router *gin.Engine, method, path string) *gin.RouteInfo {
	for _, routeInfo := range router.Routes() {
		if routeInfo.Method == method && routeInfo.Path == path {
			return &routeInfo
		}
	}
	return nil
}
