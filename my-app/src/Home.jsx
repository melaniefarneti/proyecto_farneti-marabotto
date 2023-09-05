import React, { useState, useEffect } from "react";
import HotelCard from "../../my-app/src/HotelCard";

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
      <div className="card-container">
        {hotels.map((hotel) => (
          <HotelCard key={hotel.ID} hotel={hotel} />
        ))}
      </div>
    </div>
  );
}

export default Home;
