import React, { useEffect, useState } from 'react';
import Navbar from './Navbar';
import HotelCard from './HotelCard';

const HotelList = () => {
    const [hotels, setHotels] = useState([]);

    useEffect(() => {
        fetchHotels();
    }, []);

    const fetchHotels = async () => {
        try {
            const response = await fetch('/hotels/gethotels');
            const data = await response.json();
            setHotels(data);
        } catch (error) {
            console.error('Error fetching hotels:', error.message);
        }
    };

    return (
        <div>
            <div className="hotel-list">
                {hotels.map((hotel) => (
                    <HotelCard key={hotel.id} hotel={hotel} />
                ))}
            </div>
        </div>
    );
};

export default HotelList;
