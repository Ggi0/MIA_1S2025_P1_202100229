import axios from "axios";

// Creamos una instancia de axios con la URL base del backend
const api = axios.create({
    baseURL: 'http://localhost:8080'
});

// Función para enviar el texto al backend para su análisis
export const analizarTexto = (texto) => {
    return api.post('/analizar', { text: texto });
};

// Función para verificar que el servidor esté funcionando
export const verificarServidor = () => {
    return api.get('/analizar');
};