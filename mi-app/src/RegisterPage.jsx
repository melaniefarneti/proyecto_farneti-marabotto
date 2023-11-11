import React, { useState } from "react";

function Register() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [userExistsError, setUserExistsError] = useState("");

  const handleNameChange = (e) => {
    setName(e.target.value);
  };

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name, email, password }),
        credentials: "include",
        mode: "cors",
      });

      if (response.ok) {
        // Si la respuesta es exitosa, redirige al usuario a la página deseada
        window.location.href = "/";
      } else {
        // Si la respuesta no es exitosa, muestra un mensaje de error
        const data = await response.json();
        if (response.status === 400 && data.error === "user already exists") {
          setUserExistsError("El usuario ya existe. Por favor, use otro email.");
        } else {
          setError("Error al registrar. Por favor, inténtalo de nuevo.");
        }
      }
    } catch (error) {
      // Maneja cualquier error de red u otro error
      setError("Error al registrar. Por favor, inténtalo de nuevo.");
    }
  };

  return (
    <div className="login-container">
      <h1>Registrarse</h1>
      <form className="register-form" onSubmit={handleSubmit}>
        <div>
          <label>Nombre:</label>
          <input type="text" value={name} onChange={handleNameChange} />
        </div>
        <div>
          <label>Email:</label>
          <input type="email" value={email} onChange={handleEmailChange} />
        </div>
        <div>
          <label>Contraseña:</label>
          <input type="password" value={password} onChange={handlePasswordChange} />
        </div>
        {error && <p className="register-error">{error}</p>}
        {userExistsError && <p className="register-error">{userExistsError}</p>}
        <button type="submit">Registrarse</button>
      </form>
    </div>
  );
}

export default Register;
