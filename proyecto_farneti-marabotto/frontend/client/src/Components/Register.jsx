import React, { useState } from 'react';

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleNameChange = event => {
        setName(event.target.value);
    };

    const handleEmailChange = event => {
        setEmail(event.target.value);
    };

    const handlePasswordChange = event => {
        setPassword(event.target.value);
    };

    const handleRegister = () => {
        // Aquí puedes realizar la llamada al backend para registrar al usuario con los datos ingresados
        console.log('Nombre:', name);
        console.log('Email:', email);
        console.log('Contraseña:', password);
    };

    return (
        <div>
            <h2>Registro de Usuario</h2>
            <div>
                <label htmlFor="name">Nombre:</label>
                <input type="text" id="name" value={name} onChange={handleNameChange} />
            </div>
            <div>
                <label htmlFor="email">Email:</label>
                <input type="email" id="email" value={email} onChange={handleEmailChange} />
            </div>
            <div>
                <label htmlFor="password">Contraseña:</label>
                <input type="password" id="password" value={password} onChange={handlePasswordChange} />
            </div>
            <button onClick={handleRegister}>Registrarse</button>
        </div>
    );
};

export default Register;