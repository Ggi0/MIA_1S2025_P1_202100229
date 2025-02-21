import { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [mensaje, setMensaje] = useState("");
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(true);
    axios.get("http://localhost:8080/api/saludo")
      .then(res => {
        console.log("Respuesta del backend:", res.data); // Para debug
        setMensaje(res.data.mensaje);
        setLoading(false);
      })
      .catch(err => {
        console.error("Error detallado:", err);
        setError("Error al conectar con el backend");
        setLoading(false);
      });
  }, []);

  return (
    <div>
      <h1>Frontend en React</h1>
      {loading && <p>Cargando...</p>}
      {error && <p style={{color: 'red'}}>{error}</p>}
      {mensaje && <p>{mensaje}</p>}
    </div>
  );
}

export default App;