import React from "react";
import logo from "./logo.png";

function Header() {
  return (
    <nav>
      <div className="nav-wrapper">
        <a href="#" className="brand-logo">
          <img src={logo} alt="Logo" className="logo" />
          <span className="brand-text">HiHotel!</span>
        </a>
        <ul id="nav-mobile" className="right hide-on-med-and-down">
          <li>
            <a href="/login">Login</a>
          </li>
          <li>
            <a href="badges.html">Register</a>
          </li>
        </ul>
      </div>
    </nav>
  );
}

export default Header;
