import React from "react";

function HotelCardAdmin({ hotel }) {
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
    fontSize: "24px" // Tamaño de la letra del título
  };

  const contentStyle = {
    padding: "8px"
  };

  const handleModificarClick = () => {
    // Redirigir a la página de modificación del hotel
    window.location.href = `/modificarhotel/${hotel.ID}`;
  };

  return (
    <div className="col s12 m6 l4">
      <div className="card" style={cardStyle}>
        <div className="card-image">
          <img src={hotel.Photo} alt="Hotel" style={imageStyle} />
        </div>
        <div className="card-content">
          <div style={titleStyle}>{hotel.Name}</div>
          <div style={contentStyle}>
            <p>{hotel.Description}</p>
          </div>
        </div>
        <div className="card-action">
          <button onClick={handleModificarClick}>Modificar</button>
        </div>
      </div>
    </div>
  );
}

export default HotelCardAdmin;
