import React, { useState, useEffect } from 'react';

const AdminDashboard = () => {
    const [hotels, setHotels] = useState([]);
    const [reservations, setReservations] = useState([]);
    const [filteredHotels, setFilteredHotels] = useState([]);
    const [filteredReservations, setFilteredReservations] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');

    // Lógica para cargar los hoteles y las reservas desde el backend
    useEffect(() => {
        // Aquí puedes hacer una petición al backend para obtener los hoteles y las reservas
        // y luego actualizar los estados de `hotels` y `reservations` con los datos recibidos
    }, []);

    // Filtrar hoteles y reservas según el término de búsqueda
    useEffect(() => {
        const filteredHotels = hotels.filter((hotel) =>
            hotel.title.toLowerCase().includes(searchTerm.toLowerCase())
        );
        const filteredReservations = reservations.filter((reservation) =>
            reservation.customerName.toLowerCase().includes(searchTerm.toLowerCase())
        );
        setFilteredHotels(filteredHotels);
        setFilteredReservations(filteredReservations);
    }, [hotels, reservations, searchTerm]);

    const handleDeleteHotel = (hotelId) => {
        // Lógica para eliminar el hotel con el ID proporcionado desde el backend
        // Actualiza el estado de los hoteles después de la eliminación
    };

    const handleDeleteReservation = (reservationId) => {
        // Lógica para eliminar la reserva con el ID proporcionado desde el backend
        // Actualiza el estado de las reservas después de la eliminación
    };

    return (
        <div>
            <h2>Interfaz Administrativa</h2>

            <input
                type="text"
                placeholder="Buscar..."
                value={searchTerm}
                onChange={(e) => setSearchTerm(e.target.value)}
            />

            <h3>Hoteles</h3>
            {filteredHotels.length === 0 ? (
                <p>No se encontraron hoteles.</p>
            ) : (
                <ul>
                    {filteredHotels.map((hotel) => (
                        <li key={hotel.id}>
                            <h4>{hotel.title}</h4>
                            <p>{hotel.description}</p>
                            {/* Otras propiedades del hotel */}
                            <button onClick={() => handleDeleteHotel(hotel.id)}>
                                Eliminar Hotel
                            </button>
                        </li>
                    ))}
                </ul>
            )}

            <h3>Reservas</h3>
            {filteredReservations.length === 0 ? (
                <p>No se encontraron reservas.</p>
            ) : (
                <ul>
                    {filteredReservations.map((reservation) => (
                        <li key={reservation.id}>
                            <p>Fecha: {reservation.date}</p>
                            <p>Cliente: {reservation.customerName}</p>
                            {/* Otras propiedades de la reserva */}
                            <button onClick={() => handleDeleteReservation(reservation.id)}>
                                Eliminar Reserva
                            </button>
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default AdminDashboard;