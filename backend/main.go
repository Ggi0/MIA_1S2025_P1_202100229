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
	"unicode" // para caracteres

	// administracion de discos:
	AdminDisk "Gestor/Comandos/AdminDiscos"
)

func main() {
	fmt.Print("\n\nBienvenido, ¿Que desea realizar? esperando...\n\n$: ")

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

/*
Funcion que procesa una cadena de entrada para separar los parametros.
*/
func analizar(entrada string) {

	/*
		se debe se separar la linea por parametros, guión ( - )

		comando
			-parametro1 ( obligatorio/opcional )
			-parametro2 ( obligatorio/opcional )
			-parametro3 ( obligatorio/opcional )
			- ...
	*/
	var parametros []string    // Lista donde se almacenarán los parámetros
	var buffer strings.Builder // Buffer para construir cada parámetro individualmente
	enComillas := false

	// Recorremos cada caracter de la entrada
	for i, char := range entrada {
		if char == '"' {
			enComillas = !enComillas // Cambiar estado de comillas
		}

		// Si encontramos un '-', NO estamos dentro de comillas y no es el primer carácter
		if char == '-' && !enComillas && i > 0 {
			parametros = append(parametros, buffer.String()) // Guardamos lo que llevamos en buffer
			// fmt.Println("==========> ", buffer.String())
			buffer.Reset()
			continue
		}

		// // Si NO estamos dentro de comillas y encontramos un espacio, lo ignoramos
		if !enComillas && unicode.IsSpace(char) {
			continue
		}
		// Agregamos el carácter actual al buffer
		//fmt.Println("--------->", string(rune(char)))
		buffer.WriteRune(char)
	}

	// Agregar el último parámetro si el buffer contiene algo
	if buffer.Len() > 0 {
		parametros = append(parametros, buffer.String())
	}
	/*
		// Imprimir los parámetros obtenidos
		for i, param := range parametros {
			fmt.Printf("*-*-*-*-*-*->> parametros[%d] = %s\n", i, param)
		}
	*/

	// parametros := strings.Split(entrada, " -") // lista de parametros seperados por  "-"

	if strings.ToLower(parametros[0]) == "mkdisk" {
		if len(parametros) > 1 {
			// ejecutar parametros
			fmt.Println("\n =-=-=-=-=-=-= =-=-=-=-=-=-= =-=-=-=-=-=-= =-=-=-=-=-=-= ")
			AdminDisk.Mkdisk(parametros)
			fmt.Println(" =-=-=-=-=-=-= =-=-=-=-=-=-= =-=-=-=-=-=-= =-=-=-=-=-=-= ")
		} else {
			// retornar un error
			fmt.Println(" \n --> MKDISK, ERROR: falta de parametros obligatorios")
		}
	}
}
