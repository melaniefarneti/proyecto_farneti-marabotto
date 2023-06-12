import React, { createContext, useState, useEffect } from 'react';

const UserContext = createContext();

const UserProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [isAdmin, setIsAdmin] = useState(false);

    useEffect(() => {
        // Lógica para obtener los datos del usuario actual
        const fetchUserData = async () => {
            try {
                // Realizar una llamada a la API para obtener los datos del usuario
                const response = await fetch('/api/user');
                const userData = await response.json();

                // Establecer el usuario y el estado de administrador
                setUser(userData);
                setIsAdmin(userData.role === 'admin');
            } catch (error) {
                console.error(error);
            }
        };

        // Llamar a la función para obtener los datos del usuario
        fetchUserData();
    }, []);

    return (
        <UserContext.Provider value={{ user, isAdmin }}>
            {children}
        </UserContext.Provider>
    );
};

export { UserContext, UserProvider };
