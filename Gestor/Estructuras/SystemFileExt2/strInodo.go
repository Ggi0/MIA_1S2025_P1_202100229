package SystemFileExt2

// INODO
type Inode struct {
	I_uid   int32     //ID del usuario propietario del archivo o carpeta
	I_gid   int32     //ID del grupo al que pertenece el archivo o carpeta
	I_size  int32     //tamaño del archivo en bytes
	I_atime [16]byte  //ultima fecha que se leyó el inodo sin modificarlo "02/01/2006 15:04"
	I_ctime [16]byte  //fecha en que se creo el inodo "02/01/2006 15:04"
	I_mtime [16]byte  //ultima fecha en la que se modifica el inodo "02/01/2006 15:04"
	I_block [15]int32 //-1 si no estan usados. los valores del arreglo son: primeros 12 -> bloques directo;: 13 -> bloque simple indirecto; 14->bloque doble indirecto; 15 -> bloque triple indirecto
	I_type  [1]byte   //1 -> archivo; 0 -> carpeta
	I_perm  [3]byte   //permisos del usuario o carpeta
}
