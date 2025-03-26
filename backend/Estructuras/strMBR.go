package Estructuras

import (
	"Gestor/Acciones"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
----------> MBR <----------

	Cuando se crea un nuevo disco este debe contener un MBR,
	este deberá estar en el primer sector del disco.

atributos:

	mbr_tamano         - int          - Tamaño total del disco en bytes
	mbr_fecha_creacion - time         - Fecha y hora de creación del disco
	mbr_dsk_signature  - int          - Número random, que identifica de forma única a cada disco
	dsk_fit            - char         - Tipo de ajuste de la partición. B (Best), F (First) o W (worst)
	mbr_partitions     - partition[4] - Estructura con información de las 4 particiones
*/
type MBR struct {
	Mbr_tamanio        int32        // Tamaño del DISCO en bytes
	Mbr_creation_date  [19]byte     // Fecha y hora de creación del MBR
	Mbr_disk_signature int32        // Firma del disco (ID)
	Mbr_disk_fit       [1]byte      // Tipo de ajuste
	Mbr_partitions     [4]Partition // Particiones del MBR slice del struct Partition
}

/*
Funcion para escribir el MBR en el Disco:

file *os.File   Apuntador a un archivo abierto
tam int         Tamanio del disco
fit string      Tipo de ajuste de la particion
*/
func EscribirMBR(file *os.File, tam int, fit string) (*os.File, error) {
	//obtener hora para el id
	ahora := time.Now()
	//obtener los segundos y minutos
	segundos := ahora.Second()
	minutos := ahora.Minute()
	//hora := ahora.Hour()

	//fmt.Println(hora, minutos, segundos)

	//concatenar los segundos y minutos como una cadena (de 4 digitos)
	cad := fmt.Sprintf("%02d%02d", segundos, minutos)

	//convertir la cadena a numero en un id temporal
	idTmp, err := strconv.Atoi(cad)
	if err != nil {
		fmt.Println("\t ---> ERROR [ MK DISK - mbr ]: La conversion de fecha a entero para id fue incorrecta")
	}

	fmt.Println("\t[ MK DISK - mbr ] ID:", idTmp)

	// Create a new instance of MBR
	var newMBR MBR
	newMBR.Mbr_tamanio = int32(tam)
	newMBR.Mbr_disk_signature = int32(idTmp)
	copy(newMBR.Mbr_disk_fit[:], fit)
	copy(newMBR.Mbr_creation_date[:], ahora.Format("02/01/2006 15:04:05"))

	// Write object in bin file
	if err := Acciones.WriteObject(file, newMBR, 0); err != nil {
		return nil, err
	}

	PrintMBR(newMBR)
	return file, err

}

// Reportes de los Structs
func PrintMBR(data MBR) {
	fmt.Println("\t Disco")
	fmt.Printf("CreationDate: %s, fit: %s, size: %d, id: %d\n", string(data.Mbr_creation_date[:]), string(data.Mbr_disk_fit[:]), data.Mbr_tamanio, data.Mbr_disk_signature)
	for i := 0; i < 4; i++ {
		fmt.Printf("Partition %d: %s, %s, %d, %d, %s, %d\n", i, string(data.Mbr_partitions[i].Part_name[:]), string(data.Mbr_partitions[i].Part_type[:]), data.Mbr_partitions[i].Part_start, data.Mbr_partitions[i].Part_size, string(data.Mbr_partitions[i].Part_fit[:]), data.Mbr_partitions[i].Part_correlative)
	}
}
