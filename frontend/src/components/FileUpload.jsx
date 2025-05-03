// subir un archivo desde su computadora.

import React from "react";
//                           la función recibe un objeto de props y extrae la propiedad onFileUpload.
export function FileUpload({ onFileUpload }) {
  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file && (file.name.endsWith(".txt") || file.name.endsWith(".mia"))) {
      const reader = new FileReader();
      reader.onload = (e) => {
        onFileUpload(e.target.result);
      };
      reader.readAsText(file);
    }
  };

  const handleButtonClick = () => {
    document.getElementById("fileInput").click();
  };

  return (
    <div>
      <input
        id="fileInput"
        type="file"
        accept=".txt,.smia"
        onChange={handleFileChange}
        style={{ display: "none" }}
      />
      <button onClick={handleButtonClick}>Elegir archivo</button>
    </div>
  );
}

export default FileUpload;
