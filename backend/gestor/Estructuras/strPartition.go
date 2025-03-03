package Estructuras

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
	Part_type        [1]byte  // Tipo de partición
	Part_fit         [1]byte  // Ajuste de la partición
	Part_start       int32    // Byte de inicio de la partición
	Part_size        int32    // Tamaño de la partición
	Part_name        [16]byte // Nombre de la partición
	Part_correlative int32    // Correlativo de la partición
	Part_id          [4]byte  // ID de la partición
}
