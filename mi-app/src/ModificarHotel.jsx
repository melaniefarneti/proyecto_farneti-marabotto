import React, { useState } from "react";
import { useParams } from "react-router-dom";

function ModificarHotel() {
  const { hotelId } = useParams();

  const [amenity, setAmenity] = useState("");
  const [photo, setPhoto] = useState(null);
  const [successMessage, setSuccessMessage] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const handleAmenityChange = (event) => {
    setAmenity(event.target.value);
  };

  const handlePhotoChange = (event) => {
    const file = event.target.files[0];
    setPhoto(file);
  };

  const handleAmenitySubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch(`http://localhost:8080/amenities`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          nombre: amenity,
          hotel_id: parseInt(hotelId),
        }),
      });

      if (response.ok) {
        setSuccessMessage("Amenity cargado exitosamente.");
        setAmenity("");
      } else {
        const data = await response.json();
        console.error("Error:", response.status, data);
        setErrorMessage(data.error || "Error al cargar el amenity.");
      }
    } catch (error) {
      console.error(error);
      setErrorMessage("Error al cargar el amenity. Por favor, inténtalo de nuevo.");
    }
  };

  const handlePhotoSubmit = async (event) => {
    event.preventDefault();

    try {
      const formData = new FormData();
      formData.append("hotel_id", hotelId);
      formData.append("photo", photo);

      await fetch(`http://localhost:8080/hotels/uploadphoto/${hotelId}`, {
        method: "POST",
        body: formData,
      });

      setSuccessMessage("La foto se cargó exitosamente.");
    } catch (error) {
      console.error(error);
      setErrorMessage("Error al cargar la foto. Por favor, inténtalo de nuevo.");
    }
  };

  return (
    <div>
      <h2>Modificar Hotel</h2>
      {successMessage && <p>{successMessage}</p>}
      {errorMessage && <p>{errorMessage}</p>}

      <form onSubmit={handleAmenitySubmit}>
        <h3>Cargar Amenities</h3>
        <div>
          <label htmlFor="amenity">Nombre del amenity:</label>
          <input
            type="text"
            id="amenity"
            value={amenity}
            onChange={handleAmenityChange}
          />
        </div>
        <button type="submit">Cargar Amenity</button>
      </form>

      <form onSubmit={handlePhotoSubmit}>
        <h3>Cargar Fotos</h3>
        <div>
          <label htmlFor="photo">Seleccione una foto:</label>
          <input
            type="file"
            id="photo"
            onChange={handlePhotoChange}
          />
        </div>
        <button type="submit">Cargar Foto</button>
      </form>
    </div>
  );
}

export default ModificarHotel;
