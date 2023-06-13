import React, { useState, useEffect } from 'react';
import Navbar from './Navbar';
import Header from './Header';
import FeaturedProperties from './FeaturedProperties';
import MailList from './MailList';
import Footer from './Footer';

const Home = () => {
    const [hotels, setHotels] = useState([]);

    useEffect(() => {
        const fetchHotels = async () => {
            try {
                const response = await fetch('/hotels/gethotels');
                const data = await response.json();
                setHotels(data);
            } catch (error) {
                console.error(error);
            }
        };

        fetchHotels();
    }, []);

    return (
        <div className="home">
            <Navbar />
            <Header />
            <div className="homeContainer">
                <h1 className="homeTitle">Homes guests love</h1>
                <FeaturedProperties hotels={hotels} />
                <MailList />
                <Footer />
            </div>
        </div>
    );
};

export default Home;
