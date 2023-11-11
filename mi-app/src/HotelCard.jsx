import React from "react";
import { Link } from "react-router-dom";

const cardStyle = {
  backgroundColor: "black",
  width: "250px",
  margin: "10px",
  color: "white",
  marginRight: "10px",
  marginLeft: "10px",
  display: "flex",
  flexDirection: "column",
  alignItems: "center", // Centra horizontalmente
};

const imageStyle = {
  width: "100%", // Ajusta el ancho al 100%
  maxHeight: "200px",
  objectFit: "cover",
  marginBottom: "8px",
};

const titleStyle = {
  padding: "8px",
  backgroundColor: "rgba(0, 0, 0, 0.5)",
  fontWeight: "bold",
  fontFamily: "inherit",
  textAlign: "center",
  margin: 0,
  fontSize: "24px",
};

const contentStyle = {
  padding: "8px",
  textAlign: "center", // Centra horizontalmente el contenido
};

const buttonContainerStyle = {
  padding: "10px",
  display: "flex",
  justifyContent: "center", // Centra horizontalmente
};

const buttonStyle = {
  fontSize: "13px",
  marginTop: "15px",
  alignItems: "center",
};

function HotelCard({ hotel }) {
  return (
    <div style={cardStyle}>
      <img src={hotel.Photo} alt="Hotel" style={imageStyle} />
      <div>
        <div style={titleStyle}>{hotel.Name}</div>
        <div style={contentStyle}>
          <p>{hotel.Description}</p>
        </div>
        <div style={buttonContainerStyle}>
          <Link to={`/hotels/${hotel.ID}`}>
            <button style={buttonStyle}>Ver más</button>
          </Link>
          <div style={{ margin: "10px 10px" }}></div>
          <Link to={`/reservations/${hotel.ID}`}>
            <button style={buttonStyle}>Reservar</button>
          </Link>
        </div>
      </div>
    </div>
  );
}

export default HotelCard;


/*import React, { useState, useEffect } from "react";

const cardStyle = {
  backgroundColor: "black",
  width: "250px",
  margin: "10px",
  color: "white",
  marginRight: "10px",
  marginLeft: "10px",
  display: "flex",
  flexDirection: "column",
  alignItems: "center", // Centra horizontalmente
};

const imageStyle = {
  width: "100%", // Ajusta el ancho al 100%
  maxHeight: "200px",
  objectFit: "cover",
  marginBottom: "8px",
};

const titleStyle = {
  padding: "8px",
  backgroundColor: "rgba(0, 0, 0, 0.5)",
  fontWeight: "bold",
  fontFamily: "inherit",
  textAlign: "center",
  margin: 0,
  fontSize: "24px",
};

const contentStyle = {
  padding: "8px",
  textAlign: "center", // Centra horizontalmente el contenido
};

const buttonContainerStyle = {
  padding: "10px",
  display: "flex",
  justifyContent: "center", // Centra horizontalmente
};

const buttonStyle = {
  fontSize: "13px",
  marginTop: "15px",
  alignItems: "center",
};

function HotelCard({ hotel }) {
  const handleHotelDetailClick = () => {
    window.location.href = `/hotels/${hotel.ID}`;
  };

  return (
    <div style={cardStyle}>
      <img src={hotel.Photo} alt="Hotel" style={imageStyle} />
      <div>
        <div style={titleStyle}>{hotel.Name}</div>
        <div style={contentStyle}>
          <p>{hotel.Description}</p>
        </div>
        <div style={buttonContainerStyle}>
        <button onClick={handleHotelDetailClick} style={buttonStyle}>
          Ver mas
        </button>
        </div>
      </div>
    </div>
  );
}

export default HotelCard;*/