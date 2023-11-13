import React, { useState, useEffect } from "react";

function InformationPage() {
  const [reservationsByUserId, setReservationsByUserId] = useState([]);
  const [reservationsByHotelId, setReservationsByHotelId] = useState([]);
  const [availableRooms, setAvailableRooms] = useState([]);
  const [selectedHotelId, setSelectedHotelId] = useState("");
  const [checkInDate, setCheckInDate] = useState("");
  const [checkOutDate, setCheckOutDate] = useState("");
  const [userEmail, setUserEmail] = useState("");
  const [hotelIdInput, setHotelIdInput] = useState("");

  const fetchUserByEmail = async (email) => {
    try {
      const response = await fetch(`http://localhost:8080/users/emailuser/${email}`);
      const data = await response.json();
      return data.ID;
    } catch (error) {
      console.error(error);
      return null;
    }
  };

  const fetchReservationsByUserId = async (userId) => {
    try {
      const response = await fetch(`http://localhost:8080/reservations/getreservationsbyuserid/${userId}`);
      const data = await response.json();
      setReservationsByUserId(data.reservations);
    } catch (error) {
      console.error(error);
    }
  };

  const fetchReservationsByHotelId = async (hotelId) => {
    try {
      const response = await fetch(`http://localhost:8080/reservations/getreservationsbyhotelid/${hotelId}`);
      const data = await response.json();
      setReservationsByHotelId(data.reservations);
    } catch (error) {
      console.error(error);
    }
  };

  const fetchAvailableRoomsByHotelId = async (hotelId, checkIn, checkOut) => {
    try {
      const response = await fetch(`http://localhost:8080/availablerooms/${hotelId}/${checkIn}/${checkOut}`);
      const data = await response.json();
      setAvailableRooms(data.availableRooms);
    } catch (error) {
      console.error(error);
    }
  };

  const handleUserEmailChange = (event) => {
    setUserEmail(event.target.value);
  };

  const handleHotelIdInputChange = (event) => {
    setHotelIdInput(event.target.value);
  };

  const handleCheckInDateChange = (event) => {
    setCheckInDate(event.target.value);
  };

  const handleCheckOutDateChange = (event) => {
    setCheckOutDate(event.target.value);
  };

  const handleUserReservationsSearch = async () => {
    const userId = await fetchUserByEmail(userEmail);
    if (userId) {
      fetchReservationsByUserId(userId);
    }
  };

  const handleHotelReservationsSearch = () => {
    fetchReservationsByHotelId(hotelIdInput);
  };

  const handleAvailabilitySearch = () => {
    fetchAvailableRoomsByHotelId(selectedHotelId, checkInDate, checkOutDate);
  };

  return (
    <div className="information-container">
      <div className="user-reservations-section">
        <h2>Listado de Reservas por Usuario</h2>
        <div>
          <label htmlFor="userEmail">Email del Usuario:</label>
          <input
            type="text"
            id="userEmail"
            value={userEmail}
            onChange={handleUserEmailChange}
          />
          <button onClick={handleUserReservationsSearch}>Buscar</button>
        </div>
        <ul>
          {reservationsByUserId.map((reservation) => (
            <li key={reservation.ID}>
              <p>Como realizada el día: {reservation.CreatedAt}</p>
              <p>Hotel ID: {reservation.HotelID}</p>
              <p>Check-in: {reservation.CheckIn}</p>
              <p>Check-out: {reservation.CheckOut}</p>
              <p>Client: {reservation.ClientName}</p>
            </li>
          ))}
        </ul>
      </div>

      <div className="hotel-reservations-section">
        <h2>Listado de Reservas por Hotel</h2>
        <div>
          <label htmlFor="hotelIdInput">ID del Hotel:</label>
          <input
            type="text"
            id="hotelIdInput"
            value={hotelIdInput}
            onChange={handleHotelIdInputChange}
          />
          <button onClick={handleHotelReservationsSearch}>Buscar</button>
        </div>
        <ul>
          {reservationsByHotelId.map((reservation) => (
            <li key={reservation.ID}>
              <p>Como realizada el día: {reservation.CreatedAt}</p>
              <p>Hotel ID: {reservation.HotelID}</p>
              <p>Check-in: {reservation.CheckIn}</p>
              <p>Check-out: {reservation.CheckOut}</p>
              <p>Client: {reservation.ClientName}</p>
            </li>
          ))}
        </ul>
      </div>

      <div className="availability-section">
        <h2>Disponibilidad de Habitaciones</h2>
        <div>
          <label htmlFor="selectedHotelId">ID de Hotel:</label>
          <input
            type="text"
            id="selectedHotelId"
            value={selectedHotelId}
            onChange={(event) => setSelectedHotelId(event.target.value)}
          />
        </div>
        <div>
          <label htmlFor="checkInDate">Fecha de Check-in:</label>
          <input
            type="date"
            id="checkInDate"
            value={checkInDate}
            onChange={handleCheckInDateChange}
          />
        </div>
        <div>
          <label htmlFor="checkOutDate">Fecha de Check-out:</label>
          <input
            type="date"
            id="checkOutDate"
            value={checkOutDate}
            onChange={handleCheckOutDateChange}
          />
        </div>
        <button onClick={handleAvailabilitySearch}>Buscar Disponibilidad</button>
        <p>Número de habitaciones disponibles: {availableRooms}</p>
      </div>
    </div>
  );
}

export default InformationPage;
