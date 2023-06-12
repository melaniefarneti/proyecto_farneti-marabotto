import React from "react";
import { Link } from "react-router-dom";
import "./Navbar.css";

const Navbar = () => {
    return (
        <div className="navbar">
            <div className="navContainer">
                <span className="Logo">HiHotel</span>
                <div className="navItems">
                    <Link to="/" className="NavLink">
                        Home
                    </Link>
                    <Link to="/Register" className="NavLink">
                        Register
                    </Link>
                    <Link to="/login" className="NavLink">
                        Login
                    </Link>
                </div>
            </div>
        </div>
    );
};

export default Navbar;
