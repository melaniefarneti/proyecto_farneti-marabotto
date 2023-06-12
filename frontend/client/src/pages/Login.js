import React, { useState } from 'react';
import './Login.css';
import { Link } from "react-router-dom";
import Navbar from "./Navbar";

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleUsernameChange = (event) => {
        setUsername(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();

        // Aquí puedes realizar la lógica de autenticación y enviar los datos al servidor
        // utilizando una API o realizar cualquier otra acción necesaria

        // Restablecer los valores del formulario después del envío
        setUsername('');
        setPassword('');
    };

    return (
<div className="Navbar">
    <Navbar/>
        <div id="contenedor">
            <div id="central">
                <h2 className="titulo">Login</h2>
                <form id="login" onSubmit={handleSubmit}>
                    <div>
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={handleUsernameChange}
                        />
                    </div>
                    <div>
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={handlePasswordChange}
                        />
                    </div>
                    <button type="submit">Submit</button>
                </form>
                <div className="inferior">
                    <Link to="/Register">Create an Account</Link>
                </div>
            </div>
        </div>
</div>
    );
};

export default Login;
