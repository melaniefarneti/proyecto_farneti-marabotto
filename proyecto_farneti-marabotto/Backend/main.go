package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hotel struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Otras propiedades del hotel
}

type Reservation struct {
	ID           int    `json:"id"`
	Date         string `json:"date"`
	CustomerName string `json:"customerName"`
	// Otras propiedades de la reserva
}

var (
	hotels       []Hotel
	reservations []Reservation
)

func main() {
	// Rutas para obtener hoteles y reservas
	http.HandleFunc("/api/hotels", getHotelsHandler)
	http.HandleFunc("/api/reservations", getReservationsHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getHotelsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotels)
}

func getReservationsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}