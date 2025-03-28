package Comandos

import (
	"Gestor/Acciones"
	"fmt"
	"os"
	"strings"
)

/*
elimina un archivo que representa a un disco duro.

	rmdisk

		-path (obligatorio) -Este parámetro será la ruta en el que se eliminará el archivo
*/
func Rmdisk(parametros []string) {
	fmt.Println("\t-----> [ RM DISK ] <-----")

	var path string // ruta del Disco
	var nombreDisco string

	paramCorrectos := true // validar que todos los parametros ingresen de forma correcta
	pathInit := false      // para verificar la existencia del path

	// Recorriendo los paramtros
	for _, parametro := range parametros[1:] { // a partir del primero, ya que el primero es la ruta
		fmt.Println(" -> Parametro: ", parametro)

		// token Parametro (parametro, valor) --> dos valores: ["clave", "valor"]
		tknParam := strings.Split(parametro, "=")

		// si el token parametro no tiene su identificador y valor es un error
		if len(tknParam) != 2 {
			fmt.Println("\t ---> ERROR [ RM DISK ]: Valor desconocido del parametro, mas de 2 valores para: ", tknParam[0])
			paramCorrectos = false
			break // sale de analizar el parametro y no lo ejecuta
		}

		// ---------- VALIDANDO PARAMATROS ---------------------
		switch strings.ToLower(tknParam[0]) {
		case "path":
			path = tknParam[1]

			if path != "" {
				// ruta correcta
				path = Acciones.RutaCorrecta(path)

				// nombre del disco
				//path = strings.Trim(path, `"`) // Elimina comillas si están presentes
				ruta := strings.Split(path, "/")
				nombreDisco = ruta[len(ruta)-1] // el ultimo valor de la ruta

				pathInit = true

				_, err := os.Stat(path)
				if os.IsNotExist(err) {
					fmt.Println("\t ---> ERROR [ RM DISK ]: El disco ", nombreDisco, " no existe")
					paramCorrectos = false
					break // Terminar el bucle porque encontramos un nombre único
				}
			} else {
				fmt.Println("\t ---> ERROR [ RM DISK ]: error en ruta, esta vacia")
				paramCorrectos = false
				break
			}

		default:
			fmt.Println("\t ---> ERROR [ RM DISK ]: parametro desconocido: ", tknParam[0])
			paramCorrectos = false
			break
		}

	}

	// si se llego aqui todos los parametros estan correctos.
	// -------------- validación de parametros CORRECTOS --------------
	if paramCorrectos && pathInit { // validar los parametros obligatorios (la path)
		// ----- LOGICA PARA EL RMDISK ---------

		// Intentar eliminar el archivo
		err := os.Remove(path)
		if err != nil {
			fmt.Println("\t ---> ERROR [ RM DISK ]: No se pudo eliminar el disco:", err)
		} else {
			fmt.Println("\t ---> [ RM DISK ]: Disco", nombreDisco, "eliminado correctamente")
		}

	} else {
		fmt.Println("\t ---> ERROR [ RM DISK ]: parametros ingresados incorrectamente ")
	}

}
