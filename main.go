package main // El paquete 'main' es el punto de entrada de la aplicación

import (
	"fmt" // Para imprimir en la consola
	// Importamos nuestros paquetes personalizados.
	// El path de importación debe coincidir con el nombre de tu módulo y la ruta del paquete.
	"github.com/MrOnie/IFT365-APLIT/pkg/math_utils"
	//"./IFT365-APLIT/pkg/math_utils"
	"github.com/MrOnie/IFT365-APLIT/pkg/string_utils"
	//"./IFT365-APLIT/pkg/string_utils"
)

func main() {
	fmt.Println("--- Demostración de Paquetes en Go ---")

	// --- Demostración del Paquete math_utils ---
	fmt.Println("\n## Funciones Matemáticas (pkg/math_utils)")
	num1 := 20
	num2 := 4

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

	_, err = math_utils.Dividir(num1, 0)
	if err != nil {
		fmt.Printf("Intento de división por cero: %s\n", err)
	}

	// --- Demostración del Paquete string_utils ---
	fmt.Println("\n## Funciones de Cadenas (pkg/string_utils)")
	cadenaOriginal := "Hola Go!"
	cadenaPalindromo := "Anita lava la tina"

	// Invertir cadena
	fmt.Printf("Cadena original: \"%s\"\n", cadenaOriginal)
	fmt.Printf("Cadena invertida: \"%s\"\n", string_utils.InvertirCadena(cadenaOriginal))

	// Contar vocales
	fmt.Printf("Número de vocales en \"%s\": %d\n", cadenaOriginal, string_utils.ContarVocales(cadenaOriginal))

	// Verificar palíndromo
	fmt.Printf("¿\"%s\" es un palíndromo? %t\n", cadenaPalindromo, string_utils.EsPalindromo(cadenaPalindromo))
	fmt.Printf("¿\"%s\" es un palíndromo? %t\n", cadenaOriginal, string_utils.EsPalindromo(cadenaOriginal))

	fmt.Println("\n--- Fin de la Demostración ---")
}
