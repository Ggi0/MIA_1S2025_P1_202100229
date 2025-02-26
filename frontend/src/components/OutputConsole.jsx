import { useState } from "react";

// Muestra los resultados de los comandos ejecutados.

export function OutputConsole({ output }) {
  return (
    <div>
      <pre>{output}</pre>
    </div>
  );
}
export default OutputConsole;
