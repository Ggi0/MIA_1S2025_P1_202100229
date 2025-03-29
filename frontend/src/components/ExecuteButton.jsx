
// Un botón separado que ejecuta toda la lógica de comandos cuando se presiona.

import { analizarTexto } from "../api/api";
import { useState } from "react";

/**
 * Componente para el botón de ejecución
 * @param {string} fileContent - El contenido del área de texto que se enviará al backend
 * @param {function} setOutput - Función para establecer la salida en el componente OutputConsole
 */
export function ExecuteButton({ fileContent, setOutput, setError }) {
    const [isProcessing, setIsProcessing] = useState(false);

    /**
     * Maneja el evento de clic en el botón Ejecutar
     * Envía el texto al backend y actualiza la salida con la respuesta
     */
    const handleExecute = async () => {
        if (!fileContent || fileContent.trim() === "") {
            setOutput("Error: No hay texto para analizar");
            return;
        }

        setIsProcessing(true);
        setOutput("Procesando comandos...");
        setError(null);

        try {
            // Enviar el texto al backend para su análisis
            const respuesta = await analizarTexto(fileContent);
            console.log("Respuesta del servidor:", respuesta.data);
            
            // Mostrar el texto procesado que viene del backend
            if (respuesta.data.tipo === "success") {
                setOutput(respuesta.data.detalles || respuesta.data.mensaje);
                setError(null);
            } else if (respuesta.data.tipo === "warning") {
                setOutput(respuesta.data.detalles || "");
                setError({
                    mensaje: respuesta.data.mensaje,
                    tipo: "warning"
                });
            } else {
                setOutput("");
                setError({
                    mensaje: respuesta.data.mensaje,
                    tipo: "error",
                    detalles: respuesta.data.detalles
                });
            }
        } catch (error) {
            // Manejo de errores más detallado
            let errorMessage = "Error al comunicarse con el servidor";
            
            if (error.response) {
                // El servidor respondió con un código de error
                errorMessage = `Error ${error.response.status}: ${error.response.data.mensaje || "Error en la respuesta del servidor"}`;
                setError({
                    mensaje: errorMessage,
                    tipo: "error",
                    detalles: error.response.data.detalles || ""
                });
            } else if (error.request) {
                // No se recibió respuesta del servidor
                errorMessage = "No se recibió respuesta del servidor. Verifica que el backend esté ejecutándose.";
                setError({
                    mensaje: errorMessage,
                    tipo: "error"
                });
            } else {
                // Error al configurar la solicitud
                errorMessage = `Error: ${error.message}`;
                setError({
                    mensaje: errorMessage,
                    tipo: "error"
                });
            }
            
            setOutput("");
            console.error("Error al enviar datos al servidor:", error);
        } finally {
            setIsProcessing(false);
        }
    };

    return (
        <button 
            onClick={handleExecute} 
            disabled={isProcessing}
            className={isProcessing ? "button-processing" : ""}
        >
            {isProcessing ? "Procesando..." : "Ejecutar"}
        </button>
    );
}

export default ExecuteButton;