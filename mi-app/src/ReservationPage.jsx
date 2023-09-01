import React, { useState } from "react";
import { format } from "date-fns";

function ReservationPage({ hotelId }) {
  const [reservation, setReservation] = useState({
    hotel_id: parseInt(hotelId, 10),
    checkin: "",
    checkout: "",
    email: ""
  });
  const [reservationStatus, setReservationStatus] = useState(null);

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setReservation((prevReservation) => ({
      ...prevReservation,
      [name]: value
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    
    // Convertir el formato de fecha de dd-mm-yyyy a yyyy-mm-dd
    const formattedReservation = {
      ...reservation,
      checkin: format(new Date(reservation.checkin), "yyyy-MM-dd"),
      checkout: format(new Date(reservation.checkout), "yyyy-MM-dd")
    };
    
    // Enviar la reserva al servidor para guardarla en la base de datos
    try {
      console.log(formattedReservation);
      const response = await fetch(`http://localhost:8080/reservations/${formattedReservation.hotel_id}`, {
        method: "POST",
        body: JSON.stringify(formattedReservation),
        headers: {
          "Content-Type": "application/json"
        }
      });

      if (response.ok) {
        setReservationStatus("Reservación creada con exito.");
      } else {
        const data = await response.json();
        if (response.status === 500) {
          setReservationStatus("Error al crear la reserva. No hay disponibilidad de habitaciones para esa fecha.");
        } else {
          setReservationStatus("Error al crear la reserva. El cliente no existe, debe registrarse.");
        }
      }
    } catch (error) {
      setReservationStatus("Error al crear la reserva. Inténtelo de nuevo.");
    }
  };  

  return (
    <div>
      <h2>Reservar Hotel</h2>
      {reservationStatus && <p>{reservationStatus}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="hotel_id">Hotel ID:</label>
          <input
            type="number"
            id="hotel_id"
            name="hotel_id"
            value={reservation.hotel_id}
            onChange={handleInputChange}
          />
        </div>
        <div>
          <label htmlFor="checkin">Check-in:</label>
          <input
            type="date"
            id="checkin"
            name="checkin"
            value={reservation.checkin}
            onChange={handleInputChange}
          />
        </div>
        <div>
          <label htmlFor="checkout">Check-out:</label>
          <input
            type="date"
            id="checkout"
            name="checkout"
            value={reservation.checkout}
            onChange={handleInputChange}
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            name="email"
            value={reservation.email}
            onChange={handleInputChange}
          />
        </div>
        <button type="submit">Reservar</button>
      </form>
    </div>
  );
}

export default ReservationPage;
