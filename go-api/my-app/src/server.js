const express = require("express");
const cors = require("cors");

const app = express();

app.use(cors());

// Resto de las rutas y configuraciones del servidor

app.listen(8080, () => {
  console.log("Servidor backend iniciado en el puerto 8080");
});
