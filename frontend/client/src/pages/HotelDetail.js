import React, { useContext } from 'react';
import { useParams } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext';

const HotelDetail = ({ hotels }) => {
    const { id } = useParams();
    const { user } = useContext(UserContext);

    const hotel = hotels.find(hotel => hotel.id === id);

    const handleBookNow = () => {
        // Lógica para reservar el hotel
        // Puedes implementar la funcionalidad de reserva según tus requerimientos
    };

    if (!hotel) {
        return <p>Hotel not found.</p>;
    }

    return (
        <div className="hotel-detail">
            <img src={hotel.image} alt={hotel.title} />
            <h3>{hotel.title}</h3>
            <p>{hotel.description}</p>
            <div className="hotel-details">
                <p>Rooms: {hotel.rooms}</p>
                <p>Price: ${hotel.price} per night</p>
                {user ? (
                    <button onClick={handleBookNow}>Book Now</button>
                ) : (
                    <p>Please login to make a reservation.</p>
                )}
            </div>
        </div>
    );
};

export default HotelDetail;
