package main

import (
	"fmt"
)

// isMagicMatrix verifica si una matriz cuadrada es mágica
// Una matriz es mágica si la suma de todas sus filas, columnas
// y las dos diagonales son iguales
func isMagicMatrix(matrix [][]int) bool {
	n := len(matrix)
	
	// Verificar que la matriz sea cuadrada
	for i := 0; i < n; i++ {
		if len(matrix[i]) != n {
			return false
		}
	}
	
	// Calcular la suma de la primera fila como referencia
	magicSum := 0
	for j := 0; j < n; j++ {
		magicSum += matrix[0][j]
	}
	
	// Verificar la suma de todas las filas
	for i := 1; i < n; i++ {
		rowSum := 0
		for j := 0; j < n; j++ {
			rowSum += matrix[i][j]
		}
		if rowSum != magicSum {
			return false
		}
	}
	
	// Verificar la suma de todas las columnas
	for j := 0; j < n; j++ {
		colSum := 0
		for i := 0; i < n; i++ {
			colSum += matrix[i][j]
		}
		if colSum != magicSum {
			return false
		}
	}
	
	// Verificar la diagonal principal (de arriba-izquierda a abajo-derecha)
	diagSum1 := 0
	for i := 0; i < n; i++ {
		diagSum1 += matrix[i][i]
	}
	if diagSum1 != magicSum {
		return false
	}
	
	// Verificar la diagonal secundaria (de arriba-derecha a abajo-izquierda)
	diagSum2 := 0
	for i := 0; i < n; i++ {
		diagSum2 += matrix[i][n-1-i]
	}
	if diagSum2 != magicSum {
		return false
	}
	
	// Si todas las verificaciones pasaron, la matriz es mágica
	return true
}

// printMatrix imprime una matriz de forma legible
func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// Ejemplo 1: Matriz mágica del enunciado (3x3)
	magic1 := [][]int{
		{8, 1, 6},
		{3, 5, 7},
		{4, 9, 2},
	}
	
	fmt.Println("Matriz 1:")
	printMatrix(magic1)
	fmt.Printf("¿Es mágica? %v\n\n", isMagicMatrix(magic1))
	
	// Ejemplo 2: Matriz NO mágica (3x3)
	notMagic := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	fmt.Println("Matriz 2:")
	printMatrix(notMagic)
	fmt.Printf("¿Es mágica? %v\n\n", isMagicMatrix(notMagic))
	
	// Ejemplo 3: Matriz mágica 4x4
	magic2 := [][]int{
		{16, 2, 3, 13},
		{5, 11, 10, 8},
		{9, 7, 6, 12},
		{4, 14, 15, 1},
	}
	
	fmt.Println("Matriz 3:")
	printMatrix(magic2)
	fmt.Printf("¿Es mágica? %v\n\n", isMagicMatrix(magic2))
	
	// Ejemplo 4: Matriz mágica 5x5
	magic3 := [][]int{
		{17, 24, 1, 8, 15},
		{23, 5, 7, 14, 16},
		{4, 6, 13, 20, 22},
		{10, 12, 19, 21, 3},
		{11, 18, 25, 2, 9},
	}
	
	fmt.Println("Matriz 4:")
	printMatrix(magic3)
	fmt.Printf("¿Es mágica? %v\n\n", isMagicMatrix(magic3))
	
	// Ejemplo 5: Matriz 1x1 (trivialmente mágica)
	magic4 := [][]int{
		{42},
	}
	
	fmt.Println("Matriz 5:")
	printMatrix(magic4)
	fmt.Printf("¿Es mágica? %v\n\n", isMagicMatrix(magic4))
}
