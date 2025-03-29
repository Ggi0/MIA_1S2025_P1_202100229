package Comandos

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"Gestor/Acciones"
	"Gestor/Estructuras"
)

/*
Este comando creará un archivo binario que simulará un disco (.mia)

mkdisk

	-size (obligatoria)  Tamaño del disco
	-fit  (opcional)     BF/FF/WF ajuste a utilizar
	-unit (opcional)     Kilobytes (K)/ Megabytes (M) unidades a utilizar
	-path (obligatoria)  ruta en donde se creará el archivo
*/
func Mkdisk(parametros []string) string {
	// Buffer para capturar toda la salida
	var output strings.Builder

	fmt.Println("\t-----> [ MK DISK ] <-----")
	output.WriteString("\t-----> [ MK DISK ] <-----\n")

	var size int
	fit := "F"      // valor por deferto FF
	unit := 1048576 // valor por defecto M (1024 *1024)
	var path string // para la ruta

	paramCorrectos := true // validar que todos los parametros ingresen de forma correcta
	sizeInit := false      // para saber si entro el parametro size, false cuando no esta inicializado
	pathInit := false      // para verificar la existencia del path

	// Recorriendo los paramtros
	for _, parametro := range parametros[1:] { // a partir del primero, ya que el primero es la ruta
		fmt.Println(" -> Parametro: ", parametro)
		// token Parametro (parametro, valor) --> dos valores: ["clave", "valor"]
		tknParam := strings.Split(parametro, "=")

		// si el token parametro no tiene su identificador y valor es un error
		if len(tknParam) != 2 {
			fmt.Println("\t ---> ERROR [ MK DISK ]: Valor desconocido del parametro, mas de 2 valores para: ", tknParam[0])
			paramCorrectos = false
			break // sale de analizar el parametro y no lo ejecuta
		}

		// id(parametro) - valor
		switch strings.ToLower(tknParam[0]) {
		case "size":
			sizeInit = true                       // el valor si viene dentro de las especificaciones
			var err error                         // variable para el error posible
			size, err = strconv.Atoi(tknParam[1]) // string a int
			if err != nil {
				log.Fatal(err)
			}

			if err != nil {
				fmt.Println("\t ---> ERROR [ MK DISK ]: size debe ser un valor numerico. Se leyo: ", tknParam[1])
				paramCorrectos = false
				break
			} else if size <= 0 {
				fmt.Println("\t ---> ERROR [ MK DISK ]: size debe ser mayor a cero. se leyo: ", tknParam[1])
				paramCorrectos = false
				break
			}

		case "fit":
			// es B/W/F porque en MBR espera estos valores
			if strings.ToLower(tknParam[1]) == "bf" { //Si el ajuste es BF (best fit)
				fit = "B"

			} else if strings.ToLower(tknParam[1]) == "wf" { //Si el ajuste es WF (worst fit)
				fit = "W"

			} else if strings.ToLower(tknParam[1]) != "ff" { //Si el ajuste es diferente a ff es distinto es un error
				fmt.Println("\t ---> ERROR [ MK DISK ]: para FIT los valores aceptados son: BF, FF o WF. ingreso: ", tknParam[1])
				paramCorrectos = false
				break
			}

		case "unit":
			//si la unidad es k
			if strings.ToLower(tknParam[1]) == "k" {
				unit = 1024 // bites
			} else if strings.ToLower(tknParam[1]) != "m" {
				fmt.Println("\t ---> ERROR [ MK DISK ]: para UNIT los valores aceptados son: K y M. ingreso: ", tknParam[1])
				paramCorrectos = false
				break
			}

		case "path":
			pathInit = true
			path = tknParam[1]
			// TODO: validar errores, por ejemplo la ruta existe?

		default:
			fmt.Println("\t ---> ERROR [ MK DISK ]: parametro desconocido: ", tknParam[0])
			paramCorrectos = false
			break
		}

	}

	// si se llego aqui todos los parametros estan correctos.
	// -------------- validación de parametros CORRECTOS --------------
	if paramCorrectos {
		// esta información necesaria para la CREACION real del Disco
		if sizeInit && pathInit { // validar los parametros obligatorios
			// tamanio del disco
			//fmt.Println("validando:  ", size, "*", unit)
			tamanio := size * unit
			fmt.Println("--> Tamanio del disco: ", tamanio, " Bytes.")

			// nombre del disco
			path = strings.Trim(path, `"`) // Elimina comillas si están presentes
			ruta := strings.Split(path, "/")
			nombreDisco := ruta[len(ruta)-1] // el ultimo valor de la ruta

			fmt.Println("--> Nombre del disco: ", "'", nombreDisco, "'")

			// en este punto tenemos todo lo necesario para crear el Disco
			fmt.Println("--> FIT: ", fit)

			// CREAR EL DISCO -> hacer el archivo binario que simule el disco
			err := Acciones.CrearDisco(path, nombreDisco)
			if err != nil {
				fmt.Println("\t ---> ERROR [ MK DISK ]: ", err)
			}

			// ABRIR EL DISCO -> para completar su contenido inicial (MBR)
			file, err := Acciones.OpenFile(path)

			if err != nil {
				defer file.Close()
				return output.String()
			}

			// --> aqui se valida que el tamanio no crezca

			// A traves del tamanio establecido llena de 0 hasta esa posición.
			// cantidad de valores en binario -> del tamanio del disco que necesitamos
			datos := make([]byte, tamanio)                 // llenar el disco de Ceros (0)
			newErr := Acciones.WriteObject(file, datos, 0) //--> desde la posicion 0
			if newErr != nil {
				fmt.Println("\t ---> ERROR [ MK DISK ]: ", err)
				defer file.Close()
				return output.String()
			}

			// Escribir el MBR para completar el proceso de creacion del DISCO
			file, errr := Estructuras.EscribirMBR(file, tamanio, fit)
			if errr != nil {
				defer file.Close()
				return output.String()
			}

			defer file.Close()
			fmt.Println("\n[ MK DISK ]: Proceso completado, el disco", nombreDisco, " Fue creado CORRECTAMENTE. en: ", file.Name())
			output.WriteString(fmt.Sprintf("\n[ MK DISK ]: Proceso completado, el disco %s Fue creado CORRECTAMENTE. en: %s\n", nombreDisco, file.Name()))

		}
	} else {
		fmt.Println("\t ---> ERROR [ MK DISK ]: parametros ingresados incorrectamente ")
	}

	return output.String()

}
