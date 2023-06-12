import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { UserProvider } from "./contexts/UserContext.js";
import Header from "./pages/Header";
import Home from "./pages/Home";
import Hotel from "./pages/Hotel";
import List from "./pages/HotelList";
import Login from "./pages/Login";
import Register from "./pages/Register";

const App = () => {
    return (
        <BrowserRouter>
            <UserProvider>
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/Register" element={<Register/>} />
                    <Route path="/" element={<Home />} />
                    <Route path="/hotels" element={<List />} />
                    <Route path="/hotels/:id" element={<Hotel />} />
                </Routes>
            </UserProvider>
        </BrowserRouter>
    );
};

export default App;
