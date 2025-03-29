import { useState, useEffect } from "react";

// Muestra los resultados de los comandos ejecutados.
export function OutputConsole({ output }) {
  const [displayOutput, setDisplayOutput] = useState("");

  // Efecto para actualizar el output cuando cambia
  useEffect(() => {
    // Si el output cambia, actualizamos el estado local
    setDisplayOutput(output);
    
    // Para debug
    console.log("Output cambiado:", output);
  }, [output]);

  return (
    <div className="output-console">
      {/* Usar un pre para mantener el formato del texto */}
      <pre style={{ 
        whiteSpace: "pre-wrap", 
        wordWrap: "break-word",
        minHeight: "200px",
        backgroundColor: "#f5f5f5",
        padding: "10px",
        border: "1px solid #ddd",
        borderRadius: "4px",
        overflowY: "auto"
      }}>
        {displayOutput || "Esperando ejecución de comandos..."}
      </pre>
    </div>
  );
}

export default OutputConsole;