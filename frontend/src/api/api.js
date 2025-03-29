import axios from "axios";

// Creamos una instancia de axios con la URL base del backend
const api = axios.create({
    baseURL: 'http://localhost:8080/api'
});

// Función para enviar el texto al backend para su análisis
export const analizarTexto = async (texto) => {
    console.log("Enviando al backend:", texto);
    try {
        const response = await api.post('/analizar', { text: texto });
        console.log("Respuesta recibida:", response);
        return response;
    } catch (error) {
        console.error("Error en la llamada API:", error);
        throw error;
    }
};

// Función para verificar que el servidor esté funcionando
export const verificarServidor = async () => {
    try {
        const response = await api.get('/status');
        console.log("Estado del servidor:", response.data);
        return response;
    } catch (error) {
        console.error("Error al verificar el servidor:", error);
        throw error;
    }
};