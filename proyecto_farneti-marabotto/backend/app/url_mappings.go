package app

import (
	"go-api/controllers"

	"github.com/gin-gonic/gin"
)

const (
	pathGetItem = "/items/:itemID" //RUTA PARA OBTENER EL ITEM. DEBERIAMOS CAMBIARLO
)

/* Cuando se realiza una solicitud GET a la ruta especificada,
el controlador GetItem del paquete "go-api/controllers"
se ejecutará para manejar la solicitud.*/

func mapRoutes(router *gin.Engine) {
	router.GET(pathGetItem, controllers.GetItem)
}