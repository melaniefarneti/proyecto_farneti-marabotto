import "./Hotel.css";
import Navbar from "./Navbar";
import Header from "./Header";
import MailList from "./MailList";
import Footer from "./Footer";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
    faCircleArrowLeft,
    faCircleArrowRight,
    faCircleXmark,
    faLocationDot,
} from "@fortawesome/free-solid-svg-icons";
import { useState, useEffect } from "react";

const Hotel = () => {
    const [slideNumber, setSlideNumber] = useState(0);
    const [open, setOpen] = useState(false);
    const [hotels, setHotels] = useState([]);

    useEffect(() => {
        fetch("/hotels/gethotels")
            .then(response => response.json())
            .then(data => setHotels(data))
            .catch(error => console.log(error));
    }, []);

    const handleOpen = (i) => {
        setSlideNumber(i);
        setOpen(true);
    };

    const handleMove = (direction) => {
        let newSlideNumber;

        if (direction === "l") {
            newSlideNumber = slideNumber === 0 ? hotels.length - 1 : slideNumber - 1;
        } else {
            newSlideNumber = slideNumber === hotels.length - 1 ? 0 : slideNumber + 1;
        }

        setSlideNumber(newSlideNumber);
    };

    return (
        <div>
            <Navbar />
            <Header type="list" />
            <div className="hotelContainer">
                {open && (
                    <div className="slider">
                        <FontAwesomeIcon
                            icon={faCircleXmark}
                            className="close"
                            onClick={() => setOpen(false)}
                        />
                        <FontAwesomeIcon
                            icon={faCircleArrowLeft}
                            className="arrow"
                            onClick={() => handleMove("l")}
                        />
                        <div className="sliderWrapper">
                            <img src={hotels[slideNumber]?.photos[0]?.src} alt="" className="sliderImg" />
                        </div>
                        <FontAwesomeIcon
                            icon={faCircleArrowRight}
                            className="arrow"
                            onClick={() => handleMove("r")}
                        />
                    </div>
                )}
                {hotels.map((hotel, i) => (
                    <div className="hotelWrapper" key={i}>
                        <button className="bookNow">Reserve or Book Now!</button>
                        <h1 className="hotelTitle">{hotel.name}</h1>
                        <div className="hotelAddress">
                            <FontAwesomeIcon icon={faLocationDot} />
                            <span>{hotel.address}</span>
                        </div>
                        <span className="hotelDistance">
                            Excellent location - {hotel.distance} from center
                        </span>
                        <span className="hotelPriceHighlight">{hotel.priceHighlight}</span>
                        <div className="hotelImages">
                            {hotel.rooms.map((room, j) => (
                                <div className="hotelImgWrapper" key={j}>
                                    <img
                                        onClick={() => handleOpen(j)}
                                        src={room.photo}
                                        alt=""
                                        className="hotelImg"
                                    />
                                </div>
                            ))}
                        </div>
                        <div className="hotelDetails">
                            <div className="hotelDetailsTexts">
                                <h1 className="hotelTitle">{hotel.name}</h1>
                                <p className="hotelDesc">{hotel.description}</p>
                            </div>
                            <div className="hotelDetailsPrice">
                                <h1>Perfect for a {hotel.stayDuration}-night stay!</h1>
                                <span>
                                    Located in the real heart of {hotel.location}, this property
                                    has an excellent location score of {hotel.locationScore}!
                                </span>
                                <h2>
                                    <b>${hotel.price}</b> ({hotel.stayDuration} nights)
                                </h2>
                                <button>Book Now!</button>
                            </div>
                        </div>
                    </div>
                ))}
                <MailList />
                <Footer />
            </div>
        </div>
    );
};

export default Hotel;
