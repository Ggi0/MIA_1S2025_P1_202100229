package controllers

import (
	"Gestor/models"
	"Gestor/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AnalizarComandos procesa los comandos recibidos desde el frontend
func AnalizarComandos(c *gin.Context) {
	var entrada models.EntradaComando

	// Bind JSON Body
	if err := c.ShouldBindJSON(&entrada); err != nil {
		c.JSON(http.StatusBadRequest, models.Respuesta{
			Mensaje: "Error al procesar la solicitud",
			Tipo:    "error",
			Errores: []string{err.Error()},
		})
		return
	}

	// Validar que hay texto para analizar
	if entrada.Texto == "" {
		c.JSON(http.StatusBadRequest, models.Respuesta{
			Mensaje: "No se proporcionó ningún comando para analizar",
			Tipo:    "error",
			Errores: []string{"No se proporcionó ningún comando para analizar"},
		})
		return
	}

	// Procesar cada línea del texto como un comando separado
	lineas := services.GetLineasComando(entrada.Texto)
	var salidas []string
	var errores []string
	var comandos []string
	var todosExitosos bool = true

	for _, linea := range lineas {
		if linea != "" {
			// Analizar cada línea y capturar su salida
			resultado := services.AnalizarComando(linea)
			comandos = append(comandos, resultado.Comando)

			fmt.Printf("Resultado del comando: %s\n", resultado.Comando)
			fmt.Printf("  Éxito: %v\n", resultado.Exito)

			// Si el comando generó salida, la agregamos
			if resultado.Salida != "" {
				salidas = append(salidas, resultado.Salida)
			}

			// Si el comando generó errores, los agregamos
			if resultado.Errores != "" {
				errores = append(errores, resultado.Errores)
				todosExitosos = false
			}
		}
	}

	// Construir la respuesta
	salidaFinal := ""
	for _, salida := range salidas {
		salidaFinal += salida + "\n"
	}

	if salidaFinal == "" {
		salidaFinal = "Comandos procesados pero no generaron salida."
	}

	// Tipo de respuesta basado en el resultado
	tipoRespuesta := "success"
	mensajeRespuesta := "Todos los comandos se procesaron correctamente"

	if !todosExitosos {
		tipoRespuesta = "warning"
		mensajeRespuesta = "Algunos comandos tuvieron errores"
	}

	// Devolver los resultados al frontend
	c.JSON(http.StatusOK, models.Respuesta{
		Mensaje:  mensajeRespuesta,
		Tipo:     tipoRespuesta,
		Salida:   salidaFinal,
		Errores:  errores,
		Comandos: comandos,
	})
}

// GetStatus devuelve el estado del servidor
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, models.Respuesta{
		Mensaje: "Servidor funcionando correctamente",
		Tipo:    "success",
	})
}
