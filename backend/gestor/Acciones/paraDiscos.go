package Acciones

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/*
Crear el archivo .mia que simulara el DISCO

	path -> es la ruta en donde se hara el archivo
	nombreDisco -> es el nombre del archivo

	returna un error, en dado caso falla la creacion
*/
func CrearDisco(path string, nombreDisco string) error { // retorna un error

	// Detectar si estamos en macOS
	isMacOS := runtime.GOOS == "darwin"

	// Variable para guardar la ruta procesada
	var processedPath string

	/*

		El proyecto está diseñado para ejecutarse principalmente
		en Linux donde la ruta /home es completamente accesible,

	*/
	if isMacOS && strings.HasPrefix(path, "/home") {
		// En macOS, redirigir las rutas /home a una carpeta en el directorio del usuario
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf(" \n --> MKDISK (CrearDisco), ERROR: al obtener el directorio del usuario: %w", err)
		}

		// Reemplazar /home con el directorio simulado en la carpeta del usuario
		// Por ejemplo, /home/user/archivo.mia se convierte en /Users/gio/home/user/archivo.mia
		relPath := strings.TrimPrefix(path, "/home")
		processedPath = filepath.Join(homeDir, "home", relPath)
	} else {
		// En Linux o si la ruta no comienza con /home, usar la ruta original
		processedPath = path
	}

	fmt.Println("La path que entra al crear un disco: `", path, "`") // path original
	fmt.Println("La path procesada: `", processedPath, "`")          // path funcional en macOS

	// Verificar si el archivo ya existe ANTES de crear directorios
	if _, err := os.Stat(processedPath); err == nil {
		// El archivo ya existe, retornar un error
		errorMsg := fmt.Sprintf(" \n --> MKDISK (CrearDisco), ERROR: el disco '%s' ya existe en la ruta '%s'", nombreDisco, path)
		fmt.Println(errorMsg)
		return err
	} else if !os.IsNotExist(err) {
		// Otro error al verificar el archivo
		return fmt.Errorf("error al verificar si el disco existe: %w", err)
	}

	// asegurar que exista la ruta (el directorio) creando la ruta
	// Obtener la carpeta donde se guardará el archivo para asegurarnos de que existe antes de crearlo.
	dir := filepath.Dir(processedPath) // extrae la ruta del directorio que contiene el archivo.

	// Crear el directorio si no existe y todas sus carpetas padre si no existen.
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println(" \n --> MKDISK (CrearDisco), ERROR: al crear el disco: ", err)
		return err
	}

	// Verificar si el archivo ya existe
	newFile, err := os.Create(processedPath) // Crear el archivo si no existe
	if err != nil {
		fmt.Println(" \n --> MKDISK (CrearDisco), ERROR: al crear el disco: ", err)
		return err
	}
	defer newFile.Close() // Cierra el archivo

	fmt.Println(" \n --> MKDISK, disco:", nombreDisco, "creado EXITOSAMENTE.")
	return nil

}
