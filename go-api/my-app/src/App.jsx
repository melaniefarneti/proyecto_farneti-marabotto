import React, { useState, useEffect } from "react";
import { Router } from "@reach/router";
import Home from "./Home";
import Header from "./Header";
import Footer from "./Footer";
import Login from "./Login";
import './Styles.css';

function App() {
  const [hotels, setHotels] = useState([]);

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const response = await fetch("http://localhost:8080/hotels/gethotels");
        const data = await response.json();
        setHotels(data);
      } catch (error) {
        console.error(error);
      }
    };

    fetchHotels();
  }, []);

  return (
    <>
      <div>
        <Header />
      </div>
      <Router>
        <Home path="/" hotels={hotels} />
        <Login path="/login" />
      </Router>
      <Footer />
    </>
  );
}

export default App;
