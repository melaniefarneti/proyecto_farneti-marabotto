import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext.js';

const HotelCard = ({ hotel }) => {
    const { user } = useContext(UserContext);

    const handleBookNow = () => {
        // Lógica para reservar el hotel
        // Puedes implementar la funcionalidad de reserva según tus requerimientos
    };

    return (
        <div className="hotel-card">
            <img src={hotel.image} alt={hotel.title} />
            <h3>{hotel.title}</h3>
            <p>{hotel.description}</p>
            <div className="hotel-details">
                <p>Rooms: {hotel.rooms}</p>
                <p>Price: ${hotel.price} per night</p>
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
