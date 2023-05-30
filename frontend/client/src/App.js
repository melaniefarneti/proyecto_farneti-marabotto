import logo from './logo.svg';
import './App.css';

import Hotel from './components/Hotel';

import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import HotelList from "./pages/HotelList";
import Booking from "./pages/Booking";

function App() {
  return (
    <div className="App">
      <h1>Hoteles</h1>
      <Hotel />
      <Router>
          <Routes>
              <Route path='/hotel-list' element={<HotelList />}></Route>
              <Route path='/booking' element={<Booking />}></Route>
          </Routes>
      </Router>
      <header className="App-header">
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
