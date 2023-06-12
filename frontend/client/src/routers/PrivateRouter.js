import React from 'react';
import { Navigate, Route } from 'react-router-dom';
import { UserContext } from '../contexts/UserContext'; // Actualizado el import

const PrivateRoute = ({ component: Component, ...rest }) => {
    const { user } = React.useContext(UserContext); // Actualizado el uso del contexto

    return (
        <Route
            {...rest}
            render={(props) =>
                user ? <Component {...props} /> : <Navigate to="/" />
            }
        />
    );
};

export default PrivateRoute;
