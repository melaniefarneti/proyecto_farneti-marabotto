import React, { useState } from 'react';
import PropTypes from 'prop-types';
import AmenityForm from "./AmenitiesForm.jsx";

const HotelDetail = ({ hotel }) => {
    const [selectedRoom, setSelectedRoom] = useState('');

    const handleRoomChange = event => {
        setSelectedRoom(event.target.value);
    };

    const handleReservation = () => {
        // Aquí puedes manejar la lógica para confirmar la reserva
        console.log('Hotel seleccionado:', hotel.title);
        console.log('Habitación seleccionada:', selectedRoom);
    };

    return (
        <div>
            <h2>{hotel.title}</h2>
            <img src={hotel.photo} alt={hotel.title} />
            <p>{hotel.description}</p>
            <div>
                <label htmlFor="room">Seleccione una habitación:</label>
                <select id="room" value={selectedRoom} onChange={handleRoomChange}>
                    <option value="">Seleccione...</option>
                    {hotel.rooms.map(room => (
                        <option key={room.id} value={room.id}>
                            {room.name}
                        </option>
                    ))}
                </select>
                <h2>Amenities del Hotel</h2>
                <AmenityForm />
            </div>
            <button onClick={handleReservation}>Confirmar Reserva</button>
        </div>
    );
};

HotelDetail.propTypes = {
    hotel: PropTypes.object.isRequired,
};

export default HotelDetail;