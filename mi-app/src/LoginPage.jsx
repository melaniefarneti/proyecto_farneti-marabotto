import React, { useState } from "react";

function Login() {
  const [email, setEmail] = useState("");
  const [contrasena, setContrasena] = useState("");
  const [error, setError] = useState("");

  const handleEmailChange = (e) => {
    setEmail(e.target.value);
  };

  const handleContrasenaChange = (e) => {
    setContrasena(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, contrasena }),
        credentials: "include",
        mode: "cors",
      });

      if (response.ok) {
        // Si la respuesta es exitosa, redirige al usuario a la página deseada
        const userResponse = await fetch(`http://localhost:8080/users/emailuser/${email}`);
        const user = await userResponse.json();

        if (user.Role === "administrador") {
          window.location.href = "/admin";
        } else if (user.Role === "cliente") {
          window.location.href = "/";
        }
      } else {
        // Si la respuesta no es exitosa, muestra un mensaje de error
        setError("Error de inicio de sesión. Por favor, inténtalo de nuevo.");
      }
    } catch (error) {
      // Maneja cualquier error de red u otro error
      setError("Error de inicio de sesión. Por favor, inténtalo de nuevo.");
    }
  };

  return (
    <div className="login-container">
      <h1>Iniciar sesión</h1>
      <form className="login-form" onSubmit={handleSubmit}>
        <div>
          <label>Email:</label>
          <input type="email" value={email} onChange={handleEmailChange} />
        </div>
        <div>
          <label>Contraseña:</label>
          <input type="password" value={contrasena} onChange={handleContrasenaChange} />
        </div>
        {error && <p className="login-error">{error}</p>}
        <button type="submit">Iniciar sesión</button>
      </form>
    </div>
  );
}

export default Login;
