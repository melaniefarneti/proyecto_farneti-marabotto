import React, { useState, useContext } from 'react';
import { UserContext } from '../contexts/UserContext';

const MailList = () => {
    const { user } = useContext(UserContext);
    const [reservations, setReservations] = useState([]);
    const [showReservations, setShowReservations] = useState(false);

    const fetchReservations = async () => {
        try {
            const response = await fetch("/reservations/getreservations", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    // Agrega cualquier otro encabezado necesario, como el token de autenticación si lo requieres
                },
            });

            if (response.ok) {
                const reservationsData = await response.json();
                setReservations(reservationsData);
            } else {
                console.log("Error fetching reservations");
            }
        } catch (error) {
            console.log("Error:", error);
        }
    };

    const deleteReservation = async (reservationId) => {
        try {
            const response = await fetch(`/reservations/getreservations/${reservationId}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                    // Agrega cualquier otro encabezado necesario, como el token de autenticación si lo requieres
                },
            });

            if (response.ok) {
                console.log("Reservation deleted successfully");
                // Actualizar la lista de reservas después de eliminar
                fetchReservations();
            } else {
                console.log("Error deleting reservation");
            }
        } catch (error) {
            console.log("Error:", error);
        }
    };

    return (
        <div className="mail">
            <h1 className="mailTitle">Save time, save money!</h1>
            <span className="mailDesc">Look for reservation</span>
            <div className="mailInputContainer">
                <input type="text" placeholder="Your Email" />
                {user && user.isAdmin && (
                    <button onClick={() => setShowReservations(!showReservations)}>
                        Reservations
                    </button>
                )}
            </div>
            {showReservations && user && user.isAdmin && (
                <div className="reservations-list">
                    <h2>All Reservations</h2>
                    {reservations.map((reservation) => (
                        <div key={reservation.id}>
                            <p>Reservation ID: {reservation.id}</p>
                            <p>Guest Name: {reservation.clientName}</p>
                            <p>Check-in: {reservation.checkin}</p>
                            <p>Check-out: {reservation.checkout}</p>
                            {/* Agrega más detalles de la reserva si es necesario */}
                            <button onClick={() => deleteReservation(reservation.id)}>
                                Delete
                            </button>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default MailList;


