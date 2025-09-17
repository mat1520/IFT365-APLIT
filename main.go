package main // El paquete 'main' es el punto de entrada de la aplicación

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mat1520/IFT365-APLIT/pkg/math_utils"
	"github.com/mat1520/IFT365-APLIT/pkg/net_simplify"
	"github.com/mat1520/IFT365-APLIT/pkg/string_utils"
)

func main() {
	fmt.Println("--- Demostración de Paquetes en Go ---")

	// --- Demostración del Paquete math_utils ---
	fmt.Println("\n## Funciones Matemáticas (pkg/math_utils)")

	// Variables Definidas (anteriormente usadas puedes quitar el comentario si quieres testear con valores fijos, Recuerda comentar las otras variables)
	//num1 := 20
	//num2 := 4

	// Variables sin definir
	var num1 int
	var num2 int

	// Solicitar entrada del usuario de las variables sin definir, agregando un manejo de errores
	fmt.Print("Ingrese el primer número (entero): ")
	_, err := fmt.Scanf("%d", &num1)
	if err != nil {
		fmt.Println("Error al leer el número. Asegúrese de ingresar un entero válido.")
		return
	}

	fmt.Print("Ingrese el segundo número (entero): ")
	_, err = fmt.Scanf("%d", &num2)
	if err != nil {
		fmt.Println("Error al leer el número. Asegúrese de ingresar un entero válido.")
		return
	}

	// Suma
	fmt.Printf("Suma de %d y %d: %d\n", num1, num2, math_utils.Sumar(num1, num2))

	// Resta
	fmt.Printf("Resta de %d y %d: %d\n", num1, num2, math_utils.Restar(num1, num2))

	// Multiplicación
	fmt.Printf("Multiplicación de %d y %d: %d\n", num1, num2, math_utils.Multiplicar(num1, num2))

	// División (con manejo de errores)
	division, err := math_utils.Dividir(num1, num2)
	if err != nil {
		fmt.Printf("Error al dividir: %s\n", err)
	} else {
		fmt.Printf("División de %d y %d: %.2f\n", num1, num2, division)
	}

	// Intento de división por cero para demostrar el manejo de errores
	_, err = math_utils.Dividir(num1, 0)
	if err != nil {
		fmt.Printf("Intento de división por cero: %s\n", err)
	}

	// --- Demostración del Paquete string_utils ---
	fmt.Println("\n## Funciones de Cadenas (pkg/string_utils)")
	// Variables Definidas (anteriormente usadas puedes quitar el comentario si quieres testear con valores fijos, Recuerda comentar las otras variables)
	//cadenaOriginal := "Hola Go!"
	//cadenaPalindromo := "Anita lava la tina"

	// Variables sin definir(Las variables se definen aquí para que el usuario pueda ingresarlas)
	var cadenaOriginal string
	var cadenaPalindromo string

	// Solicitar entrada del usuario de las variables sin definir
	fmt.Print("Ingrese una cadena para invertir y contar vocales: ")
	_, err = fmt.Scanf("%s", &cadenaOriginal)
	if err != nil {
		fmt.Println("Error.")
		return
	}

	fmt.Print("Ingrese una cadena para verificar si es palíndromo: ")
	_, err = fmt.Scanf("%s", &cadenaPalindromo)
	if err != nil {
		fmt.Println("Error.")
		return
	}

	// Invertir cadena
	fmt.Printf("Cadena original: \"%s\"\n", cadenaOriginal)
	fmt.Printf("Cadena invertida: \"%s\"\n", string_utils.InvertirCadena(cadenaOriginal))

	// Contar vocales
	fmt.Printf("Número de vocales en \"%s\": %d\n", cadenaOriginal, string_utils.ContarVocales(cadenaOriginal))

	// Verificar palíndromo
	fmt.Printf("¿\"%s\" es un palíndromo? %t\n", cadenaPalindromo, string_utils.EsPalindromo(cadenaPalindromo))
	fmt.Printf("¿\"%s\" es un palíndromo? %t\n", cadenaOriginal, string_utils.EsPalindromo(cadenaOriginal))

	// --- Demostración del Paquete net_simplify ---
	fmt.Println("\n## Funciones de Red (pkg/net_simplify)")

	// registrar handler para la ruta raíz (evita 404)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Servidor iniciado correctamente prueba ingresando a http://localhost:8080/")
	})

	srv := netsimplify.CreateServer(":8080")

	// channel para recibir el resultado de ListenAndServe
	done := make(chan error, 1)
	go func() {
		done <- netsimplify.StartServer(srv)
	}()

	fmt.Println("Servidor creado en el puerto 8080. Ingrese '1' para detenerlo.")

	// esperar a que el usuario ingrese '1'
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Ingrese '1' para detener el servidor: ")
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == "1" {
			break
		}
		fmt.Println("Entrada no válida. Intente de nuevo.")
	}

	// Detener el servidor y esperar el resultado de la goroutine
	if stopErr := netsimplify.StopServer(srv); stopErr != nil {
		fmt.Println("Error al detener servidor:", stopErr)
	}

	if err := <-done; err != nil && err != http.ErrServerClosed {
		fmt.Println("StartServer devolvió error:", err)
	} else {
		fmt.Println("Servidor detenido correctamente.")
	}

	fmt.Println("\n--- Fin de la Demostración ---")

}
