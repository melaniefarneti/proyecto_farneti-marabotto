// ReservationPage.jsx

import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { format } from 'date-fns';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckCircle } from '@fortawesome/free-solid-svg-icons';
import './styles.css';

function ReservationPage() {
  const { hotelId } = useParams();

  const [reservation, setReservation] = useState({
    hotel_id: parseInt(hotelId, 10) || 0,
    checkin: '',
    checkout: '',
    email: ''
  });

  const [reservationStatus, setReservationStatus] = useState(null);
  const [reservationSuccess, setReservationSuccess] = useState(false);

  useEffect(() => {
    console.log("hotelId in ReservationPage:", hotelId);
    setReservation((prevReservation) => ({
      ...prevReservation,
      hotel_id: parseInt(hotelId, 10) || 0
    }));
  }, [hotelId]);

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setReservation((prevReservation) => ({
      ...prevReservation,
      [name]: value
    }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log("Submitting with hotel_id:", reservation.hotel_id);

    // Resto del código...

    try {
      const response = await fetch(`http://localhost:8080/reservations/${reservation.hotel_id}`, {
        method: "POST",
        body: JSON.stringify(reservation),
        headers: {
          "Content-Type": "application/json"
        }
      });

      if (response.ok) {
        setReservationSuccess(true);
        setReservationStatus("Reservación creada con éxito.");
      } else {
        // Resto del código...
      }
    } catch (error) {
      console.error("Error:", error);
      setReservationStatus("Error al crear la reserva. Inténtelo de nuevo.");
    }
  };

  return (
    <div className="reservation-container">
      <h2>Reservar Hotel</h2>
      {reservationStatus && <p className={reservationSuccess ? 'success' : 'error'}>{reservationStatus}</p>}
      <form className="reservation-form" onSubmit={handleSubmit}>
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
