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
	var errores []models.ComandoError

	for _, linea := range lineas {
		if linea != "" {
			// Analizar cada línea y capturar su salida
			resultado := services.AnalizarComando(linea)
			fmt.Println("Resultado del comando:", resultado) // Depuración

			// Verificar si el resultado contiene algún mensaje de error
			if strings.Contains(strings.ToLower(resultado), "error") {
				errores = append(errores, *models.NewComandoError(
					resultado,
					"comando",
					linea,
					"Error al ejecutar el comando",
				))
			}

			resultados.WriteString(resultado + "\n")
		}
	}

	// Si hay errores, los incluimos en la respuesta
	if len(errores) > 0 {
		errorDetalles := make([]string, len(errores))
		for i, err := range errores {
			errorDetalles[i] = err.Mensaje
		}

		c.JSON(http.StatusOK, models.Respuesta{
			Mensaje:  "Comandos procesados con errores",
			Tipo:     "warning",
			Detalles: resultados.String(),
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
