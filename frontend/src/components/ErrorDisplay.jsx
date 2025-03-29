import React from 'react';

/**
 * Componente para mostrar errores
 * @param {object} error - El objeto de error a mostrar
 */
export function ErrorDisplay({ error }) {
  if (!error) return null;

  // Clases CSS para cada tipo de error
  const errorClasses = {
    error: "error-container bg-red-100 border border-red-400 text-red-700 p-4 mt-4 mb-4 rounded",
    warning: "warning-container bg-yellow-100 border border-yellow-400 text-yellow-700 p-4 mt-4 mb-4 rounded",
    info: "info-container bg-blue-100 border border-blue-400 text-blue-700 p-4 mt-4 mb-4 rounded"
  };

  // Seleccionar la clase CSS apropiada para el tipo de error
  const containerClass = errorClasses[error.tipo] || errorClasses.error;

  return (
    <div className={containerClass}>
      <h4 className="font-bold">{error.tipo === "warning" ? "Advertencia" : "Error"}</h4>
      <p>{error.mensaje}</p>
      {error.detalles && (
        <div className="mt-2">
          <details>
            <summary className="cursor-pointer font-semibold">Ver detalles</summary>
            <pre className="whitespace-pre-wrap mt-2 p-2 bg-white bg-opacity-50 rounded">
              {error.detalles}
            </pre>
          </details>
        </div>
      )}
    </div>
  );
}

export default ErrorDisplay;