# MIA_1S2025_P1_202100229
Aplicación web para gestionar un sistema de archivos EXT2, usando React para el frontend y Go para el backend. Permite crear/administrar particiones, gestionar archivos/carpetas, manejar permisos de usuarios/grupos y generar reportes. El sistema expone una API REST para las operaciones del sistema de archivos.



# backend en go utiliznado gin


1. Configuración inicial del backend con Gin
Primero, necesitamos instalar Gin y configurar el backend:
 
 En la raíz de tu proyecto Go, en la carpeta backend


go get github.com/gin-gonic/gin
go get github.com/rs/cors/wrapper/gin



estructura del backend:


Gestor/
  ├── main.go
  ├── controllers/
  │   └── comandos_controller.go
  ├── services/
  │   └── analizador_service.go
  ├── models/
  │   └── respuesta.go
  ├── Comandos/
  │   └── AdminDiscos/
  │       ├── mkdisk.go
  │       ├── rmdisk.go
  │       ├── fdisk.go
  │       └── ...
  └── ...