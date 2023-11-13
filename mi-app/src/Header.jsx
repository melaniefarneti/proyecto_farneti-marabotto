import React from 'react';
import logo from "./logo.png";
import './Styles.css'; // Asegúrate de importar tu archivo de estilos

const headerStyles = {
  backgroundColor: 'black',
  padding: '10px',
  width: '100%',
  position: 'fixed',
  top: 0,
  zIndex: 1000,
  display: 'flex',
  justifyContent: 'space-between',
  alignItems: 'center',
  flexWrap: 'wrap', // Permite que los elementos se envuelvan si es necesario
  boxSizing: 'border-box', // Incluye padding en el ancho total
};

const logoStyles = {
  maxWidth: '120px', // Logotipo ajustará automáticamente su ancho
  height: 'auto', // Mantendrá la relación de aspecto
  margin: '0 auto', // Centra horizontalmente
};

const buttonContainerStyles = {
  display: 'flex',
  alignItems: 'center', // Centra verticalmente los botones
  gap: '10px', // Espacio entre los botones
};

const buttonStyles = {
  padding: '5px 10px',
  color: 'white',
  fontWeight: 'bold',
  textDecoration: 'none',
  borderRadius: '20px',
  transition: 'background-color 0.3s',
};

function Header() {
  return (
    <header style={headerStyles}>
      <a href="/"><img src={logo} alt="Logo" style={logoStyles} /></a>
      <div className="buttons" style={buttonContainerStyles}>
        <a href="/login" className="button" style={buttonStyles}>Login</a>
        <a href="/register" className="button" style={buttonStyles}>Register</a>
      </div>
    </header>
  );
}

export { Header };
