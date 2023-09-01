import React from 'react';
import logo from "./logo.png";
import './styles.css'; // Asegúrate de importar tu archivo de estilos

const headerStyles = {
  backgroundColor: 'black',
  padding: '10px 0',
  width: '100%',
  position: 'fixed',
  top: 0,
  zIndex: 1000,
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
};

const logoStyles = {
  maxWidth: '120px',
  display: 'block',
  margin: '0 auto',
};

const buttonContainerStyles = {
  display: 'flex',
};

const buttonStyles = {
  display: 'inline-block',
  padding: '5px 10px',
  backgroundColor: 'black',
  color: 'white',
  textDecoration: 'none',
  borderRadius: '20px',
  transition: 'background-color 0.3s',
  margin: '0 5px',
  border: '1.8px solid white',
}

function Header() {
  return (
    <header style={headerStyles}>
      <div>
        <a href="/"><img src={logo} alt="Logo" style={logoStyles} /></a>
      </div>
      <div className="buttons" style={buttonContainerStyles}>
        <a href="/login" className="button" style={buttonStyles}>Login</a>
        <a href="/register" className="button" style={buttonStyles}>Register</a>
      </div>
    </header>
  );
}

export { Header };
