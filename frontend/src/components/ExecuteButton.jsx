
// Un botón separado que ejecuta toda la lógica de comandos cuando se presiona.

export function ExecuteButton({ onExecute }) {
    return <button onClick={onExecute}>Ejecutar</button>;
  }
  export default ExecuteButton;