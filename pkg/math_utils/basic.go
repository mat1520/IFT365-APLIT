package math_utils // Declaramos el nombre del paquete

import "fmt" // Importamos el paquete fmt para la función de error

// Sumar toma dos enteros y devuelve su suma.
// Las funciones que empiezan con mayúscula son exportables (públicas)
func Sumar(a, b int) int {
	return a + b
}

// Restar toma dos enteros y devuelve su resta.
func Restar(a, b int) int {
	return a - b
}

// Multiplicar toma dos enteros y devuelve su producto.
func Multiplicar(a, b int) int {
	return a * b
}

// Dividir toma dos enteros y devuelve su cociente como float64.
// Devuelve un error si el divisor es cero para evitar un pánico.
func Dividir(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: no se puede dividir por cero")
	}
	// Realizamos la división con float64 para obtener un resultado más preciso
	return float64(a) / float64(b), nil
}
