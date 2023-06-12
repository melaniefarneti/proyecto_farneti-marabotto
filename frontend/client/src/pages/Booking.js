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
                        <option value="">2 Adults</option>
                        <option value="">3 Adults</option>
                    </select>
                    <div className="nice-select" tabIndex="0">
                        <span className="current">2 Adults</span>
                        <ul className="list">
                            <li data-value="" className="option selected">
                                2 Adults
                            </li>
                            <li data-value="" className="option">
                                3 Adults
                            </li>
                        </ul>
                    </div>
                </div>
                <div className="select-option">
                    <label htmlFor="room">Room:</label>
                    <div className="nice-select" tabIndex="0">
                        <span className="current">1 Room</span>
                        <ul className="list">
                            <li data-value="" className="option selected">
                                1 Room
                            </li>
                            <li data-value="" className="option">
                                2 Room
                            </li>
                        </ul>
                    </div>
                </div>
                <button type="submit">Check Availability</button>
            </form>
        </div>
    );
}

export default Booking;


