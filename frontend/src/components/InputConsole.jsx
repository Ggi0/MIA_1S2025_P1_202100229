import { useState, useEffect } from "react";

export function InputConsole({ onCommand, fileContent, setFileContent }) {
  const [command, setCommand] = useState("");
  
  // Actualiza el comando cuando cambia fileContent desde el exterior
  useEffect(() => {
    if (fileContent) {
      setCommand(fileContent);
    }
  }, [fileContent]);

  const handleExecute = () => {
    if (command.trim()) {
      onCommand(command);
      setCommand("");
    }
  };

  const handleChange = (e) => {
    const newValue = e.target.value;
    setCommand(newValue);
    setFileContent(newValue); // Sincroniza la entrada con el contenido del archivo
  };

  return (
    <div>
      <textarea
        value={command}
        onChange={handleChange}
      />
    </div>
  );
}

export default InputConsole;