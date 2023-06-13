import React from 'react';
import './Booking.css';

function Booking() {
    return (
        <div className="booking-form">
            <h3>Booking Your Hotel</h3>
            <form action="#">
                <div className="check-date">
                    <label htmlFor="date-in">Check In:</label>
                    <input type="text" className="date-input hasDatepicker" id="date-in" />
                    <i className="icon_calendar" />
                </div>
                <div className="check-date">
                    <label htmlFor="date-out">Check Out:</label>
                    <input type="text" className="date-input hasDatepicker" id="date-out" />
                    <i className="icon_calendar" />
                </div>
                <div className="select-option">
                    <label htmlFor="guest">Guests:</label>
                    <select id="guest">
                        <option value="2">2 Adults</option>
                        <option value="3">3 Adults</option>
                    </select>
                </div>
                <div className="select-option">
                    <label htmlFor="room">Room:</label>
                    <select id="room">
                        <option value="1">1 Room</option>
                        <option value="2">2 Rooms</option>
                    </select>
                </div>
                <button type="submit">Check Availability</button>
            </form>
        </div>
    );
}

export default Booking;
