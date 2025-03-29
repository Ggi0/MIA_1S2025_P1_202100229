package services

import (
	// "bufio" // para capturar entrada
	"bufio"
	"fmt"
	"strings"
	"unicode" // para caracteres

	// administracion de discos:
	AdminDisk "Gestor/Comandos/AdminDiscos"
)

// AnalizarComando procesa un comando y devuelve su salida como string
func AnalizarComando(entrada string) string {
	// Enfoque simplificado: procesamos directamente y devolvemos el resultado
	return analizarEntrada(entrada)
}

// GetLineasComando divide un texto en líneas individuales de comandos
func GetLineasComando(texto string) []string {
	var lineas []string
	scanner := bufio.NewScanner(strings.NewReader(texto))
	for scanner.Scan() {
		linea := scanner.Text()
		// Dividir por # para ignorar comentarios
		comandoSinComentario := strings.Split(linea, "#")[0]
		if strings.TrimSpace(comandoSinComentario) != "" {
			lineas = append(lineas, comandoSinComentario)
		}
	}
	return lineas
}

/*
Funcion que procesa una cadena de entrada para separar los parametros.
*/
func analizarEntrada(entrada string) string {
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
			buffer.Reset()
			continue
		}

		// Si NO estamos dentro de comillas y encontramos un espacio, lo ignoramos
		if !enComillas && unicode.IsSpace(char) {
			continue
		}
		// Agregamos el carácter actual al buffer
		buffer.WriteRune(char)
	}

	// Agregar el último parámetro si el buffer contiene algo
	if buffer.Len() > 0 {
		parametros = append(parametros, buffer.String())
	}

	var resultado strings.Builder

	switch strings.ToLower(parametros[0]) {
	case "mkdisk":
		if len(parametros) > 1 {
			// ejecutar parametros
			resultado.WriteString("\n\n ----------------------------------- mkdisk ----------------------------------- \n")
			// Llamar directamente a la función que devuelve un string
			salidaMkdisk := AdminDisk.Mkdisk(parametros)
			resultado.WriteString(salidaMkdisk)
			resultado.WriteString("\n ------------------------------------------------------------------------------ \n\n ")
		} else {
			// retornar un error
			resultado.WriteString("\t ---> ERROR [ MK DISK ]: falta de parametros obligatorios\n")
		}
	case "rmdisk":
		if len(parametros) > 1 {
			// ejecutar parametros
			resultado.WriteString("\n\n ----------------------------------- rmdisk ----------------------------------- \n")
			// Aquí deberíamos modificar también RmDisk para que devuelva un string
			// Por ahora capturamos mensajes de error explícitamente
			if err := AdminDisk.Rmdisk(parametros); err != nil {
				resultado.WriteString(fmt.Sprintf("ERROR: %v\n", err))
			} else {
				resultado.WriteString("Disco eliminado correctamente\n")
			}
			resultado.WriteString("\n ------------------------------------------------------------------------------ \n\n ")
		} else {
			// retornar un error
			resultado.WriteString("\t ---> ERROR [ RM DISK ]: falta de parametros obligatorios\n")
		}

	case "fdisk":
		if len(parametros) > 1 {
			// ejecutar parametros
			resultado.WriteString("\n\n ----------------------------------- fdisk ----------------------------------- \n")
			// Aquí también deberíamos modificar Fdisk para que devuelva un string
			// Por ahora capturamos mensajes de error explícitamente
			if err := AdminDisk.Fdisk(parametros); err != nil {
				resultado.WriteString(fmt.Sprintf("ERROR: %v\n", err))
			} else {
				resultado.WriteString("Partición creada correctamente\n")
			}
			resultado.WriteString("\n ------------------------------------------------------------------------------ \n\n ")
		} else {
			// retornar un error
			resultado.WriteString("\t ---> ERROR [ F DISK ]: falta de parametros obligatorios\n")
		}

	default:
		resultado.WriteString(fmt.Sprintf("\t ---> ERROR [ ]: comando no reconocido %s\n", strings.ToLower(parametros[0])))
	}

	// Este printf es solo para debugging en la consola del servidor
	fmt.Printf("Resultado del análisis: %s\n", resultado.String())

	return resultado.String()
}
