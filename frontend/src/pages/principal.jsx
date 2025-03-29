import Header from "../components/Header";
import { useState } from "react"; // es un hook que permite añadir estado a componentes funcionales.
import InputConsole from "../components/InputConsole";
import OutputConsole from "../components/OutputConsole";
import FileUpload from "../components/FileUpload";
import ExecuteButton from "../components/ExecuteButton";
import ClearButton from "../components/ClearButton";
import Footer from "../components/Footer";

/**
 * Componente principal de la aplicación
 * Gestiona el estado y la interacción entre los componentes
 */
export function Principal() {
  // Estados para gestionar el contenido del archivo y la salida
  const [fileContent, setFileContent] = useState("");
  const [output, setOutput] = useState(""); // Estado para almacenar la salida
  const [error, setError] = useState(null); // Estado para manejar errores

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

  // Función para renderizar el error basado en su tipo
  const renderError = () => {
    if (!error) return null;

    // Clases CSS para cada tipo de error
    const errorClasses = {
      error:
        "bg-red-100 border border-red-400 text-red-700 p-4 mt-4 mb-4 rounded",
      warning:
        "bg-yellow-100 border border-yellow-400 text-yellow-700 p-4 mt-4 mb-4 rounded",
      info: "bg-blue-100 border border-blue-400 text-blue-700 p-4 mt-4 mb-4 rounded",
    };

    const containerClass = errorClasses[error.tipo] || errorClasses.error;

    return (
      <div className={containerClass}>
        <h4 className="font-bold">
          {error.tipo === "warning" ? "Advertencia" : "Error"}
        </h4>
        <p>{error.mensaje}</p>
        {error.detalles && (
          <div className="mt-2">
            <details>
              <summary className="cursor-pointer font-semibold">
                Ver detalles
              </summary>
              <pre className="whitespace-pre-wrap mt-2 p-2 bg-white bg-opacity-50 rounded">
                {error.detalles}
              </pre>
            </details>
          </div>
        )}
      </div>
    );
  };

  return (
    <div>
      <Header />

      <div className="button-container">
        <FileUpload onFileUpload={setFileContent} />
        <ExecuteButton
          fileContent={fileContent}
          setOutput={setOutput}
          setError={setError}
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

      {/* Sección para mostrar errores */}
      {error && (
          <section className="error-section">
            <h3>Errores:</h3>
            {renderError()}
          </section>
        )}

      <Footer />
    </div>
  );
}

export default Principal;
