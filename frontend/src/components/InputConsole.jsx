import { useState } from "react";

// Es un área de entrada tipo textarea donde el usuario escribe comandos.

export function InputConsole({ onCommand }) {
  
  const [command, setCommand] = useState(""); // para manejar el estado del texto ingresado.
  
  const handleExecute = () => {
    if (command.trim()) {
      onCommand(command);
      setCommand("");
    }
  };

  return (
    <div>
      <textarea value={command} onChange={(e) => setCommand(e.target.value)} />
    </div>
  );
}
export default InputConsole;
