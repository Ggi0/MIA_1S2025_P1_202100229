# MIA_1S2025_P1_202100229
Aplicación web para gestionar un sistema de archivos EXT2, usando React para el frontend y Go para el backend. Permite crear/administrar particiones, gestionar archivos/carpetas, manejar permisos de usuarios/grupos y generar reportes. El sistema expone una API REST para las operaciones del sistema de archivos.



# Configuración y Estructura del Backend con Gin

Configuración inicial y la estructura del backend utilizando el framework Gin en Go.

## Instalación de Dependencias
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

## Ejecutar el Backend
Para iniciar el servidor, ejecuta el siguiente comando dentro de la carpeta `Gestor`:

```sh
sudo go run main.go
```

Este comando iniciará el servidor con Gin y habilitará la API para recibir solicitudes.

---

## Notas
- **Modularidad:** La estructura del proyecto permite una mejor organización y escalabilidad.
- **Seguridad:** Se recomienda manejar correctamente los permisos de ejecución, evitando el uso de `sudo` si no es necesario.
- **Manejo de errores:** Se deben implementar validaciones adecuadas en los controladores y servicios para evitar fallos inesperados.

---
# administrción del sistema de archivos

Administración del sistema de archivos: los comandos de este apartado simularan:
		 el formateo de las particiosn
		 administracion de usuarios, carpetas y archivos 
		 -> Debe existir una sesion activa, a exepcion del mkfs y el login


# Estructura General de una Partición EXT2

Una partición EXT2 se organiza en cinco secciones principales:

```
| Superbloque | Bitmap Inodos | Bitmap Bloques | Inodos | Bloques |
```

Cada una cumple una función específica dentro del sistema de archivos.

## 1. Superbloque

El superbloque es como la "tabla de contenidos" del sistema de archivos. Contiene información esencial sobre la estructura de la partición.

### ¿Qué almacena?
- Número total de inodos y bloques
- Cantidad de inodos y bloques libres
- Fecha y hora del último montaje y desmontaje
- Tamaños de inodos y bloques
- Ubicaciones de las diferentes secciones (bitmaps, inodos, bloques)
- Un número mágico (0xEF53) que identifica al sistema como EXT2

**Ejemplo intuitivo:** Es como la primera página de un libro, donde está el índice y la información general sobre el contenido.

## 2. Bitmap de Inodos

Es un mapa de bits que indica cuáles inodos están ocupados y cuáles están libres.

### ¿Qué almacena?
- Cada bit representa un inodo:
  - `0`: Inodo libre
  - `1`: Inodo en uso

**Ejemplo intuitivo:** Piensa en un estacionamiento con 10 espacios. Un bitmap sería una lista donde los espacios ocupados se marcan con `1` y los libres con `0`:  
`[1,1,0,1,0,0,0,1,1,0]`

## 3. Bitmap de Bloques

Similar al bitmap de inodos, pero en lugar de inodos, rastrea bloques de datos.

### ¿Qué almacena?
- Cada bit indica el estado de un bloque:
  - `0`: Bloque libre
  - `1`: Bloque ocupado

**Ejemplo intuitivo:** Si tienes una libreta con 20 páginas y quieres saber cuáles están usadas, podrías hacer una lista de 20 bits donde cada bit representa una página.

## 4. Inodos (Index Nodes)

Los inodos son estructuras que contienen la información sobre cada archivo o carpeta del sistema.

### ¿Qué almacena cada inodo?
- Identificadores de usuario y grupo (UID y GID)
- Tamaño del archivo
- Fechas (creación, modificación, último acceso)
- Permisos de acceso
- Tipo de archivo (archivo regular o carpeta)
- Punteros a los bloques de datos

**Ejemplo intuitivo:** Un inodo es como la ficha bibliográfica de un libro en una biblioteca. No contiene el libro en sí, sino la información sobre él y dónde encontrarlo.

## 5. Bloques

Los bloques son las unidades donde se almacena el contenido real de los archivos o carpetas.

### Tipos de bloques en EXT2
- **Bloques de carpetas**: Almacenan nombres de archivos y la referencia a su inodo.
- **Bloques de archivos**: Contienen los datos reales de un archivo.
- **Bloques de apuntadores**: Se usan para acceder a más bloques de datos en archivos grandes.

**Ejemplo intuitivo:**  
- Si un inodo es la ficha bibliográfica, los bloques son las páginas del libro.  
- Un bloque de carpeta es como una lista con nombres de archivos.  
- Un bloque de archivo es el contenido real (como texto o imágenes).  
- Un bloque de apuntadores es como un índice adicional que te dice "para seguir leyendo, ve a estas otras páginas".

## ¿Cómo funciona la relación entre estos elementos?

1. **Cuando accedes a un archivo:**
   - El sistema busca el inodo correspondiente.
   - El inodo tiene punteros a los bloques donde está el contenido.
   - El sistema lee esos bloques para mostrar el contenido.

2. **Cuando creas un archivo nuevo:**
   - Se busca un inodo libre en el bitmap de inodos.
   - Se buscan bloques libres en el bitmap de bloques.
   - Se actualiza el inodo con los punteros a esos bloques.
   - Se actualiza el superbloque con la nueva información.

## Cálculo de la estructura del sistema EXT2

Para determinar cuántos inodos y bloques se pueden crear en una partición, se usa la siguiente fórmula:

```
tamaño_particion = sizeOf(superblock) + n + 3*n + n*sizeOf(inodos) + 3*n*sizeOf(block)
numero_estructuras = floor(n)
```

### Relación entre inodos y bloques
Por cada inodo (`n`), hay **tres** bloques (`3*n`). Esto significa que en una partición EXT2, los inodos y los bloques siguen una proporción de **1:3**.

---
# Comandos y Estructuras del Sistema de Archivos

## mkdisk
Crea un archivo binario que simula un disco duro virtual.

**Parámetros:**
- **-size** (obligatorio): Tamaño del disco
- **-fit** (opcional): Tipo de ajuste a utilizar (BF/FF/WF)
- **-unit** (opcional): Unidades a utilizar: Kilobytes (K) o Megabytes (M)
- **-path** (obligatorio): Ruta donde se creará el archivo

## fdisk
Maneja las particiones en el disco permitiendo crear, eliminar o modificar particiones.

**Parámetros:**
- **-size** (obligatorio): Tamaño de la partición a crear
- **-unit** (opcional): Unidades a utilizar: Bytes (B), Kilobytes (K) o Megabytes (M)
- **-path** (obligatorio): Ruta donde se encuentra el disco en el que se creará la partición
- **-type** (opcional): Tipo de partición: Primaria (P), Extendida (E) o Lógica (L)
- **-fit** (opcional): Tipo de ajuste de la partición: BF (Best), FF (First) o WF (Worst)
- **-name** (obligatorio): Nombre de la partición

## mount
Monta una partición del disco.


**Formato del ID:** número '29' + números de la partición (part_correlative) + letra

**Parámetros:**
- **-name** (obligatorio): Nombre de la partición
- **-path** (obligatorio): Ruta donde se encuentra el archivo

## rmdisk
Elimina un archivo que representa a un disco duro.

**Parámetros:**
- **-path** (obligatorio): Ruta del archivo que se eliminará

## mkfile
Permite crear un archivo. El propietario será el usuario que actualmente ha iniciado sesión, con permisos 664.

**Parámetros:**
- **-path** (obligatorio): Ruta del archivo que se creará. Si ya existe, muestra un mensaje preguntando si se desea sobreescribir.
- **-r** (opcional): Si se utiliza y las carpetas especificadas por path no existen, se crearán las carpetas padres.
- **-size** (opcional): Tamaño en bytes del archivo. El contenido serán números del 0 al 9 repetidos hasta cumplir el tamaño.
- **-cont** (opcional): Indica un archivo en el disco de la computadora que tendrá el contenido del archivo.

## mkfs
Formateo completo de la partición a ext2. Crea un archivo en la raíz `user.txt` que contendrá los usuarios y contraseñas.

**Parámetros:**
- **-id** (obligatorio): ID generado con el comando mount
- **-type** (opcional): Tipo de formateo. Full: realiza un formateo completo

**Proceso de formateo:**
1. Recibe un ID de partición montada
2. Valida que la partición exista
3. Calcula cuántos inodos y bloques caben en la partición
4. Crea todas las estructuras necesarias del sistema EXT2:
   - Superbloque
   - Bitmaps para inodos y bloques
   - Tabla de inodos
   - Tabla de bloques
5. Configura la carpeta raíz (/) y crea un archivo inicial (users.txt)
6. Escribe todas estas estructuras al disco

## mkdir
Permite crear un folder. El propietario será el usuario que actualmente ha iniciado sesión, con permisos 664.

**Parámetros:**
- **-path** (obligatorio): Ruta de la carpeta que se creará
- **-p** (opcional): Si se utiliza y las carpetas padres en path no existen, se crearán. Sin este parámetro, mostrará error si no existen las carpetas padres.


# Comandos de Administración de Usuarios y Grupos

## LOGIN
Comando para iniciar sesión en el sistema.

**Parámetros:**
- **-user** (obligatorio): Especifica el nombre del usuario que iniciará sesión. Si no se encuentra, muestra un mensaje indicando que el usuario no existe. *Distingue mayúsculas de minúsculas*.
- **-pass** (obligatorio): Indica la contraseña del usuario que inicia sesión. Si no coincide, muestra un mensaje de autenticación fallida. *Distingue mayúsculas y minúsculas*.
- **-id** (obligatorio): Indica el ID de la partición montada de la cual se iniciará sesión. Todas las acciones posteriores se realizarán sobre este ID.

**Observaciones:**
- No se puede iniciar otra sesión sin haber hecho LOGOUT antes. En caso contrario, muestra un mensaje de error.

**Ejemplos:**
```
#Se loguea en el sistema como usuario root
login -user=root -pass=123 -id=062A

#Debe dar error porque ya hay un usuario logueado
login -user="mi usuario" -pass="mi pwd" -id=062A
```

## LOGOUT
Comando para cerrar sesión.

**Características:**
- No recibe parámetros
- Debe haber una sesión activa para poder utilizarlo, si no, muestra un mensaje de error

**Ejemplos:**
```
#Termina la sesión del usuario
Logout

#Si se vuelve a ejecutar deberá mostrar un error ya que no hay sesión actualmente
Logout
```

## MKGRP
Crea un grupo para los usuarios de la partición y lo guarda en el archivo users.txt.

**Características:**
- Solo lo puede utilizar el usuario **root**
- Si el grupo ya existe, muestra un mensaje de error
- Distingue entre mayúsculas y minúsculas

**Parámetros:**
- **-name** (obligatorio): Indica el nombre que tendrá el grupo

**Ejemplo:**
```
#Crea el grupo usuarios en la partición de la sesión actual
mkgrp -name=usuarios
```

**Formato en users.txt:**
```
1, G, root
1, U, root, root, 123
2, G, usuarios
```

## RMGRP
Elimina un grupo para los usuarios de la partición.

**Características:**
- Solo lo puede utilizar el usuario **root**
- Si el grupo no existe, muestra un mensaje de error

**Parámetros:**
- **-name** (obligatorio): Indica el nombre del grupo a eliminar

**Ejemplo:**
```
#Elimina el grupo usuarios en la partición de la sesión actual
rmgrp -name=usuarios

#Debe mostrar mensaje de error ya que el grupo no existe porque ya fue eliminado
rmgrp -name=usuarios
```

**Formato en users.txt después de eliminar:**
```
1, G, root
1, U, root, root, 123
0, G, usuarios
```

## MKUSR
Crea un usuario en la partición.

**Características:**
- Solo lo puede ejecutar el usuario **root**

**Parámetros:**
- **-user** (obligatorio): Indica el nombre del usuario a crear. Máximo 10 caracteres. Si ya existe, muestra un error.
- **-pass** (obligatorio): Indica la contraseña del usuario. Máximo 10 caracteres.
- **-grp** (obligatorio): Indica el grupo al que pertenece el usuario. Máximo 10 caracteres. Debe existir en la partición.

**Ejemplo:**
```
#Crea usuario user1 en el grupo 'usuarios'
mkusr -user=user1 -pass=usuario -grp=usuarios

#Debe mostrar mensaje de error ya que el usuario ya existe independientemente que esté en otro grupo
mkusr -user=user1 -pass=usuario -grp=usuarios2
```

**Formato en users.txt:**
```
1, G, root
1, U, root, root, 123
2, G, usuarios
2, U, usuarios, user1, usuario
```

## RMUSR
Elimina un usuario en la partición.

**Características:**
- Solo lo puede ejecutar el usuario **root**

**Parámetros:**
- **-user** (obligatorio): Indica el nombre del usuario a eliminar. Si no existe, muestra un error.

--- 

## Visión general del disco con particiones

Primero, veamos cómo se organiza un disco completo:

```
┌────────────────────────────────────────────────────────────────┐
│                            DISCO                               │
├────────┬────────────┬───────────────┬────────────┬─────────────┤
│  MBR   │ Partición 1│  Partición 2  │ Partición 3│ Partición 4 │
│(92bytes)│   (EXT2)  │  (Sin formato)│  (EXT2)    │ (Espacio    │
│        │            │               │            │  libre)     │
└────────┴────────────┴───────────────┴────────────┴─────────────┘
```

## Estructura detallada de una partición formateada con EXT2

Ahora, veamos en detalle cómo se organiza una partición formateada con EXT2:

```
┌─────────────────────────────────────── Partición con EXT2 ────────────────────────────────────────┐
│                                                                                                   │
│  ┌────────────┐  ┌─────────────┐  ┌─────────────┐  ┌────────────────────┐  ┌───────────────────┐  │
│  │ Superbloque│  │ Bitmap      │  │ Bitmap      │  │     Tabla de       │  │    Tabla de       │  │
│  │ (92 bytes) │  │ de Inodos   │  │ de Bloques  │  │     Inodos         │  │    Bloques        │  │
│  │            │  │ (n bytes)   │  │ (3n bytes)  │  │   (n * 124 bytes)  │  │  (3n * 64 bytes)  │  │
│  └────────────┘  └─────────────┘  └─────────────┘  └────────────────────┘  └───────────────────┘  │
│                                                                                                   │
└───────────────────────────────────────────────────────────────────────────────────────────────────┘
```

Donde:
- **n** es el número de inodos que caben en la partición (calculado en la función `Mkfs`)
- Cada inodo ocupa 124 bytes
- Cada bloque ocupa 64 bytes

## Detalle de cada estructura principal

### 1. Superbloque (92 bytes)
Es la estructura que contiene toda la información del sistema de archivos:

```
┌───────────────────── Superbloque (92 bytes) ─────────────────────┐
│ S_filesystem_type   (4 bytes) - Tipo de sistema (2 = EXT2)       │
│ S_inodes_count      (4 bytes) - Total de inodos                  │
│ S_blocks_count      (4 bytes) - Total de bloques                 │
│ S_free_blocks_count (4 bytes) - Bloques libres                   │
│ S_free_inodes_count (4 bytes) - Inodos libres                    │
│ S_mtime             (19 bytes) - Fecha de montaje                │
│ S_umtime            (19 bytes) - Fecha de desmontaje             │
│ S_mnt_count         (4 bytes) - Veces montado                    │
│ S_magic             (4 bytes) - Número mágico (0xEF53)           │
│ S_inode_size        (4 bytes) - Tamaño de inodo (124)            │
│ S_block_size        (4 bytes) - Tamaño de bloque (64)            │
│ S_first_ino         (4 bytes) - Primer inodo libre               │
│ S_first_blo         (4 bytes) - Primer bloque libre              │
│ S_bm_inode_start    (4 bytes) - Inicio bitmap inodos             │
│ S_bm_block_start    (4 bytes) - Inicio bitmap bloques            │
│ S_inode_start       (4 bytes) - Inicio tabla inodos              │
│ S_block_start       (4 bytes) - Inicio tabla bloques             │
└──────────────────────────────────────────────────────────────────┘
```

### 2. Bitmap de Inodos (n bytes)
Cada bit representa un inodo:
- 0 = libre
- 1 = ocupado

```
┌─── Bitmap Inodos (n bytes) ───┐
│ 1 1 0 0 0 0 0 0 0 0 ...       │
└───────────────────────────────┘
  ^ ^
  | └─ Inodo 1 (users.txt)
  └─── Inodo 0 (carpeta raíz)
```

### 3. Bitmap de Bloques (3n bytes)
Cada bit representa un bloque:
- 0 = libre
- 1 = ocupado

```
┌─── Bitmap Bloques (3n bytes) ─┐
│ 1 1 0 0 0 0 0 0 0 0 ...       │
└───────────────────────────────┘
  ^ ^
  | └─ Bloque 1 (contenido de users.txt)
  └─── Bloque 0 (entradas de carpeta raíz)
```

### 4. Tabla de Inodos (n * 124 bytes)
Cada inodo ocupa 124 bytes:

```
┌───────────────────── Inodo (124 bytes) ─────────────────────┐
│ I_uid      (4 bytes) - ID de usuario propietario            │
│ I_gid      (4 bytes) - ID de grupo propietario              │
│ I_size     (4 bytes) - Tamaño del archivo                   │
│ I_atime    (16 bytes) - Fecha de último acceso              │
│ I_ctime    (16 bytes) - Fecha de creación                   │
│ I_mtime    (16 bytes) - Fecha de modificación               │
│ I_block    (60 bytes) - 15 apuntadores a bloques (4c/u)     │
│ I_type     (1 byte) - Tipo (0=carpeta, 1=archivo)           │
│ I_perm     (3 bytes) - Permisos (ej: 664)                   │
└─────────────────────────────────────────────────────────────┘
```

### 5. Tabla de Bloques (3n * 64 bytes)
Hay tres tipos principales de bloques:

#### a) Bloque de Carpeta (64 bytes)
```
┌───────────────── Bloque de Carpeta (64 bytes) ─────────────────┐
│ ┌───────────────────┐ ┌───────────────────┐                    │
│ │ B_name: "."       │ │ B_name: ".."      │                    │
│ │ B_inodo: 0        │ │ B_inodo: 0        │                    │
│ └───────────────────┘ └───────────────────┘                    │
│ ┌───────────────────┐ ┌───────────────────┐                    │
│ │ B_name:"users.txt"│ │ B_name: ""        │                    │
│ │ B_inodo: 1        │ │ B_inodo: -1       │                    │
│ └───────────────────┘ └───────────────────┘                    │
└────────────────────────────────────────────────────────────────┘
```

#### b) Bloque de Archivo (64 bytes)
```
┌─────────────── Bloque de Archivo (64 bytes) ───────────────┐
│                                                            │
│ B_content: "1,G,root\n1,U,root,root,123\n"                 │
│                                                            │
└────────────────────────────────────────────────────────────┘
```

#### c) Bloque de Apuntadores (64 bytes)
```
┌───────────── Bloque de Apuntadores (64 bytes) ─────────────┐
│ B_pointers[0]: X                                           │
│ B_pointers[1]: X                                           │
│ B_pointers[2]: X                                           │
│ ...                                                        │
│ B_pointers[15]: X                                          │
└────────────────────────────────────────────────────────────┘
```

## Ejemplo visual del sistema de archivos inicial

Después de formatear con MKFS, así es cómo se ve el sistema de archivos:

```
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                             SISTEMA DE ARCHIVOS EXT2                                    │
├────────────┬─────────────┬─────────────┬────────────────────┬───────────────────────────┤
│ Superbloque│ Bitmap      │ Bitmap      │     Tabla de       │       Tabla de            │
│            │ de Inodos   │ de Bloques  │     Inodos         │       Bloques             │
├────────────┼─────────────┼─────────────┼────────────────────┼───────────────────────────┤
│            │ 1 1 0 0 ... │ 1 1 0 0 ... │ ┌─────────────┐    │  ┌───────────────────┐    │
│  Contiene  │             │             │ │  Inodo 0    │    │  │     Bloque 0      │    │
│  metadata  │ (inodos     │ (bloques    │ │  (carpeta /)│    │  │  (carpeta raíz)   │    │
│  del       │  0 y 1      │  0 y 1      │ └─────────────┘    │  └───────────────────┘    │
│  sistema   │  en uso)    │  en uso)    │ ┌─────────────┐    │  ┌───────────────────┐    │
│            │             │             │ │  Inodo 1    │    │  │     Bloque 1      │    │
│            │             │             │ │ (users.txt) │    │  │ (datos users.txt) │    │
│            │             │             │ └─────────────┘    │  └───────────────────┘    │
│            │             │             │      ...           │         ...               │
└────────────┴─────────────┴─────────────┴────────────────────┴───────────────────────────┘
```

## La relación entre estructuras

Veamos cómo se conectan las estructuras para formar el sistema de archivos:

```
                                  SUPERBLOQUE
                                       │
                                       │ (contiene ubicaciones)
                                       ▼
              ┌─────────────────┬─────────────────┬─────────────────┐
              │                 │                 │                 │
              ▼                 ▼                 ▼                 ▼
        BITMAP INODOS    BITMAP BLOQUES     TABLA INODOS      TABLA BLOQUES
              │                 │                 │                 │
              │                 │                 │                 │
              │                 │                 ▼                 │
              │                 │          ┌─────────────┐          │
              │                 │          │  Inodo 0    │          │
              │                 │          │ (carpeta /) │          │
              │                 │          └─────────────┘          │
              │                 │                 │                 │
              │                 │                 │ I_block[0]=0    │
              │                 │                 │                 │
              │                 │                 │                 ▼
              │                 │                 │           ┌───────────────┐
              │                 │                 └──────────►│   Bloque 0    │
              │                 │                             │ (carpeta raíz)│
              │                 │                             └───────────────┘
              │                 │                                     │
              │                 │                                     │ B_content[2].B_inodo=1
              │                 │                                     ▼
              │                 │                             ┌───────────────┐
              │                 │                             │   Inodo 1     │
              │                 │                             │  (users.txt)  │
              │                 │                             └───────────────┘
              │                 │                                     │
              │                 │                                     │ I_block[0]=1
              │                 │                                     ▼
              │                 │                             ┌───────────────┐
              │                 │                             │   Bloque 1    │
              │                 │                             │  (users.txt)  │
              │                 │                             └───────────────┘
```

## Papel de los structs en el sistema

Los structs que defines en Go tienen dos propósitos principales:

1. **Representación en memoria**: Permiten manipular las estructuras del sistema de archivos desde tu programa.



El sistema EXT2 es un sistema de archivos estructurado que almacena tanto los datos como los metadatos en un formato específico dentro de una partición. Los structs que defines en Go sirven tanto para representar esas estructuras en memoria (para que tu programa pueda manipularlas) como para definir exactamente cómo se almacenan en el disco.

Cuando formateas una partición con MKFS, lo que haces es escribir todas estas estructuras en posiciones específicas dentro de la partición, creando un sistema de archivos vacío con solo la carpeta raíz y el archivo users.txt inicial.
