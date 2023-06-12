export const formatDate = (dateString) => {
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    const date = new Date(dateString);
    return date.toLocaleDateString(undefined, options);
};

export const isDateAvailable = (date, reservations) => {
    const formattedDate = new Date(date).toISOString().split('T')[0];
    return !reservations.find((reservation) => reservation.date === formattedDate);
};

export const getAvailableDates = (startDate, endDate, reservations) => {
    const availableDates = [];
    const currentDate = new Date(startDate);

    while (currentDate <= new Date(endDate)) {
        if (isDateAvailable(currentDate, reservations)) {
            availableDates.push(currentDate.toISOString().split('T')[0]);
        }
        currentDate.setDate(currentDate.getDate() + 1);
    }

    return availableDates;
};
