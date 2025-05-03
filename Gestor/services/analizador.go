package services

import (
	// "bufio" // para capturar entrada
	"bufio"
	"fmt"
	"strings"
	"unicode" // para caracteres

	// administracion de discos:
	admindiscos "Gestor/Comandos/AdminDiscos"
	filesFolders "Gestor/Comandos/AdminFiles"
	users "Gestor/Comandos/AdminUserGroup"
	rep "Gestor/Comandos/Rep"
	fileSystem "Gestor/Comandos/adminSistemaArchivos"

	"Gestor/models"
)

// AnalizarComando procesa un comando y devuelve el resultado
func AnalizarComando(entrada string) models.ResultadoComando {
	resultado := analizarEntrada(entrada)
	return resultado
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
func analizarEntrada(entrada string) models.ResultadoComando {
	var parametros []string    // Lista donde se almacenarán los parámetros
	var buffer strings.Builder // Buffer para construir cada parámetro individualmente
	enComillas := false

	// Preparamos el resultado
	resultado := models.ResultadoComando{
		Comando: entrada,
		Exito:   true, // Por defecto asumimos éxito
	}

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

	var salidaNormal strings.Builder
	var salidaError strings.Builder

	// Si no hay parámetros, retornamos error
	if len(parametros) == 0 {
		errorMsg := "ERROR: No se proporcionó ningún comando"
		salidaError.WriteString(errorMsg + "\n")
		resultado.Errores = salidaError.String()
		resultado.Exito = false
		return resultado
	}

	switch strings.ToLower(parametros[0]) {
	case "mkdisk":
		salidaNormal.WriteString("\n =================================== mkdisk =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaMkdisk := admindiscos.Mkdisk(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaMkdisk), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaMkdisk))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaMkdisk))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MK DISK ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "rmdisk":
		salidaNormal.WriteString("\n =================================== rmdisk =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaRmdisk := admindiscos.Rmdisk(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaRmdisk), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaRmdisk))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaRmdisk))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MK DISK ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "fdisk":
		salidaNormal.WriteString("\n =================================== fdisk =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaFdis := admindiscos.Fdisk(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaFdis), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaFdis))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaFdis))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MK DISK ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mount":
		salidaNormal.WriteString("\n =================================== mount =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaMount := admindiscos.Mount(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaMount), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaMount))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaMount))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MOUNT ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mounted":
		salidaNormal.WriteString("\n =================================== mount =================================== \n")
		if len(parametros) == 1 {
			// ejecutar parametros
			salidaMounted := admindiscos.Mounted(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaMounted), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaMounted))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaMounted))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MOUNTED ]: No permite más parametros"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mkfs":
		salidaNormal.WriteString("\n =================================== mkfs =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaMkfs := fileSystem.Mkfs(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaMkfs), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaMkfs))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaMkfs))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MKFS ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "cat":
		salidaNormal.WriteString("\n =================================== cat =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaMkfs := fileSystem.Cat(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaMkfs), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaMkfs))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaMkfs))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ CAT ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "login":
		salidaNormal.WriteString("\n =================================== login =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaLogin := users.Login(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaLogin), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaLogin))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaLogin))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ LOGIN ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "logout":
		salidaNormal.WriteString("\n =================================== logout =================================== \n")
		if len(parametros) == 1 {
			// ejecutar parametros
			salida := users.Logout(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ LOGOUT ]: No permite parametros adicionales"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mkgrp":
		salidaNormal.WriteString("\n =================================== mkgrp =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := users.Mkgrp(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MKGRP ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "rmgrp":
		salidaNormal.WriteString("\n =================================== rmgrp =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := users.Rmgrp(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MRGRP ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mkusr":
		salidaNormal.WriteString("\n =================================== mkusr =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := users.Mkusr(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MRGRP ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "rmusr":
		salidaNormal.WriteString("\n =================================== rmusr =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := users.Rmusr(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MRGRP ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "chgrp":
		salidaNormal.WriteString("\n =================================== chgrp =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := users.Chgrp(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MRGRP ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mkdir":
		salidaNormal.WriteString("\n =================================== mkdir =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := filesFolders.Mkdir(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MKDIR ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "mkfile":
		salidaNormal.WriteString("\n =================================== mkfile =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salida := filesFolders.Mkfile(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salida), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salida))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salida))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ MKFILE ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	case "rep":
		salidaNormal.WriteString("\n =================================== Rep =================================== \n")
		if len(parametros) > 1 {
			// ejecutar parametros
			salidaRep := rep.Reportes(parametros)

			// Detectar si hay errores en la salida
			if strings.Contains(strings.ToLower(salidaRep), "error") {
				salidaError.WriteString(fmt.Sprintf("Error en comando: %s\n%s", entrada, salidaRep))
				resultado.Exito = false
			} else {
				salidaNormal.WriteString(fmt.Sprintf("\tComando Ejecutado: %s\n%s", entrada, salidaRep))
			}
		} else {
			// retornar un error
			errorMsg := "\t ---> ERROR [ REP ]: falta de parametros obligatorios"
			salidaError.WriteString(errorMsg + "\n")
			resultado.Exito = false
		}
		salidaNormal.WriteString("\n ============================================================================== \n\n")

	default:
		errorMsg := fmt.Sprintf("\t ---> ERROR [ ]: comando no reconocido %s", strings.ToLower(parametros[0]))
		salidaError.WriteString(errorMsg + "\n")
		resultado.Exito = false
	}

	// Establecer la salida normal y los errores
	resultado.Salida = salidaNormal.String()
	resultado.Errores = salidaError.String()

	fmt.Printf("Comando: %s, Éxito: %v\n", entrada, resultado.Exito)
	if resultado.Salida != "" {
		fmt.Printf("Salida:\n%s\n", resultado.Salida)
	}

	if resultado.Errores != "" {
		fmt.Printf("Errores:\n%s\n", resultado.Errores)
	}

	return resultado
}
