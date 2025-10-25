package main

import (
	"fmt"
)

// modPow calcula la potenciación modulada a^b mod c
// Utiliza la fórmula recursiva:
// a^b mod c = 1 si b = 0
// a^b mod c = ((a mod c) × (a^(b-1) mod c)) mod c si b > 0
func modPow(a, b, c int) int {
	// Caso base: cualquier número elevado a 0 es 1
	if b == 0 {
		return 1
	}
	
	// Caso recursivo:
	// Calculamos (a mod c) para reducir el tamaño de a
	aMod := a % c
	
	// Calculamos recursivamente a^(b-1) mod c
	prevPow := modPow(a, b-1, c)
	
	// Aplicamos la fórmula: ((a mod c) × (a^(b-1) mod c)) mod c
	result := (aMod * prevPow) % c
	
	return result
}

func main() {
	// Ejemplos de uso
	fmt.Println("Ejemplos de potenciación modulada:")
	fmt.Println()
	
	// Ejemplo 1: 2^5 mod 13
	a, b, c := 2, 5, 13
	result := modPow(a, b, c)
	fmt.Printf("%d^%d mod %d = %d\n", a, b, c, result)
	
	// Ejemplo 2: 3^7 mod 10
	a, b, c = 3, 7, 10
	result = modPow(a, b, c)
	fmt.Printf("%d^%d mod %d = %d\n", a, b, c, result)
	
	// Ejemplo 3: 5^0 mod 7
	a, b, c = 5, 0, 7
	result = modPow(a, b, c)
	fmt.Printf("%d^%d mod %d = %d\n", a, b, c, result)
	
	// Ejemplo 4: 7^10 mod 13
	a, b, c = 7, 10, 13
	result = modPow(a, b, c)
	fmt.Printf("%d^%d mod %d = %d\n", a, b, c, result)
	
	// Ejemplo 5: 123^456 mod 789
	a, b, c = 123, 456, 789
	result = modPow(a, b, c)
	fmt.Printf("%d^%d mod %d = %d\n", a, b, c, result)
}
