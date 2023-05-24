import React from 'react';
import PropTypes from 'prop-types';


const ReservationList = ({ reservations }) => {
    return (
        <div>
            <h2>Listado de Reservas</h2>
            {reservations.map(reservation => (
                <div key={reservation.id}>
                    <p>Hotel: {reservation.hotel}</p>
                    <p>Habitacion: {reservation.room}</p>
                    <p>Fecha: {reservation.date}</p>
                    <hr />
                    </div>
            ))}
        </div>
    );
};

ReservationList.propTypes = {
    hotel: PropTypes.object.isRequired,
};