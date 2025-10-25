package quaternion

import (
	"fmt"
	"math"
)

// Quaternion representa un cuaternión de la forma a + bi + cj + dk
type Quaternion struct {
	A, B, C, D float64
}

// New crea un nuevo cuaternión con los coeficientes dados
func New(a, b, c, d float64) Quaternion {
	return Quaternion{A: a, B: b, C: c, D: d}
}

// FromReal crea un cuaternión a partir de un número real
func FromReal(a float64) Quaternion {
	return Quaternion{A: a, B: 0, C: 0, D: 0}
}

// String devuelve una representación en string del cuaternión
func (q Quaternion) String() string {
	return fmt.Sprintf("%.2f + %.2fi + %.2fj + %.2fk", q.A, q.B, q.C, q.D)
}

// Add suma dos cuaterniones
// (a₁ + b₁i + c₁j + d₁k) + (a₂ + b₂i + c₂j + d₂k) = 
// (a₁ + a₂) + (b₁ + b₂)i + (c₁ + c₂)j + (d₁ + d₂)k
func (q Quaternion) Add(other Quaternion) Quaternion {
	return Quaternion{
		A: q.A + other.A,
		B: q.B + other.B,
		C: q.C + other.C,
		D: q.D + other.D,
	}
}

// AddReal suma un número real a un cuaternión
// (a + bi + cj + dk) + n = (a + n) + bi + cj + dk
func (q Quaternion) AddReal(n float64) Quaternion {
	return Quaternion{
		A: q.A + n,
		B: q.B,
		C: q.C,
		D: q.D,
	}
}

// Conjugate calcula la conjugada del cuaternión
// ~(a + bi + cj + dk) = a - bi - cj - dk
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{
		A: q.A,
		B: -q.B,
		C: -q.C,
		D: -q.D,
	}
}

// Multiply multiplica dos cuaterniones
// La multiplicación de cuaterniones no es conmutativa
func (q Quaternion) Multiply(other Quaternion) Quaternion {
	// (a₁ + b₁i + c₁j + d₁k) * (a₂ + b₂i + c₂j + d₂k)
	a1, b1, c1, d1 := q.A, q.B, q.C, q.D
	a2, b2, c2, d2 := other.A, other.B, other.C, other.D
	
	return Quaternion{
		A: a1*a2 - b1*b2 - c1*c2 - d1*d2,
		B: a1*b2 + b1*a2 + c1*d2 - d1*c2,
		C: a1*c2 - b1*d2 + c1*a2 + d1*b2,
		D: a1*d2 + b1*c2 - c1*b2 + d1*a2,
	}
}

// MultiplyReal multiplica un cuaternión por un número real
// (a + bi + cj + dk) * n = (a*n) + (b*n)i + (c*n)j + (d*n)k
func (q Quaternion) MultiplyReal(n float64) Quaternion {
	return Quaternion{
		A: q.A * n,
		B: q.B * n,
		C: q.C * n,
		D: q.D * n,
	}
}

// Abs calcula la medida o valor absoluto del cuaternión
// &(a + bi + cj + dk) = √(a² + b² + c² + d²)
func (q Quaternion) Abs() float64 {
	return math.Sqrt(q.A*q.A + q.B*q.B + q.C*q.C + q.D*q.D)
}

// Equals compara dos cuaterniones con una tolerancia pequeña
func (q Quaternion) Equals(other Quaternion) bool {
	const epsilon = 1e-10
	return math.Abs(q.A-other.A) < epsilon &&
		math.Abs(q.B-other.B) < epsilon &&
		math.Abs(q.C-other.C) < epsilon &&
		math.Abs(q.D-other.D) < epsilon
}

// Operadores adicionales para conveniencia

// Subtract resta dos cuaterniones
func (q Quaternion) Subtract(other Quaternion) Quaternion {
	return Quaternion{
		A: q.A - other.A,
		B: q.B - other.B,
		C: q.C - other.C,
		D: q.D - other.D,
	}
}

// Negate niega un cuaternión (multiplica por -1)
func (q Quaternion) Negate() Quaternion {
	return Quaternion{
		A: -q.A,
		B: -q.B,
		C: -q.C,
		D: -q.D,
	}
}
