import React, { createContext, useState } from "react";


export const UserProvider = ({ children }) => {
    const [user, setUser] = useState(null);

    const logout = () => {
        // Lógica para cerrar sesión y actualizar el estado del usuario
        setUser(null);
    };

    return (
        <UserContext.Provider value={{ user, setUser, logout }}>
            {children}
        </UserContext.Provider>
    );
};
