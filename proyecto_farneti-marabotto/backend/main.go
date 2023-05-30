package main

import (
    "database/sql"
    "encoding/json"
    "time"
	"fmt"
	"go-api/handlers"
	"go-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

)

// Estructura del usuario en la base de datos
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Estructura de la respuesta del token JWT
type TokenResponse struct {
	Token string `json:"token"`
}

// Clave secreta para firmar los tokens JWT
var secretKey = []byte("clave_secreta")

// Conexión a la base de datos MySQL
var db *sql.DB

func main() {
	// Inicializar los servicios
	hotelService := services.NewHotelService()
	reservationService := services.NewReservationService()
	amenityService := services.NewAmenityService()

	// Crear el enrutador Gin
	router := gin.Default()

	// Definir las rutas y los manejadores
	router.GET("/hotels", handlers.GetHotels(hotelService))
	router.GET("/reservations/:id", handlers.GetReservationByID(reservationService))
	router.POST("/reservations", handlers.CreateReservation(reservationService))
	router.GET("/reservations/hotel/:hotelID/date/:date", handlers.GetReservationsByHotelAndDay(reservationService))
	router.GET("/amenities", handlers.GetAmenities(amenityService))

	// Iniciar el servidor HTTP
	port := 8080
	serverAddress := fmt.Sprintf(":%d", port)
	log.Printf("Server listening on %s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, router))

	// Inicializar la conexión a la base de datos
    	db, err := sql.Open("mysql", "usuario:contraseña@tcp(localhost:3306)/basededatos")
    	if err != nil {
    		log.Fatal(err)
    	}
    	defer db.Close()

    	// Ruta para el endpoint de autenticación
    	http.HandleFunc("/auth", AuthHandler)

    	// Iniciar el servidor en el puerto 8000
    	log.Println("Servidor iniciado en http://localhost:8000")
    	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Handler para el endpoint de autenticación
func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Leer los datos del usuario y contraseña del body de la solicitud
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validar las credenciales del usuario
	validUser, err := ValidateUser(user.Username, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !validUser {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Generar el token JWT
	tokenString, err := GenerateToken(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Devolver el token en la respuesta
	response := TokenResponse{Token: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Función para validar las credenciales del usuario
func ValidateUser(username, password string) (bool, error) {
	// Realizar la consulta a la base de datos para obtener los datos del usuario
	query := "SELECT password FROM users WHERE username = ?"
	row := db.QueryRow(query, username)

	var hashedPassword string
	err := row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// El usuario no existe
			return false, nil
		}
		return false, err
	}

	// Comparar la contraseña ingresada con la contraseña almacenada en la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// La contraseña no coincide
		return false, nil
	}

	// Las credenciales son válidas
	return true, nil
}

// Función para generar un token JWT
func GenerateToken(username string) (string, error) {
	// Crear la estructura del token con los claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // El token expira en 24 horas
	})

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
