import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext';

const HotelCard = ({ hotel }) => {
    const { user } = useContext(UserContext);

    const handleBookNow = async () => {
        if (user) {
            try {
                const response = await fetch('/reservations', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        hotelId: hotel.id,
                        userId: user.id,
                    }),
                });

                if (response.ok) {
                    console.log('Reservation created successfully');
                    // Aquí puedes realizar alguna acción adicional después de crear la reserva
                } else {
                    console.log('Failed to create reservation');
                }
            } catch (error) {
                console.log('Error:', error);
            }
        } else {
            console.log('User not authenticated. Redirect to login page.');
            // Aquí redirige al usuario a la página de inicio de sesión
        }
    };

    return (
        <div className="hotel-card">
            <img src={hotel.image} alt={hotel.title} />
            <h3>{hotel.title}</h3>
            <p>{hotel.description}</p>
            <div className="hotel-details">
                <p>Rooms: {hotel.rooms}</p>
                {user ? (
                    <button onClick={handleBookNow}>Book Now</button>
                ) : (
                    <p>
                        Please <Link to="/login">login</Link> to make a reservation.
                    </p>
                )}
            </div>
        </div>
    );
};

export default HotelCard;
