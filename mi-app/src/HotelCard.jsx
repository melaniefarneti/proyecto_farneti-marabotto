import React, { useState, useEffect } from "react";

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
  const handleReservarClick = () => {
    // Redirigir a la página de reservas del hotel
    window.location.href = `/reservations/${hotel.ID}`;
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
        <button onClick={handleReservarClick} style={buttonStyle}>
          Ver mas
        </button>
        </div>
      </div>
    </div>
  );
}

export default HotelCard;


/*import React, { useState, useEffect } from "react";

function HotelCard({ hotel }) {
  const [activeTab, setActiveTab] = useState(0);
  const [amenities, setAmenities] = useState([]);
  const [amenitiesLoaded, setAmenitiesLoaded] = useState(false);
  const [photos, setPhotos] = useState([]);
  const [photosLoaded, setPhotosLoaded] = useState(false);

  const cardStyle = {
    backgroundColor: "black",
    width: "250px",
    margin: "10px",
    color: "white",
    marginRight: "10px",
    marginLeft: "10px"
  };

  const imageStyle = {
    maxHeight: "200px",
    objectFit: "cover",
    marginBottom: "8px"
  };

  const titleStyle = {
    padding: "8px",
    backgroundColor: "rgba(0, 0, 0, 0.5)",
    fontWeight: "bold",
    fontFamily: "inherit",
    textAlign: "left",
    margin: 0,
    fontSize: "24px"
  };

  const contentStyle = {
    padding: "8px"
  };

  const tabStyle = {
    display: "flex",
    justifyContent: "space-between",
    marginBottom: "10px"
  };

  const tabButtonStyle = {
    fontSize: "9px",
    padding: "4px 8px",
    margin: "0 3px"
  };

  const buttonStyle = {
    fontSize: "13px",
    marginTop: "15px",
    marginLeft: "3px"
  };

  const handleTabClick = (tabIndex) => {
    setActiveTab(tabIndex);
  };

  const handleReservarClick = () => {
    // Redirigir a la página de reservas del hotel
    window.location.href = `/reservations/${hotel.ID}`;
  };

  return (

      <div style={cardStyle}>
        <div>
          <img src={hotel.Photo} alt="Hotel" style={imageStyle} />
        </div>
        <div>
          <div style={titleStyle}>{hotel.Name}</div>
          <div style={contentStyle}>
            <div style={tabStyle}><p>{hotel.Description}</p></div>
        </div>
        <div>
          <button onClick={handleReservarClick} style={buttonStyle}>
            Ver mas
          </button>
        </div>
      </div>
    </div>
  );
}

export default HotelCard;*/
