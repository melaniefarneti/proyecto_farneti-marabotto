import React, { useState, useEffect } from 'react';
import { getHotels } from '../services/api';
import Navbar from './Navbar';
import Header from './Header';
import Featured from './featured/Featured';
import PropertyList from './PropertyList';
import FeaturedProperties from './featuredProperties/FeaturedProperties';
import MailList from './MailList';
import Footer from './Footer';

const Home = () => {
    const [hotels, setHotels] = useState([]);

    useEffect(() => {
        const fetchHotels = async () => {
            try {
                const hotelsData = await getHotels();
                setHotels(hotelsData);
            } catch (error) {
                console.error(error);
            }
        };

        fetchHotels();
    }, []);

    return (
        <div className="home">
            <Navbar />
            <Header/>
            <div className="homeContainer">
                <h1 className="homeTitle">Homes guests love</h1>
                <FeaturedProperties />
                <MailList />
                <Footer />
            </div>
        </div>
    );
};

export default Home;
