import React, { useState, useEffect } from "react";
import { BrowserRouter as Router} from 'react-router-dom';
import Home from './Home';
import {Header} from "./Header";
import './Styles.css';


function App() {
  return (
    <>
      <div>
        <Header />
      </div>
      <Router>
        <Home path="/" />
      </Router>
    </>
  );
}

export default App;
