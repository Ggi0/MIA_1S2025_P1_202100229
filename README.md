# MIA_1S2025_P1_202100229
Aplicación web para gestionar un sistema de archivos EXT2, usando React para el frontend y Go para el backend. Permite crear/administrar particiones, gestionar archivos/carpetas, manejar permisos de usuarios/grupos y generar reportes. El sistema expone una API REST para las operaciones del sistema de archivos.



# 📌 Configuración y Estructura del Backend con Gin

Configuración inicial y la estructura del backend utilizando el framework Gin en Go.

## 🚀 Instalación de Dependencias
Antes de comenzar, asegúrate de instalar Gin y otras dependencias necesarias ejecutando los siguientes comandos en la raíz de tu proyecto:

```sh
# Instalar Gin
go get github.com/gin-gonic/gin

# Middleware para manejo de CORS
go get github.com/rs/cors/wrapper/gin
```

---

## Estructura del Proyecto

El backend está organizado en diferentes carpetas para mantener un código limpio y modular:

```
Gestor/
  ├── main.go                       # Punto de entrada principal de la aplicación
  ├── controllers/                  # Controladores, gestionan las solicitudes y lógica principal
  │   ├── comandos_controller.go    # Controlador específico para manejar comandos
  ├── services/                     # Servicios, lógica de negocio y procesos
  │   └── analizador_service.go     # Lógica de análisis principal
  ├── models/                       # Definición de estructuras de datos y modelos
  │   ├── respuesta.go              # Modelo para estructurar las respuestas
  │   ├── errorCap.go               # Modelo para manejar errores capturados
  ├── Comandos/                     # Módulo de comandos organizados por categorías
  │   ├── AdminDiscos/              # Comandos relacionados con discos
  │   │   ├── mkdisk.go             # Crear discos
  │   │   ├── rmdisk.go             # Eliminar discos
  │   │   ├── fdisk.go              # Configurar particiones
  │   │   └── ...                   # Otros comandos relacionados
  │   ├── AdminSistemaArchivos/     # Comandos relacionados con sistemas de archivos
  │       ├── mount.go              # Montar sistemas de archivos
  │       └── ...                   # Más comandos
  ├── Estructuras/                  # Estructuras específicas del sistema
  │   ├── strMBR.go                 # Estructura para MBR (Master Boot Record)
  │   ├── strEBR.go                 # Estructura para EBR (Extended Boot Record)
  │   ├── strPartition.go           # Estructura para particiones
  │   └── estructuras.md            # Documentación sobre las estructuras
  ├── utils/                        # Utilidades auxiliares y herramientas
  │   └── logger_service.go         # Servicio para manejo de registros
  ├── go.mod                        # Archivo del módulo Go
  └── go.sum                        # Hashes de dependencias de Go
```

---

## ▶️ Ejecutar el Backend
Para iniciar el servidor, ejecuta el siguiente comando dentro de la carpeta `Gestor`:

```sh
sudo go run main.go
```

Este comando iniciará el servidor con Gin y habilitará la API para recibir solicitudes.

---

## 📌 Notas
- **Modularidad:** La estructura del proyecto permite una mejor organización y escalabilidad.
- **Seguridad:** Se recomienda manejar correctamente los permisos de ejecución, evitando el uso de `sudo` si no es necesario.
- **Manejo de errores:** Se deben implementar validaciones adecuadas en los controladores y servicios para evitar fallos inesperados.

---


