import { analizarTexto } from "../api/api";
import { useState } from "react";

/**
 * Componente para el botón de ejecución
 * @param {string} fileContent - El contenido del área de texto que se enviará al backend
 * @param {function} setOutput - Función para establecer la salida en el componente OutputConsole
 * @param {function} setErrors - Función para establecer errores
 * @param {function} setComandos - Función para establecer comandos procesados
 */
export function ExecuteButton({ fileContent, setOutput, setErrors, setComandos }) {
    const [isProcessing, setIsProcessing] = useState(false);

    /**
     * Maneja el evento de clic en el botón Ejecutar
     * Envía el texto al backend y actualiza la salida con la respuesta
     */
    const handleExecute = async () => {
        if (!fileContent || fileContent.trim() === "") {
            setOutput("Error: No hay texto para analizar");
            setErrors(["No hay texto para analizar"]);
            return;
        }

        setIsProcessing(true);
        setOutput("Procesando comandos...");
        setErrors([]); // Limpiar errores anteriores
        setComandos([]); // Limpiar comandos anteriores

        try {
            // Enviar el texto al backend para su análisis
            const respuesta = await analizarTexto(fileContent);
            console.log("Respuesta del servidor:", respuesta.data);
            
            // Actualizar los estados según la respuesta
            setOutput(respuesta.data.salida || "");
            setErrors(respuesta.data.errores || []);
            setComandos(respuesta.data.comandos || []);
            
        } catch (error) {
            // Manejo de errores más detallado
            let errorMessage = "Error al comunicarse con el servidor";
            
            if (error.response) {
                // El servidor respondió con un código de error
                errorMessage = `Error ${error.response.status}: ${error.response.data.mensaje || "Error en la respuesta del servidor"}`;
                setErrors([errorMessage]);
                if (error.response.data.errores) {
                    setErrors(error.response.data.errores);
                }
            } else if (error.request) {
                // No se recibió respuesta del servidor
                errorMessage = "No se recibió respuesta del servidor. Verifica que el backend esté ejecutándose.";
                setErrors([errorMessage]);
            } else {
                // Error al configurar la solicitud
                errorMessage = `Error: ${error.message}`;
                setErrors([errorMessage]);
            }
            
            setOutput(""); // Limpiar la salida en caso de error
            console.error("Error al enviar datos al servidor:", error);
        } finally {
            setIsProcessing(false);
        }
    };

    return (
        <button 
            onClick={handleExecute} 
            disabled={isProcessing}
        >
            {isProcessing ? "Procesando..." : "Ejecutar"}
        </button>
    );
}

export default ExecuteButton;