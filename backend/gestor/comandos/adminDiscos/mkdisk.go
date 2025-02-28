package comandos

import (
	"fmt"
)

/*
mkdisk

	-size (obligatoria)
	-fit  (opcional)
	-unit (opcional)
	-path (obligatoria)
*/
func Mkdisk(parametros []string) {
	fmt.Println("desde el metodo: ", parametros)

}
