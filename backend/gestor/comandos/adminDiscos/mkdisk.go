package comandos

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
Este comando creará un archivo binario que simulará un disco (.mia)

mkdisk

	-size (obligatoria)  Tamaño del disco
	-fit  (opcional)     BF/FF/WF ajuste a utilizar
	-unit (opcional)     Kilobytes (K)/ Megabytes (M) unidades a utilizar
	-path (obligatoria)  ruta en donde se creará el archivo
*/
func Mkdisk(parametros []string) {
	fmt.Println("          ---> mkdisk <---")

	var size int
	fit := "F"      // valor por deferto FF
	unit := 1048576 // valor por defecto M (1024 *1024)
	var path string

	paramCorrectos := true // validar que todos los parametros ingresen de forma correcta
	sizeInit := false      // para saber si entro el parametro size, false cuando no esta inicializado
	pathInit := false      // para verificar la existencia del path

	// Recorriendo los paramtros
	for _, parametro := range parametros[1:] {
		fmt.Println("parametro: ", parametro)
		// token Parametro (parametro, valor) -> (size, 300)
		tknParam := strings.Split(parametro, "=")

		// si el token parametro no tiene su identificador y valor es un error
		if len(tknParam) != 2 {
			fmt.Println(" \n --> MKDISK, ERROR: Valor desconocido del parametro ", tknParam[0])
			paramCorrectos = false
			break // sale de analizar el parametro y no lo ejecuta
		}

		// id(parametro) - valor
		switch strings.ToLower(tknParam[0]) {
		case "size":
			sizeInit = true // el valor si viene dentro de las especificaciones
			var err error   // variable para el error posible
			size, err = strconv.Atoi(tknParam[1])
			if err != nil {
				log.Fatal(err)
			}

			if err != nil {
				fmt.Println(" \n --> MKDISK, ERROR: size debe ser un valor numerico. Se leyo: ", tknParam[1])
				paramCorrectos = false
				break
			} else if size <= 0 {
				fmt.Println("MKDISK, Error: size debe ser mayor a cero. se leyo ", tknParam[1])
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
				fmt.Println("MKDISK, ERROR, para FIT los valores aceptados son: BF, FF o WF. ingreso: ", tknParam[1])
				paramCorrectos = false
				break
			}

		case "unit":
			//si la unidad es k
			if strings.ToLower(tknParam[1]) == "k" {
				unit = 1024 // bites
			} else if strings.ToLower(tknParam[1]) != "m" {
				fmt.Println("MKDISK, ERROR, para UNIT los valores aceptados son: K, M. ingreso: ", tknParam[1])
				paramCorrectos = false
				break
			}

		case "path":
			pathInit = true
			path = tknParam[1]
			// validar errores, por ejemplo la ruta existe?
		default:
			fmt.Println(" \n --> MKDISK, ERROR: parametro desconocido ", tknParam[0])
			paramCorrectos = false
			break
		}

		// validación de parametros correcta
		if paramCorrectos {
			// esta información necesaria para la CREACION real del Disco
			if sizeInit && pathInit { // validar los parametros obligatorios
				// tamanio del disco
				fmt.Println("validando:  ", size, "*", unit)
				tamanio := size * unit
				fmt.Println("tamanio del disco: ", tamanio)

				// nombre del disco
				path = strings.Trim(path, `"`) // Elimina comillas si están presentes
				ruta := strings.Split(path, "/")
				nombreDisco := ruta[len(ruta)-1]
				fmt.Println("nombre del disco: ", "'", nombreDisco, "'")

				// en este punto tenemos todo lo necesario para crear el Disco
				fmt.Println("FIT: ", fit)
				break

			}
		} else {
			fmt.Println(" \n --> MKDISK, ERROR: parametros ingresados incorrectamente ")
		}

	}
}
