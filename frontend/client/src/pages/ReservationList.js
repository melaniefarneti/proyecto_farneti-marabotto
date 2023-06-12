import React, { useContext, useEffect, useState } from 'react';
import { UserContext } from '../contexts/UserContext';
import { getReservations, getAdminReservations } from '../services/api';

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
                response = await getAdminReservations(user.hotelId, new Date().toISOString().split('T')[0]);
            } else {
                response = await getReservations();
            }
            setReservations(response);
        } catch (error) {
            console.error('Error fetching reservations:', error.message);
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
                </tr>
                </thead>
                <tbody>
                {reservations.map((reservation) => (
                    <tr key={reservation.id}>
                        <td>{reservation.date}</td>
                        <td>{reservation.hotel}</td>
                        <td>{reservation.name}</td>
                        <td>{reservation.email}</td>
                    </tr>
                ))}
                </tbody>
            </table>
        </div>
    );
};

export default ReservationList;
