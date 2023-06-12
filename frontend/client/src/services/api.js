import axios from 'axios';

const BASE_URL = 'http://localhost:3000'; // Cambia la URL de acuerdo a tu configuración del backend

const api = axios.create({
    baseURL: BASE_URL,
});

export const login = async (email, password) => {
    try {
        const response = await api.post('/login', { email, password });
        return response.data;
    } catch (error) {
        throw new Error('Failed to login');
    }
};

export const getHotels = async () => {
    try {
        const response = await api.get('/hotels');
        return response.data;
    } catch (error) {
        throw new Error('Failed to fetch hotels');
    }
};

export const getHotelById = async (id) => {
    try {
        const response = await api.get(`/hotels/${id}`);
        return response.data;
    } catch (error) {
        throw new Error('Failed to fetch hotel');
    }
};

export const makeReservation = async (reservationData) => {
    try {
        const response = await api.post('/reservations', reservationData);
        return response.data;
    } catch (error) {
        throw new Error('Failed to make reservation');
    }
};

export const getReservations = async () => {
    try {
        const response = await api.get('/reservations');
        return response.data;
    } catch (error) {
        throw new Error('Failed to fetch reservations');
    }
};

export const getAdminReservations = async (hotelId, date) => {
    try {
        const response = await api.get(`/reservations?hotelId=${hotelId}&date=${date}`);
        return response.data;
    } catch (error) {
        throw new Error('Failed to fetch admin reservations');
    }
};
