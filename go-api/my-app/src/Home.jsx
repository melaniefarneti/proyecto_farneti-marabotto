import React, { useState, useEffect } from "react";
import HotelCard from "./HotelCard";

function Home() {
  const [hotels, setHotels] = useState([]);

  const fetchHotels = async () => {
    try {
      const response = await fetch("http://localhost:8080/hotels/gethotels");
      const data = await response.json();
      setHotels(data);
    } catch (error) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchHotels();
  }, []);

  return (
    <div>
      <h1>Home</h1>
      <div className="card-container">
        {hotels.map((hotel) => (
          <HotelCard key={hotel.ID} hotel={hotel} />
        ))}
      </div>
    </div>
  );
}

export default Home;
