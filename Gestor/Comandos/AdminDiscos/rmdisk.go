package Comandos

import (
	"fmt"
	"os"
	"strings"
)

/*
elimina un archivo que representa a un disco duro.

	rmdisk

		-path (obligatorio) -Este parámetro será la ruta en el que se eliminará el archivo
*/
func Rmdisk(parametros []string) error {
	// Buffer para capturar toda la salida
	var output strings.Builder

	output.WriteString("\t-----> [ RM DISK ] <-----\n")
	fmt.Println("\t-----> [ RM DISK ] <-----")

	pathInit := false
	var path string

	// Recorriendo los paramtros
	for _, parametro := range parametros[1:] {
		fmt.Println(" -> Parametro: ", parametro)
		output.WriteString(fmt.Sprintf(" -> Parametro: %s\n", parametro))

		// token Parametro (parametro, valor)
		tknParam := strings.Split(parametro, "=")

		// si el token parametro no tiene su identificador y valor es un error
		if len(tknParam) != 2 {
			fmt.Println("\t ---> ERROR [ RM DISK ]: Valor desconocido del parametro, más de 2 valores para: ", tknParam[0])
			output.WriteString(fmt.Sprintf("\t ---> ERROR [ RM DISK ]: Valor desconocido del parametro, más de 2 valores para: %s\n", tknParam[0]))
			return fmt.Errorf("Valor desconocido del parametro, más de 2 valores para: %s", tknParam[0])
		}

		// id(parametro) - valor
		switch strings.ToLower(tknParam[0]) {
		case "path":
			pathInit = true
			path = tknParam[1]
			path = strings.Trim(path, `"`) // Elimina comillas si están presentes

		default:
			fmt.Println("\t ---> ERROR [ RM DISK ]: parametro desconocido: ", tknParam[0])
			output.WriteString(fmt.Sprintf("\t ---> ERROR [ RM DISK ]: parametro desconocido: %s\n", tknParam[0]))
			return fmt.Errorf("parametro desconocido: %s", tknParam[0])
		}
	}

	if pathInit {
		// Validar que el archivo exista
		_, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("\t ---> ERROR [ RM DISK ]: El disco no existe: ", path)
				output.WriteString(fmt.Sprintf("\t ---> ERROR [ RM DISK ]: El disco no existe: %s\n", path))
				return fmt.Errorf("El disco no existe: %s", path)
			}
			fmt.Println("\t ---> ERROR [ RM DISK ]: Error al verificar el disco: ", err)
			output.WriteString(fmt.Sprintf("\t ---> ERROR [ RM DISK ]: Error al verificar el disco: %v\n", err))
			return err
		}

		// Eliminar el archivo
		err = os.Remove(path)
		if err != nil {
			fmt.Println("\t ---> ERROR [ RM DISK ]: Error al eliminar el disco: ", err)
			output.WriteString(fmt.Sprintf("\t ---> ERROR [ RM DISK ]: Error al eliminar el disco: %v\n", err))
			return err
		}

		fmt.Println("\t ---> [ RM DISK ]: Disco eliminado correctamente: ", path)
		output.WriteString(fmt.Sprintf("\t ---> [ RM DISK ]: Disco eliminado correctamente: %s\n", path))
	} else {
		fmt.Println("\t ---> ERROR [ RM DISK ]: falta el parámetro obligatorio: path")
		output.WriteString("\t ---> ERROR [ RM DISK ]: falta el parámetro obligatorio: path\n")
		return fmt.Errorf("falta el parámetro obligatorio: path")
	}

	return nil
}
