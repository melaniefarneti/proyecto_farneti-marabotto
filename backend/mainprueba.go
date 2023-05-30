package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client representa la estructura de la tabla "usuarios"
type Client struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"column:nombre"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:contrasena"`
	Role        string `gorm:"column:rol"`
}

// Hotel representa la estructura de la tabla "hotel"
type Hotel struct {
	ID          int    `gorm:"type:int;primaryKey"`
	Name        string `gorm:"type:varchar(255);column:nombre"`
	Photo       string `gorm:"type:varchar(255);column:foto"`
	Description string `gorm:"type:varchar(1000);column:descripcion"`
	Location    string `gorm:"type:varchar(50);column:ubicacion"`
	Rooms       string `gorm:"type:varchar(255);column:cuartos"`
}

// Reservation representa la estructura de la tabla "reservas"
type Reservation struct {
	ID          int    `gorm:"primaryKey"`
	HotelID     int    `gorm:"column:hotel_id"`
	CheckIn     string `gorm:"column:fecha_desde"`
	CheckOut    string `gorm:"column:fecha_hasta"`
	ClientName  string `gorm:"column:cliente_nombre"`
}

// Amenity representa la estructura de la tabla "amenities"
type Amenity struct {
	ID      int    `gorm:"type:int;primaryKey"`
	Name    string `gorm:"type:longtext;column:nombre"`
	HotelID int    `gorm:"type:int not null;column:hotel_id"`
}

func main() {
	dsn := "root:mel1@tcp(localhost:3306)/proyectohotel?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Automigrate para crear automáticamente las tablas si no existen
	err = db.AutoMigrate(&Client{}, &Hotel{}, &Reservation{}, &Amenity{})
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/clientes", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Obtener todos los clientes
			var clientes []Client
			db.Find(&clientes)
			// Devolver los clientes como respuesta JSON
			json.NewEncoder(w).Encode(clientes)
		case http.MethodPost:
			// Crear un nuevo cliente
			// Obtener los datos del formulario de la solicitud
			nombre := r.FormValue("nombre")
			email := r.FormValue("email")
			contrasena := r.FormValue("contrasena")
			rol := r.FormValue("rol")
			// Crear un nuevo registro de cliente en la base de datos
			cliente := Client{
				Name:     nombre,
				Email:    email,
				Password: contrasena,
				Role:     rol,
			}
			db.Create(&cliente)
			// Devolver una respuesta de éxito
			response := map[string]interface{}{
				"message": "Cliente creado exitosamente",
				"client":  cliente,
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	r.HandleFunc("/clientes/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Obtener un cliente por ID
			// Obtener el ID de la URL
			vars := mux.Vars(r)
			id := vars["id"]

			// Utilizar el ID para buscar el cliente en la base de datos
			var cliente Client
			db.First(&cliente, id)
			if cliente.ID == 0 {
				// Manejar el caso en que no se encuentre el cliente
				response := map[string]interface{}{
					"message": "Cliente no encontrado",
				}
				json.NewEncoder(w).Encode(response)
			} else {
				// Devolver el cliente como respuesta JSON
				json.NewEncoder(w).Encode(cliente)
			}
		case http.MethodDelete:
			// Eliminar un cliente por ID
			// Obtener el ID de la URL
			vars := mux.Vars(r)
			id := vars["id"]

			// Utilizar el ID para buscar el cliente en la base de datos y eliminarlo
			var cliente Client
			db.Delete(&cliente, id)
			// Devolver una respuesta de éxito
			response := map[string]interface{}{
				"message": "Cliente eliminado exitosamente",
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	r.HandleFunc("/hoteles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Obtener todos los hoteles
			var hoteles []Hotel
			db.Find(&hoteles)
			// Devolver los hoteles como respuesta JSON
			json.NewEncoder(w).Encode(hoteles)
		case http.MethodPost:
			// Crear un nuevo hotel
			// Obtener los datos del formulario de la solicitud
			nombre := r.FormValue("nombre")
			foto := r.FormValue("foto")
			descripcion := r.FormValue("descripcion")
			ubicacion := r.FormValue("ubicacion")
			cuartos := r.FormValue("cuartos")
			// Crear un nuevo registro de hotel en la base de datos
			hotel := Hotel{
				Name:        nombre,
				Photo:       foto,
				Description: descripcion,
				Location:    ubicacion,
				Rooms:       cuartos,
			}
			db.Create(&hotel)
			// Devolver una respuesta de éxito
			response := map[string]interface{}{
				"message": "Hotel creado exitosamente",
				"hotel":   hotel,
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	r.HandleFunc("/hoteles/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Obtener un hotel por ID
			// Obtener el ID de la URL
			vars := mux.Vars(r)
			id := vars["id"]

			// Utilizar el ID para buscar el hotel en la base de datos
			var hotel Hotel
			db.First(&hotel, id)
			if hotel.ID == 0 {
				// Manejar el caso en que no se encuentre el hotel
				response := map[string]interface{}{
					"message": "Hotel no encontrado",
				}
				json.NewEncoder(w).Encode(response)
			} else {
				// Devolver el hotel como respuesta JSON
				json.NewEncoder(w).Encode(hotel)
			}
		case http.MethodDelete:
			// Eliminar un hotel por ID
			// Obtener el ID de la URL
			vars := mux.Vars(r)
			id := vars["id"]

			// Utilizar el ID para buscar el hotel en la base de datos y eliminarlo
			var hotel Hotel
			db.Delete(&hotel, id)
			// Devolver una respuesta de éxito
			response := map[string]interface{}{
				"message": "Hotel eliminado exitosamente",
			}
			json.NewEncoder(w).Encode(response)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}