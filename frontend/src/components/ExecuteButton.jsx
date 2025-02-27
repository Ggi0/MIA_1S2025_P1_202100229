
// Un botón separado que ejecuta toda la lógica de comandos cuando se presiona.

import { analizarTexto } from "../api/api";

/**
 * Componente para el botón de ejecución
 * @param {string} fileContent - El contenido del área de texto que se enviará al backend
 * @param {function} setOutput - Función para establecer la salida en el componente OutputConsole
 */
export function ExecuteButton({ fileContent, setOutput }) {
    /**
     * Maneja el evento de clic en el botón Ejecutar
     * Envía el texto al backend y actualiza la salida con la respuesta
     */
    const handleExecute = async () => {
        if (!fileContent || fileContent.trim() === "") {
            setOutput("Error: No hay texto para analizar");
            return;
        }

        try {
            // Enviar el texto al backend para su análisis
            const respuesta = await analizarTexto(fileContent);
            console.log("Respuesta del servidor:", respuesta.data);
            
            // Mostrar el texto procesado que viene del backend
            setOutput(respuesta.data.processedText);
        } catch (error) {
            setOutput(`Error: ${error.message}`);
            console.error("Error al enviar datos al servidor:", error);
        }
    };

    return <button onClick={handleExecute}>Ejecutar</button>;
}

export default ExecuteButton;