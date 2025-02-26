import Header from "../components/Header";
import { useState } from "react";
import InputConsole from "../components/InputConsole";
import OutputConsole from "../components/OutputConsole";
import FileUpload from "../components/FileUpload";
import ExecuteButton from "../components/ExecuteButton";
import ClearButton from "../components/ClearButton";
import Footer from "../components/Footer";

export function Principal() {
  // Esto TIENE que estar ANTES de cualquier uso de setFileContent
  const [fileContent, setFileContent] = useState("");
  
  // También es importante definir onCommand para InputConsole
  const handleCommand = (cmd) => {
    console.log("Comando ejecutado:", cmd);
    // Aquí irá la lógica para procesar el comando
  };
  
  return (
    <div>
      
      <Header />
      
      <div className="button-container">
        <FileUpload onFileUpload={setFileContent} />
        <ExecuteButton />
        <ClearButton />
      </div>
      
      <main>
        
        <section>
          <h3>Entrada:</h3>
          <InputConsole 
            fileContent={fileContent} 
            setFileContent={setFileContent} 
            onCommand={handleCommand} 
          />
        </section>

        <section>
          <h3>Salida:</h3>
          <OutputConsole />
        </section>

      </main>
      <Footer />
    </div>
  );
}

export default Principal;