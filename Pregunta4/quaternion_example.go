package main

import (
	"fmt"
	"quaternion"
)

func main() {
	fmt.Println("=== Ejemplos de Operaciones con Cuaterniones ===\n")
	
	// Crear cuaterniones
	q1 := quaternion.New(1, 2, 3, 4)
	q2 := quaternion.New(2, 3, 4, 5)
	q3 := quaternion.New(1, 0, 1, 0)
	
	fmt.Println("Cuaterniones definidos:")
	fmt.Printf("q1 = %s\n", q1)
	fmt.Printf("q2 = %s\n", q2)
	fmt.Printf("q3 = %s\n\n", q3)
	
	// Suma de cuaterniones: q1 + q2
	fmt.Println("1. Suma de cuaterniones (q1 + q2):")
	suma := q1.Add(q2)
	fmt.Printf("   %s\n\n", suma)
	
	// Conjugada: ~q1
	fmt.Println("2. Conjugada (~q1):")
	conjugada := q1.Conjugate()
	fmt.Printf("   %s\n\n", conjugada)
	
	// Producto de cuaterniones: q1 * q2
	fmt.Println("3. Producto de cuaterniones (q1 * q2):")
	producto := q1.Multiply(q2)
	fmt.Printf("   %s\n\n", producto)
	
	// Medida o valor absoluto: &q1
	fmt.Println("4. Medida o valor absoluto (&q1):")
	medida := q1.Abs()
	fmt.Printf("   %.4f\n\n", medida)
	
	// Expresiones compuestas
	fmt.Println("=== Expresiones Compuestas ===\n")
	
	// (q1 + q2) * q3
	fmt.Println("5. (q1 + q2) * q3:")
	expr1 := q1.Add(q2).Multiply(q3)
	fmt.Printf("   %s\n\n", expr1)
	
	// (q1 + q1) * (q3 + ~q2)
	fmt.Println("6. (q1 + q1) * (q3 + ~q2):")
	expr2 := q1.Add(q1).Multiply(q3.Add(q2.Conjugate()))
	fmt.Printf("   %s\n\n", expr2)
	
	// &(q3 * q1)
	fmt.Println("7. &(q3 * q1):")
	expr3 := q3.Multiply(q1).Abs()
	fmt.Printf("   %.4f\n\n", expr3)
	
	// Operaciones con números reales
	fmt.Println("=== Operaciones con Números Reales ===\n")
	
	// q1 + 3
	fmt.Println("8. q1 + 3:")
	suma_real := q1.AddReal(3)
	fmt.Printf("   %s\n\n", suma_real)
	
	// q2 * 2.5
	fmt.Println("9. q2 * 2.5:")
	prod_real := q2.MultiplyReal(2.5)
	fmt.Printf("   %s\n\n", prod_real)
	
	// q1 * 3.0 + 7.0
	fmt.Println("10. q1 * 3.0 + 7.0:")
	expr4 := q1.MultiplyReal(3.0).AddReal(7.0)
	fmt.Printf("    %s\n\n", expr4)
	
	// (q1 + q1) * &q3
	fmt.Println("11. (q1 + q1) * &q3:")
	abs_q3 := q3.Abs()
	expr5 := q1.Add(q1).MultiplyReal(abs_q3)
	fmt.Printf("    %s\n", expr5)
	fmt.Printf("    (donde &q3 = %.4f)\n\n", abs_q3)
	
	// Verificación de propiedades
	fmt.Println("=== Verificación de Propiedades ===\n")
	
	// Verificar i² = j² = k² = ijk = -1
	i := quaternion.New(0, 1, 0, 0)
	j := quaternion.New(0, 0, 1, 0)
	k := quaternion.New(0, 0, 0, 1)
	
	i2 := i.Multiply(i)
	j2 := j.Multiply(j)
	k2 := k.Multiply(k)
	ijk := i.Multiply(j).Multiply(k)
	
	fmt.Println("12. Verificando i² = j² = k² = ijk = -1:")
	fmt.Printf("    i² = %s\n", i2)
	fmt.Printf("    j² = %s\n", j2)
	fmt.Printf("    k² = %s\n", k2)
	fmt.Printf("    ijk = %s\n\n", ijk)
	
	// Conjugada de conjugada
	fmt.Println("13. Verificando ~~q1 = q1:")
	doble_conj := q1.Conjugate().Conjugate()
	fmt.Printf("    q1 = %s\n", q1)
	fmt.Printf("    ~~q1 = %s\n", doble_conj)
	fmt.Printf("    ¿Son iguales? %v\n\n", q1.Equals(doble_conj))
	
	// |q| = |~q|
	fmt.Println("14. Verificando |q1| = |~q1|:")
	abs_q1 := q1.Abs()
	abs_conj_q1 := q1.Conjugate().Abs()
	fmt.Printf("    |q1| = %.4f\n", abs_q1)
	fmt.Printf("    |~q1| = %.4f\n", abs_conj_q1)
	fmt.Printf("    ¿Son iguales? %v\n", math.Abs(abs_q1-abs_conj_q1) < 1e-10)
}
