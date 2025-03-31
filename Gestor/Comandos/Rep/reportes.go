package Rep

import (
	"Gestor/Acciones"
	"Gestor/Estructuras"
	"Gestor/utils"
	"fmt"
	"os"
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
			case "ebr":
				fmt.Println("reporte ebr")
				ebr(path, id)
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

func ebr(path string, id string) {
	var pathDisco string
	var particionExtendida Estructuras.Partition
	encontrada := false
	existe := false

	// Busca en struct de particiones montadas el id ingresado
	for _, montado := range Estructuras.Montadas {
		if montado.Id == id {
			pathDisco = montado.PathM
			existe = true
			break
		}
	}

	if existe {
		// Reporte
		tmp := strings.Split(path, "/") // /dir1/dir2/reporte
		nombreReporte := strings.Split(tmp[len(tmp)-1], ".")[0]

		// Disco a reportar
		tmp = strings.Split(pathDisco, "/")
		disco := strings.Split(tmp[len(tmp)-1], ".")[0]

		file, err := Acciones.OpenFile(pathDisco)
		if err != nil {
			fmt.Println("REP Error: No se pudo abrir el disco")
			return
		}
		defer file.Close()

		// Leer el MBR para encontrar la partición extendida
		var mbr Estructuras.MBR
		if err := Acciones.ReadObject(file, &mbr, 0); err != nil {
			fmt.Println("REP Error: No se pudo leer el MBR")
			return
		}

		// Buscar la partición extendida en el MBR
		for i := 0; i < 4; i++ {
			tipo := string(mbr.Mbr_partitions[i].Part_type[:])
			if tipo == "E" {
				particionExtendida = mbr.Mbr_partitions[i]
				encontrada = true
				break
			}
		}

		if !encontrada {
			fmt.Println("REP Error: No se encontró partición extendida en el disco")
			return
		}

		// Inicio del reporte de EBRs y particiones lógicas
		cad := "digraph { \nnode [ shape=none ] \nTablaReportNodo [ label = < <table border=\"1\"> \n"
		cad += " <tr>\n  <td bgcolor='SlateBlue' COLSPAN=\"2\"> Reporte EBR de Particiones Lógicas </td> \n </tr> \n"
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='LightSteelBlue' COLSPAN=\"2\"> Partición Extendida: %s </td> \n </tr> \n",
			Estructuras.GetName(string(particionExtendida.Part_name[:])))
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='#AFA1D1'> Inicio Partición Extendida </td> \n  <td bgcolor='#AFA1D1'> %d </td> \n </tr> \n",
			particionExtendida.Part_start)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='Azure'> Tamaño Partición Extendida </td> \n  <td bgcolor='Azure'> %d </td> \n </tr> \n",
			particionExtendida.Part_size)

		// Añadir información de cada EBR y su partición lógica asociada
		cad += reporteEBRs(particionExtendida, file)
		cad += "</table> > ]\n}"

		// Generar reporte
		carpeta := filepath.Dir(path)
		rutaReporte := "." + carpeta + "/" + nombreReporte + ".dot"

		Acciones.RepGraphizMBR(rutaReporte, cad, nombreReporte)
		fmt.Println(" Reporte EBR del disco " + disco + " creado exitosamente")
	} else {
		fmt.Println("REP Error: Id no existe")
	}
}

// reporteEBRs genera el reporte de todos los EBRs y sus particiones lógicas asociadas
func reporteEBRs(particionExtendida Estructuras.Partition, disco *os.File) string {
	cad := ""

	// Leer el primer EBR que está al inicio de la partición extendida
	var ebrActual Estructuras.EBR
	var posEBR int64 = int64(particionExtendida.Part_start)

	if err := Acciones.ReadObject(disco, &ebrActual, posEBR); err != nil {
		fmt.Println("REP Error: No se pudo leer el EBR inicial")
		return ""
	}

	// Contador para numerar los EBRs
	contador := 1

	// Recorrer toda la lista enlazada de EBRs
	for {
		// Añadir información del EBR actual
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='MediumSlateBlue' COLSPAN=\"2\"> EBR #%d </td> \n </tr> \n", contador)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='GhostWhite'> Posición en Disco </td> \n  <td bgcolor='GhostWhite'> %d </td> \n </tr> \n", posEBR)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='Lavender'> Estado </td> \n  <td bgcolor='Lavender'> %s </td> \n </tr> \n", string(ebrActual.EbrP_mount[:]))
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='GhostWhite'> Ajuste </td> \n  <td bgcolor='GhostWhite'> %s </td> \n </tr> \n", string(ebrActual.EbrP_fit[:]))
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='Lavender'> Inicio Partición </td> \n  <td bgcolor='Lavender'> %d </td> \n </tr> \n", ebrActual.EbrP_start)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='GhostWhite'> Tamaño Partición </td> \n  <td bgcolor='GhostWhite'> %d </td> \n </tr> \n", ebrActual.EbrP_size)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='Lavender'> Siguiente EBR </td> \n  <td bgcolor='Lavender'> %d </td> \n </tr> \n", ebrActual.EbrP_next)
		cad += fmt.Sprintf(" <tr>\n  <td bgcolor='GhostWhite'> Nombre Partición </td> \n  <td bgcolor='GhostWhite'> %s </td> \n </tr> \n",
			Estructuras.GetName(string(ebrActual.EbrP_name[:])))

		// Si hay una partición lógica válida en este EBR (tamaño > 0), mostrar info adicional
		if ebrActual.EbrP_size > 0 {
			cad += " <tr>\n  <td bgcolor='SteelBlue' COLSPAN=\"2\"> Partición Lógica </td> \n </tr> \n"
			cad += fmt.Sprintf(" <tr>\n  <td bgcolor='LightCyan'> Nombre </td> \n  <td bgcolor='LightCyan'> %s </td> \n </tr> \n",
				Estructuras.GetName(string(ebrActual.EbrP_name[:])))
			cad += fmt.Sprintf(" <tr>\n  <td bgcolor='PowderBlue'> Inicio </td> \n  <td bgcolor='PowderBlue'> %d </td> \n </tr> \n", ebrActual.EbrP_start)
			cad += fmt.Sprintf(" <tr>\n  <td bgcolor='LightCyan'> Tamaño </td> \n  <td bgcolor='LightCyan'> %d </td> \n </tr> \n", ebrActual.EbrP_size)
			cad += fmt.Sprintf(" <tr>\n  <td bgcolor='PowderBlue'> Tipo </td> \n  <td bgcolor='PowderBlue'> %s </td> \n </tr> \n", string(ebrActual.EbrType[:]))
		} else {
			cad += " <tr>\n  <td bgcolor='LightGray' COLSPAN=\"2\"> EBR sin partición activa </td> \n </tr> \n"
		}

		// Añadir separador entre EBRs
		cad += " <tr>\n  <td bgcolor='#E6E6FA' COLSPAN=\"2\"> </td> \n </tr> \n"

		// Si no hay más EBRs en la cadena, terminamos
		if ebrActual.EbrP_next == -1 {
			break
		}

		// Avanzar al siguiente EBR
		posEBR = int64(ebrActual.EbrP_next)
		if err := Acciones.ReadObject(disco, &ebrActual, posEBR); err != nil {
			fmt.Println("REP Error: Error al leer un EBR en la cadena")
			return cad // Devolver lo que ya tenemos acumulado
		}

		// Incrementar contador
		contador++
	}

	return cad
}
