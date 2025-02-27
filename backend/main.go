package main

import (
    "bufio"
    "fmt"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "net/http"
    "strings"
)

// Entrada representa la estructura de los datos enviados desde el frontend
type Entrada struct {
    Text string `json:"text"`
}

// StatusResponse representa la respuesta del servidor
type StatusResponse struct {
    Message string `json:"message"`
    Type    string `json:"type"`
}

func main() {
    // Crear una nueva instancia de Gin
    r := gin.Default()

    // Configurar CORS para Gin
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // Origen de tu aplicación React
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    // Configurar rutas
    r.POST("/analizar", getCadenaAnalizar)
    r.GET("/analizar", healthCheck)

     // Iniciar el servidor en el puerto 8080
     fmt.Println("Servidor ejecutándose en http://localhost:8080")
     r.Run(":8080")
}

// **GET /analizar** - Endpoint para verificar que el servidor esté corriendo
func healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, StatusResponse{
        Message: "El servidor está funcionando correctamente",
        Type:    "success",
    })
    fmt.Println("Servidor funcionando")
}

// **POST /analizar** - Recibe y analiza un texto enviado desde el frontend
func getCadenaAnalizar(c *gin.Context) {
    var entrada Entrada

    // Intentar decodificar el JSON enviado en la solicitud
    if err := c.ShouldBindJSON(&entrada); err != nil {
        c.JSON(http.StatusBadRequest, StatusResponse{
            Message: "Error al decodificar JSON",
            Type:    "unsuccess",
        })
        return
    }

    // Analizar cada línea del texto recibido
    lector := bufio.NewScanner(strings.NewReader(entrada.Text))
    for lector.Scan() {
        linea := lector.Text()
        analizar(linea)
    }

    fmt.Println("Cadena recibida (desde el metodo):", entrada.Text)

    // Responder con el texto procesado
    c.JSON(http.StatusOK, gin.H{
        "message":       "Recibido correctamente de la TextArea",
        "type":          "success",
        "processedText": "comando " + entrada.Text + " analizado desde el back",
    })
}

// **Función para analizar el texto recibido**
func analizar(linea string) {
    fmt.Println("Analizando línea:", linea)
    // Aquí podrías agregar cualquier lógica para procesar el texto
}