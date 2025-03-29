package controllers

import (
	"Gestor/models"
	"Gestor/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AnalizarComandos procesa los comandos recibidos desde el frontend
func AnalizarComandos(c *gin.Context) {
	var entrada models.EntradaComando

	// Bind JSON Body
	if err := c.ShouldBindJSON(&entrada); err != nil {
		c.JSON(http.StatusBadRequest, models.Respuesta{
			Mensaje:  "Error al procesar la solicitud",
			Tipo:     "error",
			Detalles: err.Error(),
		})
		return
	}

	// Validar que hay texto para analizar
	if entrada.Texto == "" {
		c.JSON(http.StatusBadRequest, models.Respuesta{
			Mensaje: "No se proporcionó ningún comando para analizar",
			Tipo:    "error",
		})
		return
	}

	// Procesar cada línea del texto como un comando separado
	lineas := services.GetLineasComando(entrada.Texto)
	var resultados strings.Builder
	var errores []string

	for _, linea := range lineas {
		if linea != "" {
			// Analizar cada línea y capturar su salida
			resultado := services.AnalizarComando(linea)
			fmt.Println("Resultado del comando:", resultado) // Depuración

			// Verificar si el resultado contiene algún mensaje de error
			if strings.Contains(strings.ToLower(resultado), "error") {
				errores = append(errores, fmt.Sprintf("Error en comando: %s\n%s", linea, resultado))
			}

			resultados.WriteString(resultado + "\n")
		}
	}

	// Si hay errores, los incluimos en la respuesta
	if len(errores) > 0 {
		errorDetalles := strings.Join(errores, "\n\n")

		c.JSON(http.StatusOK, models.Respuesta{
			Mensaje:  "Algunos comandos tuvieron errores",
			Tipo:     "warning",
			Detalles: errorDetalles,
		})
		return
	}

	// Si no hay errores, devuelve los resultados
	salidaFinal := resultados.String()
	if salidaFinal == "" {
		salidaFinal = "Comando procesado pero no generó salida."
	}

	// Devolver los resultados al frontend
	c.JSON(http.StatusOK, models.Respuesta{
		Mensaje:  "Comandos procesados correctamente",
		Tipo:     "success",
		Detalles: salidaFinal,
	})
}

// GetStatus devuelve el estado del servidor
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, models.Respuesta{
		Mensaje: "Servidor funcionando correctamente",
		Tipo:    "success",
	})
}
