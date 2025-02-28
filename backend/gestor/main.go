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
	"bufio"
	"fmt"
	"os"
	"strings"

	// administracion de discos:
	adminDisk "gestor/comandos/adminDiscos"
)

func main() {
	fmt.Println("Que desea realizar? esperando...")

	reader := bufio.NewScanner(os.Stdin)

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

}

func analizar(entrada string) {
	/*
		se debe se separar la linea por parametros ( - )

		comadno
			-parametro1 ( obligatorio/opcional )
			-parametro2 ( obligatorio/opcional )
			-parametro3 ( obligatorio/opcional )
			- ...
	*/

	parametros := strings.Split(entrada, " -") // lista de parametros seperados por  "-"

	if strings.ToLower(parametros[0]) == "mkdisk" {
		if len(parametros) > 1 {
			// ejecutar parametros
			fmt.Println("\n --- mkdisk ---- ")
			adminDisk.Mkdisk(parametros)
		} else {
			// retornar un error
			fmt.Println(" \n --> MKDISK, ERROR: falta de parametros obligatorios")
		}
	}
}
