import Header from "../components/Header";
import { useState } from "react"; // es un hook que permite añadir estado a componentes funcionales.
import InputConsole from "../components/InputConsole";
import OutputConsole from "../components/OutputConsole";
import FileUpload from "../components/FileUpload";
import ExecuteButton from "../components/ExecuteButton";
import ClearButton from "../components/ClearButton";
import ErrorConsole from "../components/ErrorConsole";
import Footer from "../components/Footer";

/**
 * Componente principal de la aplicación
 * Gestiona el estado y la interacción entre los componentes
 */
export function Principal() {
  // Estados para gestionar el contenido del archivo y la salida
  const [fileContent, setFileContent] = useState("");
  const [output, setOutput] = useState(""); // Estado para almacenar la salida
  const [errors, setErrors] = useState([]); // Estado para almacenar errores (array)
  const [comandos, setComandos] = useState([]); // Estado para almacenar comandos procesados

  /**
   * Maneja el comando ejecutado desde InputConsole
   * @param {string} cmd - El comando ejecutado
   */
  const handleCommand = (cmd) => {
    console.log("Comando ejecutado:", cmd);
    // Esta función ahora no se utiliza ya que el botón Ejecutar maneja la comunicación
  };

  /**
   * Limpia el contenido del área de texto y la salida
   */
  const handleClear = () => {
    setFileContent("");
    setOutput("");
    setError(null);
  };


  return (
    <div>
      <Header />
      
      <div className="button-container">
        <FileUpload onFileUpload={setFileContent} />
        <ExecuteButton 
          fileContent={fileContent} 
          setOutput={setOutput} 
          setErrors={setErrors}
          setComandos={setComandos}
        />
        <ClearButton onClear={handleClear} />
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
          <OutputConsole output={output} />
        </section>
        
        
      </main>
      {/* Sección para mostrar errores solo si hay errores */}
      {errors && errors.length > 0 && (
          <section className="error-section">
            <h3>Errores:</h3>
            <ErrorConsole errors={errors} />
          </section>
        )}
      
      <Footer />
    </div>
  );
}

export default Principal;