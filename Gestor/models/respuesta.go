package models

// Respuesta representa la estructura de respuesta que enviaremos al frontend
type Respuesta struct {
	Mensaje  string `json:"mensaje"`            // Mensaje principal
	Tipo     string `json:"tipo"`               // Tipo de mensaje: "error", "success", "info"
	Detalles string `json:"detalles,omitempty"` // Detalles adicionales (opcional)
}

// EntradaComando representa la estructura JSON que recibiremos desde el frontend
type EntradaComando struct {
	Texto string `json:"text"` // El texto del comando a analizar
}
