import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
    const [hotels, setHotels] = useState([]);
    const [selectedHotel, setSelectedHotel] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');

    useEffect(() => {
        // Aquí se realiza la llamada al backend para obtener los hoteles disponibles
        fetch('/api/hotels')
            .then(response => response.json())
            .then(data => setHotels(data))
            .catch(error => console.log(error));
    }, []);

    const handleHotelChange = event => {
        setSelectedHotel(event.target.value);
    };

    const handleStartDateChange = event => {
        setStartDate(event.target.value);
    };

    const handleEndDateChange = event => {
        setEndDate(event.target.value);
    };

    const handleSearch = () => {
        // Aquí puedes redirigir a la página de detalle del hotel seleccionado con las fechas elegidas
        console.log('Hotel seleccionado:', selectedHotel);
        console.log('Fecha de inicio:', startDate);
        console.log('Fecha de fin:', endDate);
    };

    return (
        <div>
            <h1>Bienvenido al sitio de reserva de habitaciones</h1>
            <div>
                <label htmlFor="hotel">Seleccione un hotel:</label>
                <select id="hotel" value={selectedHotel} onChange={handleHotelChange}>
                    <option value="">Seleccione...</option>
                    {hotels.map(hotel => (
                        <option key={hotel.id} value={hotel.id}>
                            {hotel.title}
                        </option>
                    ))}
                </select>
            </div>
            <div>
                <label htmlFor="startDate">Fecha desde:</label>
                <input type="date" id="startDate" value={startDate} onChange={handleStartDateChange} />
            </div>
            <div>
                <label htmlFor="endDate">Fecha hasta:</label>
                <input type="date" id="endDate" value={endDate} onChange={handleEndDateChange} />
            </div>
            <button onClick={handleSearch}>Buscar</button>
            <Link to="/register">Registrarse</Link>
        </div>
    );
};

export default Home;