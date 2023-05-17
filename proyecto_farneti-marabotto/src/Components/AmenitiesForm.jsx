import React, { useState } from 'react';

const AmenityForm = () => {
    const [amenities, setAmenities] = useState([{ name: '', description: '' }]);

    const handleInputChange = (index, event) => {
        const { name, value } = event.target;
        const updatedAmenities = [...amenities];
        updatedAmenities[index][name] = value;
        setAmenities(updatedAmenities);
    };

    const handleAddAmenity = () => {
        setAmenities([...amenities, { name: '', description: '' }]);
    };

    const handleRemoveAmenity = index => {
        const updatedAmenities = [...amenities];
        updatedAmenities.splice(index, 1);
        setAmenities(updatedAmenities);
    };

    const handleSubmit = event => {
        event.preventDefault();
        // Aquí puedes enviar los datos de las amenidades al backend para su almacenamiento
        console.log('Amenidades:', amenities);
    };

    return (
        <div>
            <h2>Carga de Amenidades</h2>
            <form onSubmit={handleSubmit}>
                {amenities.map((amenity, index) => (
                    <div key={index}>
                        <div>
                            <label htmlFor={`name-${index}`}>Nombre:</label>
                            <input
                                type="text"
                                id={`name-${index}`}
                                name="name"
                                value={amenity.name}
                                onChange={event => handleInputChange(index, event)}
                            />
                        </div>
                        <div>
                            <label htmlFor={`description-${index}`}>Descripción:</label>
                            <input
                                type="text"
                                id={`description-${index}`}
                                name="description"
                                value={amenity.description}
                                onChange={event => handleInputChange(index, event)}
                            />
                        </div>
                        {index > 0 && (
                            <button type="button" onClick={() => handleRemoveAmenity(index)}>
                                Eliminar Amenity
                            </button>
                        )}
                    </div>
                ))}
                <button type="button" onClick={handleAddAmenity}>
                    Agregar Amenity
                </button>
                <button type="submit">Guardar Amenidades</button>
            </form>
        </div>
    );
};

export default AmenityForm;