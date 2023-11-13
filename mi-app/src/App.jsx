import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom"; // Importa Routes
import Home from "./Home";
import { Header } from "./Header";
import HotelDetailPage from "./HotelDetailPage";
import Login from "./LoginPage";
import Register from "./RegisterPage";
import AdminPage from "./AdminPage";
import ReservationPage from "./ReservationPage";
import ModificarHotel from "./ModificarHotel";
import InformationPage from "./InformationPage";


function App() {
  return (
    <>
      <div>
        <Header />
      </div>
      <Router>
        <Routes> 
          <Route path="/" element={<Home />} />
          <Route path="/hotels/:hotelId" element={<HotelDetailPage />} />
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
          <Route path="admin" element={<AdminPage/>} />
          <Route path='reservations/:hotelId' element={<ReservationPage/>}/>
          <Route path="/modificarhotel/:hotelId" element={<ModificarHotel/>}/>
          <Route path="/information" element={<InformationPage/>}/>
        </Routes>
      </Router>
    </>
  );
}

export default App;


/*import React, { useState, useEffect } from "react";
import { BrowserRouter as Router} from 'react-router-dom';
import Home from './Home';
import {Header} from "./Header";
import './Styles.css';


function App() {
  return (
    <>
      <div>
        <Header />
      </div>
      <Router>
        <Home path="/" />
      </Router>
    </>
  );
}

export default App;*/
