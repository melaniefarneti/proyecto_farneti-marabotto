import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { UserProvider } from '../contexts/UserContext';
import PrivateRoute from './PrivateRouter';
import Header from '../pages/Header';
import Home from '../pages/Home';
import HotelList from '../pages/HotelList';
import HotelDetail from '../pages/HotelDetail';
import ReservationForm from '../pages/reservation/ReservationForm';
import ReservationList from '../pages/ReservationList';

const AppRouter = () => {
    return (
        <Router>
            <UserProvider>
                <Header />
                    <Route exact path="/" component={Home} />
                    <Route exact path="/hotels" component={HotelList} />
                    <Route exact path="/hotels/:id" component={HotelDetail} />
                    <PrivateRoute exact path="/reservation" component={ReservationForm} />
                    <PrivateRoute exact path="/reservations" component={ReservationList} />
            </UserProvider>
        </Router>
    );
};

export default AppRouter;
