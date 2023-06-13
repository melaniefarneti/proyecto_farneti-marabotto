import React, { useContext, useEffect, useState } from 'react';
import axios from 'axios';
import { UserContext } from '../contexts/UserContext';
import { getReservations, deleteReservation } from '../services/api';

const ReservationList = () => {
    const { user } = useContext(UserContext);
    const [reservations, setReservations] = useState([]);

    useEffect(() => {
        fetchReservations();
    }, []);

    const fetchReservations = async () => {
        try {
            let response;
            if (user.isAdmin) {
                response = await axios.get(`/admin/reservations`);
            } else {
                response = await getReservations();
            }
            setReservations(response.data);
        } catch (error) {
            console.error('Error fetching reservations:', error.message);
        }
    };

    const handleDeleteReservation = async (reservationId) => {
        try {
            if (user.isAdmin) {
                await deleteReservation(reservationId);
                fetchReservations();
            } else {
                console.log('Only administrators can delete reservations.');
            }
        } catch (error) {
            console.error('Error deleting reservation:', error.message);
        }
    };

    return (
        <div className="reservation-list">
            <h3>Reservation List</h3>
            <table>
                <thead>
                <tr>
                    <th>Date</th>
                    <th>Hotel</th>
                    <th>Name</th>
                    <th>Email</th>
                    {user.isAdmin && <th>Action</th>}
                </tr>
                </thead>
                <tbody>
                {reservations.map((reservation) => (
                    <tr key={reservation.id}>
                        <td>{reservation.date}</td>
                        <td>{reservation.hotel}</td>
                        <td>{reservation.name}</td>
                        <td>{reservation.email}</td>
                        {user.isAdmin && (
                            <td>
                                <button onClick={() => handleDeleteReservation(reservation.id)}>Delete</button>
                            </td>
                        )}
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default ReservationList;
