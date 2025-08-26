# IFT365-APLIT 🚀
Este repositorio es un proyecto de Go base diseñado para que estudiantes y principiantes puedan explorar y entender la estructura de módulos y paquetes en Go. Contiene ejemplos simples de funciones matemáticas y de manipulación de cadenas de texto, organizadas en sus propios paquetes.

## Descripción del Proyecto ✨
El objetivo de este proyecto es ilustrar cómo organizar código Go en módulos y paquetes, cómo definir funciones exportables y cómo importar y utilizar esas funciones desde el programa principal (main.go). Es un punto de partida para construir aplicaciones más complejas y fomentar las buenas prácticas de organización de código desde el principio.

## Estructura del Proyecto 📁
```
IFT365-APLIT/
├── .gitignore             # Reglas para ignorar archivos en Git
├── go.mod                 # Define el módulo Go y sus dependencias
├── README.md              # Este archivo de documentación
├── main.go                # El programa principal que usa los paquetes
├── pkg/                   # Directorio para los paquetes internos del proyecto
│   ├── math_utils/        # Paquete para operaciones matemáticas básicas
│   │   └── basic.go       # Implementación de funciones como Sumar, Restar, Multiplicar, Dividir
│   └── string_utils/      # Paquete para manipulación de cadenas de texto
│       └── manip.go       # Implementación de funciones como InvertirCadena, ContarVocales, EsPalindromo
└── .vscode/               # (Opcional) Configuraciones recomendadas para VS Code
    └── settings.json
```

## Cómo Usar este Repositorio ⚙️
Sigue estos pasos para clonar el repositorio, ejecutar el ejemplo y entender su funcionamiento:

**1. Clonar el Repositorio:**
Abre tu terminal y clona este repositorio:

```
git clone https://github.com/MrOnie/IFT365-APLIT.git
cd IFT365-APLIT
```

Si quieres clonar el repositorio y solamente usar los paquetes, desde tu main.go puedes importar directamente desde el repositorio de GitHub con:

```
import "github.com/MrOnie/IFT365-APLIT/pkg/math_utils"
```

Y no te olvides de usar desde tu proyecto:

```
go get github.com/MrOnie/IFT365-APLIT/pkg/math_utils
```

**2. Inicializar el Módulo (si es necesario):**
Si no se ha inicializado automáticamente o si estás trabajando en un entorno nuevo, ejecuta:

```
go mod tidy
```

Esto asegurará que tu go.mod esté al día con las dependencias.

**3. Ejecutar la Aplicación:**
Desde la raíz del proyecto (IFT365-APLIT/), ejecuta el programa principal:

```
go run .
```

Verás la salida de las funciones de demostración de los paquetes math_utils y string_utils.

**4. Explorar el Código:**

- main.go: Observa cómo se importan los paquetes (import "github.com/MrOnie/IFT365-APLIT/pkg/math_utils") y cómo se llaman sus funciones (math_utils.Sumar(...)).

- pkg/math_utils/basic.go: Revisa la declaración package math_utils y las funciones (Sumar, Restar, etc.). Nota que comienzan con mayúscula para ser exportables.

- pkg/string_utils/manip.go: Similar al paquete math_utils, pero con funciones para cadenas.

## Cómo Contribuir (Opcional para Estudiantes) 💡
¡Te animamos a extender este proyecto! Puedes:

- Añadir más funciones al paquete math_utils (ej: Potencia, RaizCuadrada).

- Añadir más funciones al paquete string_utils (ej: ContarPalabras, RemoverEspacios).

- Crear un nuevo paquete bajo pkg/ (ej: pkg/array_utils con funciones para slices).

- Implementar pruebas unitarias para cada función de los paquetes.

Esto te ayudará a practicar la creación de tu propio código modular en Go.