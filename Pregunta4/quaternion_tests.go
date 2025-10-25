package quaternion

import (
	"math"
	"testing"
)

// TestQuaternionCreation prueba la creación de cuaterniones
func TestQuaternionCreation(t *testing.T) {
	q := New(1.0, 2.0, 3.0, 4.0)
	
	if q.A != 1.0 {
		t.Errorf("Expected a=1.0, got %f", q.A)
	}
	if q.B != 2.0 {
		t.Errorf("Expected b=2.0, got %f", q.B)
	}
	if q.C != 3.0 {
		t.Errorf("Expected c=3.0, got %f", q.C)
	}
	if q.D != 4.0 {
		t.Errorf("Expected d=4.0, got %f", q.D)
	}
}

// TestQuaternionZero prueba el cuaternión cero
func TestQuaternionZero(t *testing.T) {
	q := New(0, 0, 0, 0)
	
	if q.A != 0 || q.B != 0 || q.C != 0 || q.D != 0 {
		t.Errorf("Expected zero quaternion, got %v", q)
	}
}

// TestAdditionBasic prueba la suma básica de cuaterniones
func TestAdditionBasic(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	
	result := q1.Add(q2)
	
	if result.A != 6 || result.B != 8 || result.C != 10 || result.D != 12 {
		t.Errorf("Expected (6+8i+10j+12k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestAdditionWithZero prueba la suma con el cuaternión cero
func TestAdditionWithZero(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	qZero := New(0, 0, 0, 0)
	
	result := q1.Add(qZero)
	
	if result.A != q1.A || result.B != q1.B || result.C != q1.C || result.D != q1.D {
		t.Errorf("Addition with zero should not change quaternion")
	}
}

// TestAdditionCommutative verifica que la suma es conmutativa
func TestAdditionCommutative(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	
	r1 := q1.Add(q2)
	r2 := q2.Add(q1)
	
	if !almostEqual(r1, r2) {
		t.Errorf("Addition should be commutative")
	}
}

// TestAdditionWithScalar prueba la suma con un escalar (entero/flotante)
func TestAdditionWithScalar(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.AddScalar(3)
	
	if result.A != 4 || result.B != 2 || result.C != 3 || result.D != 4 {
		t.Errorf("Expected (4+2i+3j+4k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestAdditionWithFloat prueba la suma con un flotante
func TestAdditionWithFloat(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.AddScalar(2.5)
	
	if !floatEqual(result.A, 3.5) || result.B != 2 || result.C != 3 || result.D != 4 {
		t.Errorf("Expected (3.5+2i+3j+4k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestAdditionNegative prueba la suma con valores negativos
func TestAdditionNegative(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	q2 := New(-1, -2, -3, -4)
	
	result := q1.Add(q2)
	
	if result.A != 0 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("Expected zero quaternion, got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestAdditionChain prueba sumas encadenadas
func TestAdditionChain(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	q3 := New(1, 1, 1, 1)
	
	result := q1.Add(q2).Add(q3)
	
	if result.A != 7 || result.B != 9 || result.C != 11 || result.D != 13 {
		t.Errorf("Expected (7+9i+11j+13k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestConjugateBasic prueba el conjugado básico
func TestConjugateBasic(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.Conjugate()
	
	if result.A != 1 || result.B != -2 || result.C != -3 || result.D != -4 {
		t.Errorf("Expected (1-2i-3j-4k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestConjugateZero prueba el conjugado del cuaternión cero
func TestConjugateZero(t *testing.T) {
	q := New(0, 0, 0, 0)
	
	result := q.Conjugate()
	
	if result.A != 0 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("Conjugate of zero should be zero")
	}
}

// TestConjugateReal prueba el conjugado de un cuaternión real
func TestConjugateReal(t *testing.T) {
	q := New(5, 0, 0, 0)
	
	result := q.Conjugate()
	
	if result.A != 5 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("Conjugate of real quaternion should be itself")
	}
}

// TestDoubleConjugate prueba que el doble conjugado es el original
func TestDoubleConjugate(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.Conjugate().Conjugate()
	
	if !almostEqual(q, result) {
		t.Errorf("Double conjugate should equal original")
	}
}

// TestConjugateNegative prueba el conjugado con valores negativos
func TestConjugateNegative(t *testing.T) {
	q := New(-1, -2, -3, -4)
	
	result := q.Conjugate()
	
	if result.A != -1 || result.B != 2 || result.C != 3 || result.D != 4 {
		t.Errorf("Expected (-1+2i+3j+4k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductBasic prueba el producto básico
func TestProductBasic(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	
	result := q1.Multiply(q2)
	
	// Cálculo manual: (1+2i+3j+4k) * (5+6i+7j+8k)
	// a = 1*5 - 2*6 - 3*7 - 4*8 = 5 - 12 - 21 - 32 = -60
	// b = 1*6 + 2*5 + 3*8 - 4*7 = 6 + 10 + 24 - 28 = 12
	// c = 1*7 - 2*8 + 3*5 + 4*6 = 7 - 16 + 15 + 24 = 30
	// d = 1*8 + 2*7 - 3*6 + 4*5 = 8 + 14 - 18 + 20 = 24
	
	if result.A != -60 || result.B != 12 || result.C != 30 || result.D != 24 {
		t.Errorf("Expected (-60+12i+30j+24k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductWithUnit prueba el producto con el cuaternión unitario
func TestProductWithUnit(t *testing.T) {
	q := New(1, 2, 3, 4)
	unit := New(1, 0, 0, 0)
	
	result := q.Multiply(unit)
	
	if !almostEqual(q, result) {
		t.Errorf("Product with unit quaternion should not change quaternion")
	}
}

// TestProductISquared prueba que i^2 = -1
func TestProductISquared(t *testing.T) {
	qi := New(0, 1, 0, 0)
	
	result := qi.Multiply(qi)
	
	if result.A != -1 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("i^2 should equal -1, got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductJSquared prueba que j^2 = -1
func TestProductJSquared(t *testing.T) {
	qj := New(0, 0, 1, 0)
	
	result := qj.Multiply(qj)
	
	if result.A != -1 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("j^2 should equal -1, got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductKSquared prueba que k^2 = -1
func TestProductKSquared(t *testing.T) {
	qk := New(0, 0, 0, 1)
	
	result := qk.Multiply(qk)
	
	if result.A != -1 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("k^2 should equal -1, got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductIJK prueba que ijk = -1
func TestProductIJK(t *testing.T) {
	qi := New(0, 1, 0, 0)
	qj := New(0, 0, 1, 0)
	qk := New(0, 0, 0, 1)
	
	result := qi.Multiply(qj).Multiply(qk)
	
	if result.A != -1 || result.B != 0 || result.C != 0 || result.D != 0 {
		t.Errorf("ijk should equal -1, got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductNonCommutative verifica que el producto NO es conmutativo
func TestProductNonCommutative(t *testing.T) {
	qi := New(0, 1, 0, 0)
	qj := New(0, 0, 1, 0)
	
	r1 := qi.Multiply(qj)  // ij = k
	r2 := qj.Multiply(qi)  // ji = -k
	
	if almostEqual(r1, r2) {
		t.Errorf("Product should NOT be commutative")
	}
}

// TestProductWithScalar prueba el producto con un escalar
func TestProductWithScalar(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.MultiplyScalar(3)
	
	if result.A != 3 || result.B != 6 || result.C != 9 || result.D != 12 {
		t.Errorf("Expected (3+6i+9j+12k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductWithFloat prueba el producto con un flotante
func TestProductWithFloat(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.MultiplyScalar(2.5)
	
	if !floatEqual(result.A, 2.5) || !floatEqual(result.B, 5.0) || 
	   !floatEqual(result.C, 7.5) || !floatEqual(result.D, 10.0) {
		t.Errorf("Expected (2.5+5i+7.5j+10k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestProductAssociative verifica que el producto es asociativo
func TestProductAssociative(t *testing.T) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	q3 := New(2, 3, 4, 5)
	
	r1 := q1.Multiply(q2).Multiply(q3)
	r2 := q1.Multiply(q2.Multiply(q3))
	
	if !almostEqual(r1, r2) {
		t.Errorf("Product should be associative")
	}
}

// TestMagnitudeBasic prueba el cálculo de la magnitud
func TestMagnitudeBasic(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	result := q.Magnitude()
	expected := math.Sqrt(1*1 + 2*2 + 3*3 + 4*4) // sqrt(30)
	
	if !floatEqual(result, expected) {
		t.Errorf("Expected magnitude %f, got %f", expected, result)
	}
}

// TestMagnitudeZero prueba la magnitud del cuaternión cero
func TestMagnitudeZero(t *testing.T) {
	q := New(0, 0, 0, 0)
	
	result := q.Magnitude()
	
	if result != 0 {
		t.Errorf("Magnitude of zero quaternion should be 0, got %f", result)
	}
}

// TestMagnitudeUnit prueba la magnitud de cuaterniones unitarios
func TestMagnitudeUnit(t *testing.T) {
	tests := []Quaternion{
		New(1, 0, 0, 0),
		New(0, 1, 0, 0),
		New(0, 0, 1, 0),
		New(0, 0, 0, 1),
	}
	
	for _, q := range tests {
		result := q.Magnitude()
		if !floatEqual(result, 1.0) {
			t.Errorf("Magnitude of unit quaternion should be 1, got %f", result)
		}
	}
}

// TestMagnitudePositive verifica que la magnitud siempre es positiva
func TestMagnitudePositive(t *testing.T) {
	q := New(-1, -2, -3, -4)
	
	result := q.Magnitude()
	
	if result < 0 {
		t.Errorf("Magnitude should always be positive, got %f", result)
	}
}

// TestMagnitudeProperty verifica que |q*conj(q)| = |q|^2
func TestMagnitudeProperty(t *testing.T) {
	q := New(1, 2, 3, 4)
	
	product := q.Multiply(q.Conjugate())
	magnitude := q.Magnitude()
	
	if !floatEqual(product.A, magnitude*magnitude) {
		t.Errorf("q * conj(q) should equal |q|^2")
	}
}

// TestComplexExpression prueba expresiones complejas
func TestComplexExpression(t *testing.T) {
	a := New(1, 2, 3, 4)
	b := New(5, 6, 7, 8)
	c := New(2, 1, 0, -1)
	
	// (b + b) * (c + ~a)
	result := b.Add(b).Multiply(c.Add(a.Conjugate()))
	
	// Verificamos que no hay errores y el resultado tiene valores razonables
	if math.IsNaN(result.A) || math.IsNaN(result.B) || 
	   math.IsNaN(result.C) || math.IsNaN(result.D) {
		t.Errorf("Complex expression resulted in NaN")
	}
}

// TestMixedOperations prueba operaciones mixtas
func TestMixedOperations(t *testing.T) {
	a := New(1, 2, 3, 4)
	b := New(5, 6, 7, 8)
	
	// a * 3.0 + 7.0
	result := a.MultiplyScalar(3.0).AddScalar(7.0)
	
	if !floatEqual(result.A, 10.0) || result.B != 6 || result.C != 9 || result.D != 12 {
		t.Errorf("Expected (10+6i+9j+12k), got (%f+%fi+%fj+%fk)", 
			result.A, result.B, result.C, result.D)
	}
}

// TestOperationWithMagnitude prueba operaciones con la magnitud
func TestOperationWithMagnitude(t *testing.T) {
	b := New(1, 2, 3, 4)
	c := New(5, 6, 7, 8)
	
	// (b + b) * |c|
	mag := c.Magnitude()
	result := b.Add(b).MultiplyScalar(mag)
	
	if math.IsNaN(result.A) || math.IsInf(result.A, 0) {
		t.Errorf("Operation with magnitude resulted in NaN or Inf")
	}
}

// Funciones auxiliares

func floatEqual(a, b float64) bool {
	const epsilon = 1e-9
	return math.Abs(a-b) < epsilon
}

func almostEqual(q1, q2 Quaternion) bool {
	return floatEqual(q1.A, q2.A) && 
	       floatEqual(q1.B, q2.B) && 
	       floatEqual(q1.C, q2.C) && 
	       floatEqual(q1.D, q2.D)
}

// Benchmarks

func BenchmarkAddition(b *testing.B) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	
	for i := 0; i < b.N; i++ {
		_ = q1.Add(q2)
	}
}

func BenchmarkMultiplication(b *testing.B) {
	q1 := New(1, 2, 3, 4)
	q2 := New(5, 6, 7, 8)
	
	for i := 0; i < b.N; i++ {
		_ = q1.Multiply(q2)
	}
}

func BenchmarkConjugate(b *testing.B) {
	q := New(1, 2, 3, 4)
	
	for i := 0; i < b.N; i++ {
		_ = q.Conjugate()
	}
}

func BenchmarkMagnitude(b *testing.B) {
	q := New(1, 2, 3, 4)
	
	for i := 0; i < b.N; i++ {
		_ = q.Magnitude()
	}
}
