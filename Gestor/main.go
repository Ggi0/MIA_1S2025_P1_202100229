package main

/*
	Giovanni Concohá - Ggi0

	Aplicación web para gestionar un sistema de archivos EXT2,
	usando React para el frontend y Go para el backend.
	Permite crear/administrar particiones, gestionar archivos/carpetas,
	manejar permisos de usuarios/grupos y generar reportes. El sistema
	expone una API REST para las operaciones del sistema de archivos.

*/

import (
	"Gestor/controllers" // Importamos el paquete de controladores que crearemos
	"log"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin" // Importamos cors para permitir peticiones desde el frontend
)

/*
	main funcional para salidas en consola:
*/
/*func main() {
	fmt.Print("\n\nBienvenido, ¿Que desea realizar? esperando...\n\n$: ")

	reader := bufio.NewScanner(os.Stdin) // variable con info desde el front

	for {

		// capturar la entrada:
		reader.Scan()

		entrada := strings.TrimRight(reader.Text(), " ") // omitir espacios vacios
		linea := strings.Split(entrada, "#")             // ignorar comentarios

		if strings.ToLower(linea[0]) != "exit" { // pasar todo a minusculas
			analizar(linea[0])

		} else {
			fmt.Println("\n ... Adios ... ")
			break
		}

	}

}*/

func main() {
	// Iniciamos el router de Gin
	router := gin.Default()

	// Habilitamos CORS para que el frontend pueda comunicarse con el backend
	router.Use(cors.Default())

	// Configuramos las rutas de nuestra API
	api := router.Group("/api")
	{
		// Ruta para analizar comandos
		api.POST("/analizar", controllers.AnalizarComandos)

		// Podemos agregar más rutas según sea necesario
		api.GET("/status", controllers.GetStatus)
	}

	// Iniciamos el servidor en el puerto 8080
	log.Println("Servidor escuchando en http://localhost:8080")
	router.Run(":8080")
}
