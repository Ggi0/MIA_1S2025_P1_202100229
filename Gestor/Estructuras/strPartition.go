package Estructuras

import (
	"Gestor/Acciones"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

/*
----------> PARTITION <----------

	Una PARTICION es una división lógica de un disco que
	los sistemas de archivos tratan como una unidad separada.

atributos:

	part_status      - char      - Indica si la partición está MONTADA o no
	part_type        - char      - Indica el tipo de partición, primaria (P) o extendida (E).
	part_fit         - char      - Tipo de ajuste de la partición. B(Best), F (First) o W (worst)
	part_start       - int       - Indica en qué byte del disco inicia la partición
	part_s           - int       - Contiene el tamaño total de la partición en bytes
	part_name        - char[16]  - Nombre de la partición
	part_correlative - int       - Indica el correlativo de la partición este valor será inicialmente -1 hasta que sea montado
	part_id          - char[4]   - Indica el ID de la partición generada al montar esta partición, esto se explicará más adelante
*/
type Partition struct {
	Part_status      [1]byte  // Estado de la partición
	Part_type        [1]byte  // Tipo de partición P E
	Part_fit         [1]byte  // Ajuste de la partición
	Part_start       int32    // Byte de inicio de la partición
	Part_size        int32    // Tamaño de la partición
	Part_name        [16]byte // Nombre de la partición
	Part_correlative int32    // Correlativo de la partición
	Part_id          [4]byte  // ID de la partición
}

// Setear valores de la particion
func (p *Partition) SetInfo(newType string, fit string, newStart int32, newSize int32, name string, correlativo int32) {
	p.Part_size = newSize
	p.Part_start = newStart
	p.Part_correlative = correlativo
	copy(p.Part_name[:], name)
	copy(p.Part_fit[:], fit)
	copy(p.Part_status[:], "I")
	copy(p.Part_type[:], newType)
}

// Metodos de Partition --> para obtener el nombre de la particion
func GetName(nombre string) string {
	posicionNulo := strings.IndexByte(nombre, 0)
	//Si posicionNulo retorna -1 no hay bytes nulos
	if posicionNulo != -1 {
		//guarda la cadena hasta el primer byte nulo (elimina los bytes nulos)
		nombre = nombre[:posicionNulo]
	}
	return nombre
}

func GetId(nombre string) string {
	//si existe id, no contiene bytes nulos
	posicionNulo := strings.IndexByte(nombre, 0)
	//si posicionNulo  no es -1, no existe id.
	if posicionNulo != -1 {
		nombre = "-"
	}
	return nombre
}

func (p *Partition) GetEnd() int32 {
	return p.Part_start + p.Part_size
}

/*
disco *os.File       ---> disco abierto
typePartition string ---> P (primaria), E (extendida), L (logica)
name string          ---> nombre de la particion
size int             ---> tamanio de la particion
unit int             ---> B (bytes), K (kilobytes), M (megabytes)
*/
func EscribirParticion(disco *os.File, typePartition string, name string, size int, unit int, fit string) {
	//Se crea un mbr para cargar el mbr del disco --> la info
	var mbr MBR

	//Guardo el mbr leido --> y se lee desde la posicion (0, 0) --> lee el MBR que tiene el archivo disco
	if err := Acciones.ReadObject(disco, &mbr, 0); err != nil {
		return
	}

	//Si la particion es tipo extendida validar que no exista alguna extendida
	// solo puede haber una por disco
	isPartExtend := false // Indica si se puede usar la particion extendida

	isName := true // Valida si el nombre no se repite (true no se repite)

	if typePartition == "E" {
		for i := 0; i < 4; i++ {
			// leyendo la informacion de las  particiones en el MBR del disco.
			tipo := string(mbr.Mbr_partitions[i].Part_type[:])
			//fmt.Println("tipo ", tipo)
			if tipo != "E" {
				isPartExtend = true
			} else {
				isPartExtend = false
				isName = false
				fmt.Println("ERROR [ F DISK ]: Ya existe una particion extendida")
				fmt.Println("ERROR [ F DISK ]: No se puede crear la nueva particion con nombre: ", name)
				defer disco.Close() // cerrar el disco
				break
			}
		}
	}

	//verificar si  el nombre existe en las particiones primarias o extendida
	if isName {
		for i := 0; i < 4; i++ {
			nombre := GetName(string(mbr.Mbr_partitions[i].Part_name[:]))
			if nombre == name {
				isName = false
				fmt.Println("ERROR [ F DISK ]: Ya existe la particion : ", name)
				fmt.Println("ERROR [ F DISK ]: No se puede crear la nueva particion con nombre: ", name)
				defer disco.Close() // cerrar el disco
				break
			}
		}
	}

	//INGRESO DE PARTICIONES PRIMARIAS Y/O EXTENDIDA (SIN LOGICAS)
	sizeNewPart := size * unit //Tamaño de la nueva particion (tamaño * unidades)
	guardar := false           //Indica si se debe guardar la particion, es decir, escribir en el disco
	var newPart Partition      // nuevo struct tipo particion --> en donde se va guardar la particion

	if (typePartition == "P" || (isPartExtend && typePartition == "E")) && isName { //para que isPartExtend sea true, el tipo de la particion tendra que ser "E"

		// obtener el espacio Real fisico que ocupa el MBR en el disco (el tamaño de la estructura - estructuras.md)
		sizeMBR := int32(binary.Size(mbr))

		//Para manejar los demas ajustes hacer un if del FIT para llamar a la funcion adecuada

		// TODO: validar los ajustes
		// F = primer ajuste;  (BF)
		// B = mejor ajuste;   (FF)
		// else -> peor ajuste (WF --> por defecto)

		//INSERTAR PARTICION (Primer ajuste)
		switch fit {
		case "FF": // busca el primer espacio libre que encuentre
			fmt.Println("fit: ", fit)

		case "BF": // Busca el espacio más pequeño donde quepa
			fmt.Println("fit: ", fit)

		case "WF": // busca el espacio libre mas grande disponible
			// worstfit
			fmt.Println("fit: ", fit)

		default:
			// ERROR
		}

		mbr, newPart = primerAjuste(mbr, typePartition, sizeMBR, int32(sizeNewPart), name, fit) //int32(sizeNewPart) es para castear el int a int32 que es el tipo que tiene el atributo en el struct Partition
		// si la particion es mayor a 0 se puede guardar
		guardar = (newPart.Part_size > 0)

		//escribimos el MBR en el archivo. Lo que no se llegue a escribir en el archivo (aqui) se pierde, es decir, los cambios no se guardan
		if guardar {

			//sobreescribir el mbr --> con la nueva info de la particion
			if err := Acciones.WriteObject(disco, mbr, 0); err != nil {
				return
			}

			//SI es extendida ademas se agrega el ebr de la particion extendida en el disco
			if isPartExtend {
				var ebr EBR // EBR por "defecto"

				ebr.ebrP_start = newPart.Part_start // el nuevo EBR inicia a apartir de donde incia la nueva particion
				ebr.ebrP_next = -1                  // enlazada a otro ebr

				// escribiendo el EBR en la posicion int64(ebr.ebrP_start)
				if err := Acciones.WriteObject(disco, ebr, int64(ebr.ebrP_start)); err != nil {
					return
				}
			}

			// para verificar que lo guardo
			var TempMBR2 MBR
			// se lee de nuevo el MBR del disco
			if err := Acciones.ReadObject(disco, &TempMBR2, 0); err != nil {
				return
			}
			fmt.Println("[ F DISK ]: Particion con nombre " + name + ", de tipo: " + typePartition + " creada exitosamente")
			PrintMBR(TempMBR2)
		} else {
			//Lo podría eliminar pero tendria que modificar en el metodo del ajuste todos los errores para que aparezca el nombre que se intento ingresar como nueva particion
			fmt.Println("ERROR [ F DISK ]: No se puede crear la nueva particion con nombre: ", name)
			defer disco.Close()
		}

	}
	// TODO: else if para ingreso de particiones logicas

}

func primerAjuste(mbr MBR, typee string, sizeMBR int32, sizeNewPart int32, name string, fit string) (MBR, Partition) {

	var newPart Partition // struct de particion
	var noPart Partition  //para revertir el set info (simula volverla null)

	//PARTICION 1 (libre) - (size = 0 no se ha creado) caso1
	if mbr.Mbr_partitions[0].Part_size == 0 {
		newPart.SetInfo(typee, fit, sizeMBR, sizeNewPart, name, 1) // nueva particion
		if mbr.Mbr_partitions[1].Part_size == 0 {
			if mbr.Mbr_partitions[2].Part_size == 0 {
				//caso particion 4 (no existe)
				if mbr.Mbr_partitions[3].Part_size == 0 {
					//859 <= 1024 - 165
					// validar que el tamanio de la particion quepa
					if sizeNewPart <= mbr.Mbr_tamanio-sizeMBR { // tamanio del disco - tamanio de la estructura MBR
						mbr.Mbr_partitions[0] = newPart
					} else {
						newPart = noPart // regresa a un case Particion vacio
						fmt.Println("ERROR [FDISK]: Espacio insuficiente para nueva particion")
					}
				} // else caso 2
			}
		}
		//Fin de 1 no existe

		//PARTICION 2 (no existe)
		/*
			part0 part1 part2 part3
			1	  0		0	  0
		*/
	} else if mbr.Mbr_partitions[1].Part_size == 0 {
		//Si no hay espacio antes de particion 1
		newPart.SetInfo(typee, fit, mbr.Mbr_partitions[0].GetEnd(), sizeNewPart, name, 2) //el nuevo inicio es donde termina 1
		if mbr.Mbr_partitions[2].Part_size == 0 {
			if mbr.Mbr_partitions[3].Part_size == 0 {
				if sizeNewPart <= mbr.Mbr_tamanio-newPart.Part_start {
					mbr.Mbr_partitions[1] = newPart
				} else {
					newPart = noPart
					fmt.Println("ERROR [FDISK]: Espacio insuficiente para nueva particion")
				}
			}
		}
		//Fin particion 2 no existe

		//PARTICION 3
		/*
			part0 part1 part2 part3
			1	  1		0	  0
		*/
	} else if mbr.Mbr_partitions[2].Part_size == 0 {
		//despues de 2
		newPart.SetInfo(typee, fit, mbr.Mbr_partitions[1].GetEnd(), sizeNewPart, name, 3)
		if mbr.Mbr_partitions[3].Part_size == 0 {
			if sizeNewPart <= mbr.Mbr_tamanio-newPart.Part_start {
				mbr.Mbr_partitions[2] = newPart
			} else {
				newPart = noPart
				fmt.Println("ERROR [FDISK]: Espacio insuficiente para nueva particion")
			}
		}
		//Fin particion 3

		//PARTICION 4
		/*
			part0 part1 part2 part3
			1	  1		1	  0
		*/
	} else if mbr.Mbr_partitions[3].Part_size == 0 {
		if sizeNewPart <= mbr.Mbr_tamanio-mbr.Mbr_partitions[2].GetEnd() {
			//despues de 3
			newPart.SetInfo(typee, fit, mbr.Mbr_partitions[2].GetEnd(), sizeNewPart, name, 4)
			mbr.Mbr_partitions[3] = newPart
		} else {
			newPart = noPart
			fmt.Println("ERROR [FDISK]: Espacio insuficiente")
		}
		//Fin particion 4
		/*
			part0 part1 part2 part3
			1	  1		1	  1
		*/
	} else {
		newPart = noPart
		fmt.Println("ERROR [FDISK]: Particiones primarias y/o extendidas ya no disponibles")
	}

	return mbr, newPart
}
