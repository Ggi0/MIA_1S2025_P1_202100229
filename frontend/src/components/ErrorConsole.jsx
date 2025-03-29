import React from 'react';

// Componente específico para mostrar errores.
export function ErrorConsole({ errors }) {
  // Si no hay errores, no se muestra nada
  if (!errors || errors.length === 0) {
    return null;
  }

  return (
    <div className="error-console">
      <pre>
        {errors.map((error, index) => (
          <div key={index} className="error-item">
            {error}
            {index < errors.length - 1 && <hr />}
          </div>
        ))}
      </pre>
    </div>
  );
}

export default ErrorConsole;
