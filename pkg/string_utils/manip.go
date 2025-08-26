package string_utils // Declaramos el nombre del paquete

import "strings" // Importamos el paquete strings para funciones de cadena

// InvertirCadena toma una cadena y devuelve su versión invertida.
func InvertirCadena(s string) string {
	runes := []rune(s) // Convertimos la cadena a un slice de runas para manejar caracteres Unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Intercambiamos los caracteres
	}
	return string(runes) // Convertimos el slice de runas de nuevo a cadena
}

// ContarVocales cuenta el número de vocales (a, e, i, o, u, mayúsculas y minúsculas) en una cadena.
func ContarVocales(s string) int {
	count := 0
	lowerS := strings.ToLower(s) // Convertimos a minúsculas para simplificar la comparación
	vocales := "aeiouáéíóú"      // Incluimos vocales acentuadas si es relevante

	for _, char := range lowerS {
		if strings.ContainsRune(vocales, char) {
			count++
		}
	}
	return count
}

// EsPalindromo verifica si una cadena es un palíndromo (se lee igual al revés).
// Ignora espacios y es insensible a mayúsculas/minúsculas.
func EsPalindromo(s string) bool {
	cleanString := strings.ToLower(strings.ReplaceAll(s, " ", "")) // Limpiamos la cadena
	invertedString := InvertirCadena(cleanString)
	return cleanString == invertedString
}
