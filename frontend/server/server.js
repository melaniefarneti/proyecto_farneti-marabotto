import express from 'express';

const app = express();

const router = express.Router();

const fetch = (...args) => 
    import('node-fetch').then(({default: fetch}) => fetch(...args));

const backendBaseURL = "http://localhost:8080";

app.get("/api/HotelList.js/:id", async function(req, res) {
    res.set('Access-Control-Allow-Origin', '*');
    const url = `${backendBaseURL}/hotels/${req.params.id}`;
    const options = {method: 'GET'};
    try {
        let response = await fetch(url, options);
        let status = response.status;
        response = await response.json();
        res.status(status).json(response);
    } catch (err) {
        console.log(err);
        res.status(500).json({msg: 'Internal Server Error'})
    }
})


// port 5000 cannot be used in some versions of macOS
app.listen(5001, () => {
    console.log("Server started on port 5001")
});
