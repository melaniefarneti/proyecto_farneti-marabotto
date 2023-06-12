import React, { useState } from 'react';
import Navbar from "./Navbar";
import './Register.css'

const Register = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');

    const handleEmailChange = (e) => {
        setEmail(e.target.value);
    };

    const handlePasswordChange = (e) => {
        setPassword(e.target.value);
    };

    const handleConfirmPasswordChange = (e) => {
        setConfirmPassword(e.target.value);
    };

    const handleRegister = async (e) => {
        e.preventDefault();
        // Call the backend API to register the user
        try {
            const response = await fetch('/api/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email,
                    password,
                }),
            });

            if (response.ok) {
                console.log('User registered successfully');
                // Reset form fields
                setEmail('');
                setPassword('');
                setConfirmPassword('');
            } else {
                console.log('Registration failed');
            }
        } catch (error) {
            console.log('Error:', error);
        }
    };

    return (
        <div className="Navbar">
            <Navbar/>
        <div className="center-align">
            <div className="titulo">
            <h3>Register</h3>
            </div>
            <div className="body">
            <form onSubmit={handleRegister}>
                <div className="form-group">
                    <label htmlFor="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        value={email}
                        onChange={handleEmailChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={handlePasswordChange}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="confirmPassword">Confirm Password:</label>
                    <input
                        type="password"
                        id="confirmPassword"
                        value={confirmPassword}
                        onChange={handleConfirmPasswordChange}
                        required
                    />
                </div>
                <button type="submit">Register</button>
            </form>
            </div>
        </div>
        </div>
    );
};

export default Register;
