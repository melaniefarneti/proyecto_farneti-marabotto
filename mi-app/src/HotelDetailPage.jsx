import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import "./styles.css";

function HotelDetailPage() {
  const { hotelId } = useParams(); // Usar useParams para obtener el hotelId de la URL

  const [hotel, setHotel] = useState({});
  const [amenities, setAmenities] = useState([]);
  const [photos, setPhotos] = useState([]);

  useEffect(() => {
    if (hotelId) {
      const fetchHotelDetails = async () => {
        try {
          const response = await fetch(`http://localhost:8080/hotels/${hotelId}`);
          if (response.ok) {
            const data = await response.json();
            setHotel(data);
          } else {
            throw new Error("Request failed with status code " + response.status);
          }
        } catch (error) {
          console.log(error);
        }
      };

      const fetchAmenities = async () => {
        try {
          const response = await fetch(`http://localhost:8080/getamenities/${hotelId}`);
          if (response.ok) {
            const data = await response.json();
            setAmenities(data);
          } else {
            throw new Error("Request failed with status code " + response.status);
          }
        } catch (error) {
          console.log(error);
        }
      };

      const fetchHotelPhotos = async () => {
        try {
          const response = await fetch(`http://localhost:8080/gethotelphoto/${hotelId}`);
          if (response.ok) {
            const data = await response.json();
            setPhotos(data);
          } else {
            throw new Error("Request failed with status code " + response.status);
          }
        } catch (error) {
          console.log(error);
        }
      };
      

      fetchHotelDetails();
      fetchAmenities();
      fetchHotelPhotos();
    }
  }, [hotelId]);

  return (
    <div>
      <div className="hotel-container">
        <img src={hotel.Photo} alt={hotel.Name} style={{ width: '100%' }} />
        <h2 className="hotel-name" style={{ fontSize: '50px' }}>{hotel.Name} </h2>
        <p className="hotel-description">{hotel.Description}</p>
      

      <div className="amenities-container">
        <h3>Amenidades:</h3>
        <ul>
          {amenities.map((amenity) => (
            <li key={amenity.ID} className="amenity">{amenity.Nombre}</li>
          ))}
        </ul>
      </div>

      <div className="photos-container">
        <h3>Fotos del Hotel:</h3>
        <div className="photos">
        {photos.length > 0 ? (
          photos.map((photo) => (
          <img key={photo.ID} src={`${photo.Filename}`} alt="Hotel Photo" />
        ))
      ) : (
        <p>No hay fotos cargadas para este hotel.</p>
        )}
        </div>
      </div>
    </div>
    </div>
  );
}

export default HotelDetailPage;

