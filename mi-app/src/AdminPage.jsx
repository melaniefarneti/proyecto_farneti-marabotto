import React, { useState, useEffect } from "react";
import HotelCardAdmin from "./HotelCardAdmin";

function AdminPage() {
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
      <div>
      <a href="/information" style={{ textDecoration: "none" }}>
          <button>
            Ver listado de reservas
          </button>
        </a>
      </div>
      <div className="card-container">
        {hotels.map((hotel) => (
          <HotelCardAdmin key={hotel.ID} hotel={hotel} />
        ))}
      </div>
    </div>
  );
}

export default AdminPage;
