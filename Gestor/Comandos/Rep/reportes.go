package Rep

import (
	"Gestor/Acciones"
	"Gestor/Estructuras"
	"Gestor/utils"
	"fmt"
	"path/filepath"
	"strings"
)

func Reportes(parametros []string) string {
	// 1) validar salidas

	logger := utils.NewLogger("rep")
	// Encabezado
	logger.LogInfo("[ REP ]")

	// 2) validar parametros

	var name string //obligatorio Nombre del tipo de reporte a generar
	var path string //obligatorio Nombre que tendrá el reporte
	var id string   //obligatorio sera el del disco o el de la particion
	//var ruta string //opcional para file y ls
	paramC := true //valida que todos los parametros sean correctos

	for _, parametro := range parametros[1:] {
		//quito los espacios en blano despues de cada parametro
		tmp2 := strings.TrimRight(parametro, " ")
		//divido cada parametro entre nombre del parametro y su valor # -size=25 -> -size, 25
		tmp := strings.Split(tmp2, "=")

		//Si falta el valor del parametro actual lo reconoce como error e interrumpe el proceso
		if len(tmp) != 2 {
			fmt.Println("REP Error: Valor desconocido del parametro ", tmp[0])
			paramC = false
			break //para finalizar el ciclo for con el error y no ejecutar lo que haga falta
		}

		if strings.ToLower(tmp[0]) == "name" {
			name = strings.ToLower(tmp[1])
		} else if strings.ToLower(tmp[0]) == "path" {
			// Eliminar comillas
			name = strings.ReplaceAll(tmp[1], "\"", "")
			path = name
		} else if strings.ToLower(tmp[0]) == "id" {
			id = strings.ToUpper(tmp[1]) //Mayusculas para tratarlo como case insensitive
		} else if strings.ToLower(tmp[0]) == "ruta" {
			//ruta = strings.ToLower(tmp[1])
		} else {
			fmt.Println("REP Error: Parametro desconocido: ", tmp[0])
			paramC = false
			break //por si en el camino reconoce algo invalido de una vez se sale
		}
	}

	// 3) validar logica comando

	if paramC {
		if name != "" && id != "" && path != "" {
			switch name {
			case "mbr":
				fmt.Println("reporte mbr")
				mbr(path, id)
			case "disk":
				fmt.Println("reporte disk")
			default:
				fmt.Println("REP Error: Reporte ", name, " desconocido")
			}
		} else {
			fmt.Println("REP Error: Faltan parametros")
		}
	}

	// 4) validar salida
	if logger.HasErrors() {
		return logger.GetErrors()
	}
	return logger.GetOutput()

}

func mbr(path string, id string) {
	var pathDico string
	existe := false

	//BUsca en struck de particiones montadas el id ingresado
	for _, montado := range Estructuras.Montadas {
		if montado.Id == id {
			pathDico = montado.PathM
			existe = true
			break
		}
	}

	//if true { //para probar los reporte hayan o no particiones montadas
	if existe {
		//Reporte
		tmp := strings.Split(path, "/") // /dir1/dir2/reporte
		nombreReporte := strings.Split(tmp[len(tmp)-1], ".")[0]

		//Disco a reportar
		tmp = strings.Split(pathDico, "/")
		disco := strings.Split(tmp[len(tmp)-1], ".")[0]

		file, err := Acciones.OpenFile(pathDico)
		if err != nil {
			return
		}

		var mbr Estructuras.MBR
		// Read object from bin file
		if err := Acciones.ReadObject(file, &mbr, 0); err != nil {
			return
		}

		// Close bin file
		defer file.Close()

		//reporte graphviz (cad es el contenido del reporte)
		//mbr
		cad := "digraph { \nnode [ shape=none ] \nTablaReportNodo [ label = < <table border=\"1\"> \n"
		cad += " <tr>\n  <td bgcolor='SlateBlue' COLSPAN=\"2\"> Reporte MBR </td> \n </tr> \n"
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='Azure'> mbr_tamano </td> \n  <td bgcolor='Azure'> %d </td> \n </tr> \n", mbr.Mbr_tamanio)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='#AFA1D1'> mbr_fecha_creacion </td> \n  <td bgcolor='#AFA1D1'> %s </td> \n </tr> \n", string(mbr.Mbr_creation_date[:]))
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='Azure'> mbr_disk_signature </td> \n  <td bgcolor='Azure'> %d </td> \n </tr>  \n", mbr.Mbr_disk_signature)
		cad += Estructuras.RepGraphviz(mbr, file)
		cad += "</table> > ]\n}"

		//reporte requerido
		carpeta := filepath.Dir(path)
		rutaReporte := "." + carpeta + "/" + nombreReporte + ".dot"

		Acciones.RepGraphizMBR(rutaReporte, cad, nombreReporte)
		fmt.Println(" Reporte MBR del disco " + disco + " creado exitosamente")
	} else {
		fmt.Println("REP Error: Id no existe")
	}
}
