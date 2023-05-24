package app

import (
	"fmt" //para formatear y mostrar mensajes
	"github.com/gin-gonic/gin" //framework Gin
)

const (
	port = ":8080"
)

/*Aquí es donde se configura y se inicia el servidor web*/
func StartApp() {
	router := gin.Default() // configura el enrutador con la configuración predeterminada de Gin,
	                        // que incluye el registro de solicitudes y recuperación automática de errores.

	mapRoutes(router) //mapea y define las rutas y controladores de la aplicación en el enrutador.

	err := router.Run(port) // escucha las solicitudes entrantes y maneja las rutas registradas en el
	                        // enrutador. Si ocurre algún error durante la ejecución del servidor, se
	                        // captura y se produce un pánico.
	if err != nil {
		panic(fmt.Errorf("error running app: %w", err))
	}
}