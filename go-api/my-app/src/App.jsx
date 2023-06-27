import React, { useState, useEffect } from "react";
import HotelCard from "./HotelCard";
import Header from "./Header";
import Footer from "./Footer";
import './Styles.css';

function App() {
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
    <>
    <div> <Header /></div>
    <div className="card-container">
      {hotels.map((hotel) => (
        <HotelCard key={hotel.ID} hotel={hotel} />
      ))}
    </div>
    <Footer />
    </>
  );
}

export default App;

