import React, { useState, useEffect } from "react";

function HotelCardAdmin({ hotel }) {
  const [activeTab, setActiveTab] = useState(0);
  const [amenities, setAmenities] = useState([]);
  const [amenitiesLoaded, setAmenitiesLoaded] = useState(false);
  const [photos, setPhotos] = useState([]);
  const [photosLoaded, setPhotosLoaded] = useState(false);

  const cardStyle = {
    backgroundColor: "black",
    width: "250px",
    margin: "10px",
    color: "white"
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
    padding: "8px",
    marginLeft: 0, // Elimina el margen interno a la izquierda
    marginRight: 0 // Elimina el margen interno a la derecha
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
    marginLeft: "3px" // Ajusta el margen izquierdo del segundo botón
  };

  const handleTabClick = (tabIndex) => {
    setActiveTab(tabIndex);
  };

  const handleModificarClick = () => {
    // Redirigir a la página de modificación del hotel
    window.location.href = `/modificarhotel/${hotel.ID}`;
  };

  const handleReservarClick = () => {
    // Redirigir a la página de reservas del hotel
    window.location.href = `/reservations/${hotel.ID}`;
  };

  const fetchAmenities = async () => {
    try {
      const response = await fetch(`http://localhost:8080/getamenities/${hotel.ID}`);
      if (response.ok) {
        const data = await response.json();
        setAmenities(data);
        setAmenitiesLoaded(true);
      } else {
        throw new Error("Request failed with status code " + response.status);
      }
    } catch (error) {
      console.log(error);
    }
  };

  const fetchHotelPhotos = async () => {
    try {
      const response = await fetch(`http://localhost:8080/gethotelphoto/${hotel.ID}`);
      if (response.ok) {
        const data = await response.json();
        setPhotos(data);
        setPhotosLoaded(true);
      } else {
        throw new Error("Request failed with status code " + response.status);
      }
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    if (activeTab === 1 && !amenitiesLoaded) {
      fetchAmenities();
    } else if (activeTab === 2 && !photosLoaded) {
      fetchHotelPhotos();
    }
  }, [activeTab, amenitiesLoaded, photosLoaded]);

  return (
    <div className="col s12 m6 l4">
      <div className="card" style={cardStyle}>
        <div className="card-image">
          <img src={hotel.Photo} alt="Hotel" style={imageStyle} />
        </div>
        <div className="card-content">
          <div style={titleStyle}>{hotel.Name}</div>
          <div className="card-action" style={contentStyle}>
            <div style={tabStyle}>
              <button
                onClick={() => handleTabClick(0)}
                className={activeTab === 0 ? "active" : ""}
                style={tabButtonStyle}
              >
                Descripción
              </button>
              <button
                onClick={() => handleTabClick(1)}
                className={activeTab === 1 ? "active" : ""}
                style={tabButtonStyle}
              >
                Amenidades
              </button>
              <button
                onClick={() => handleTabClick(2)}
                className={activeTab === 2 ? "active" : ""}
                style={tabButtonStyle}
              >
                Fotos Cargadas
              </button>
            </div>
            {activeTab === 0 && <p>{hotel.Description}</p>}
            {activeTab === 1 && (
              <div>
                {amenitiesLoaded ? (
                  amenities.map((amenity) => (
                    <p key={amenity.ID}>{amenity.Nombre}</p>
                  ))
                ) : (
                  <p className="loader">Cargando amenidades...</p>
                )}
              </div>
            )}
            {activeTab === 2 && (
              <div>
                {photosLoaded ? (
                  photos.map((photo) => (
                    <img key={photo.ID} src={photo.URL} alt="Hotel Photo" />
                  ))
                ) : (
                  <p className="loader">Cargando fotos...</p>
                )}
              </div>
            )}
            <button onClick={handleReservarClick} style={buttonStyle}>
              Reservar
            </button>
            <button onClick={handleModificarClick} style={buttonStyle}>
              Modificar
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default HotelCardAdmin;
