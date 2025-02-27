/**
 * Componente para el botón de limpiar
 * @param {function} onClear - Función para limpiar la entrada y salida
 */
export function ClearButton({ onClear }) {
  return <button onClick={onClear}>Limpiar</button>;
}

export default ClearButton;