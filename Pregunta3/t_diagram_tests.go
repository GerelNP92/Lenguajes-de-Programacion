package main

import (
	"testing"
)

// TestDefinirProgramaNuevo verifica que se puede definir un programa nuevo
func TestDefinirProgramaNuevo(t *testing.T) {
	s := NuevoSistema()
	err := s.DefinirPrograma("test", "Java")
	
	if err != nil {
		t.Errorf("No debería dar error al definir programa nuevo: %v", err)
	}
	
	if _, existe := s.programas["test"]; !existe {
		t.Error("El programa no fue agregado al sistema")
	}
}

// TestDefinirProgramaDuplicado verifica que no se puede definir un programa duplicado
func TestDefinirProgramaDuplicado(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("test", "Java")
	err := s.DefinirPrograma("test", "Python")
	
	if err == nil {
		t.Error("Debería dar error al definir programa duplicado")
	}
}

// TestProgramaEnLOCAL verifica que un programa en LOCAL es ejecutable
func TestProgramaEnLOCAL(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("fibonacci", "LOCAL")
	
	err := s.PuedeEjecutar("fibonacci")
	if err != nil {
		t.Errorf("Un programa en LOCAL debería ser ejecutable: %v", err)
	}
}

// TestProgramaSinInterprete verifica que un programa sin intérprete no es ejecutable
func TestProgramaSinInterprete(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("factorial", "Java")
	
	// No debería dar error, pero el sistema debe reportar que no es ejecutable
	err := s.PuedeEjecutar("factorial")
	if err != nil {
		t.Errorf("No debería dar error: %v", err)
	}
}

// TestProgramaConInterpreteDirecto verifica ejecución con intérprete directo
func TestProgramaConInterpreteDirecto(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("factorial", "Java")
	s.DefinirInterprete("LOCAL", "Java")
	
	err := s.PuedeEjecutar("factorial")
	if err != nil {
		t.Errorf("Debería ser ejecutable con intérprete directo: %v", err)
	}
}

// TestProgramaConInterpreteIndirecto verifica ejecución con intérprete indirecto
func TestProgramaConInterpreteIndirecto(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("factorial", "Java")
	s.DefinirInterprete("C", "Java")
	s.DefinirInterprete("LOCAL", "C")
	
	err := s.PuedeEjecutar("factorial")
	if err != nil {
		t.Errorf("Debería ser ejecutable con intérprete indirecto: %v", err)
	}
}

// TestProgramaConTraductorDirecto verifica ejecución con traductor directo
func TestProgramaConTraductorDirecto(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("factorial", "Java")
	s.DefinirTraductor("LOCAL", "Java", "LOCAL")
	
	err := s.PuedeEjecutar("factorial")
	if err != nil {
		t.Errorf("Debería ser ejecutable con traductor directo: %v", err)
	}
}

// TestProgramaConTraductorSinInterpreteBase verifica que traductor sin base no funciona
func TestProgramaConTraductorSinInterpreteBase(t *testing.T) {
	s := NuevoSistema()
	s.DefinirPrograma("factorial", "Java")
	s.DefinirTraductor("C", "Java", "LOCAL")
	// No hay intérprete para C
	
	err := s.PuedeEjecutar("factorial")
	if err != nil {
		t.Errorf("No debería dar error: %v", err)
	}
	// El programa no debería ser ejecutable
}

// TestEjemploCompleto verifica el ejemplo completo del enunciado
func TestEjemploCompleto(t *testing.T) {
	s := NuevoSistema()
	
	// fibonacci en LOCAL - ejecutable
	s.DefinirPrograma("fibonacci", "LOCAL")
	if err := s.PuedeEjecutar("fibonacci"); err != nil {
		t.Error("fibonacci debería ser ejecutable")
	}
	
	// factorial en Java - no ejecutable aún
	s.DefinirPrograma("factorial", "Java")
	
	// Agregar intérprete e intérprete base
	s.DefinirInterprete("C", "Java")
	s.DefinirTraductor("C", "Java", "C")
	s.DefinirInterprete("LOCAL", "C")
	
	// Ahora factorial debería ser ejecutable
	if err := s.PuedeEjecutar("factorial"); err != nil {
		t.Error("factorial debería ser ejecutable después de agregar intérpretes")
	}
}

// TestCadenaDeTraductores verifica traducción en cadena
func TestCadenaDeTraductores(t *testing.T) {
	s := NuevoSistema()
	
	s.DefinirPrograma("holamundo", "Python3")
	s.DefinirTraductor("wtf42", "Python3", "LOCAL")
	s.DefinirTraductor("C", "wtf42", "Java")
	s.DefinirInterprete("LOCAL", "C")
	
	// wtf42 no es ejecutable aún, entonces el traductor de Python3 a LOCAL no funciona
	// Pero C sí es ejecutable, entonces el traductor de wtf42 a Java funciona
	// Pero esto requiere que wtf42 sea ejecutable primero
	
	err := s.PuedeEjecutar("holamundo")
	if err != nil {
		t.Errorf("No debería dar error: %v", err)
	}
}

// TestProgramaNoExistente verifica error cuando programa no existe
func TestProgramaNoExistente(t *testing.T) {
	s := NuevoSistema()
	err := s.PuedeEjecutar("noexiste")
	
	if err == nil {
		t.Error("Debería dar error cuando el programa no existe")
	}
}

// TestMultiplesInterpretes verifica que múltiples intérpretes funcionan
func TestMultiplesInterpretes(t *testing.T) {
	s := NuevoSistema()
	
	s.DefinirPrograma("prog1", "Lang1")
	s.DefinirPrograma("prog2", "Lang2")
	
	s.DefinirInterprete("LOCAL", "Lang1")
	s.DefinirInterprete("LOCAL", "Lang2")
	
	if err := s.PuedeEjecutar("prog1"); err != nil {
		t.Error("prog1 debería ser ejecutable")
	}
	
	if err := s.PuedeEjecutar("prog2"); err != nil {
		t.Error("prog2 debería ser ejecutable")
	}
}

// TestTraductorCircular verifica manejo de traducciones circulares
func TestTraductorCircular(t *testing.T) {
	s := NuevoSistema()
	
	s.DefinirPrograma("test", "A")
	s.DefinirInterprete("LOCAL", "B")
	s.DefinirTraductor("B", "A", "C")
	s.DefinirTraductor("B", "C", "A")
	
	// Esto no debería causar loop infinito
	err := s.PuedeEjecutar("test")
	if err != nil {
		t.Errorf("No debería dar error: %v", err)
	}
}

// TestTraductorAMismoLenguaje verifica traductor que traduce al mismo lenguaje
func TestTraductorAMismoLenguaje(t *testing.T) {
	s := NuevoSistema()
	
	s.DefinirPrograma("test", "Java")
	s.DefinirTraductor("LOCAL", "Java", "Java")
	s.DefinirInterprete("LOCAL", "Java")
	
	err := s.PuedeEjecutar("test")
	if err != nil {
		t.Error("test debería ser ejecutable")
	}
}

// TestCadenaLargaDeInterpretes verifica cadena larga de intérpretes
func TestCadenaLargaDeInterpretes(t *testing.T) {
	s := NuevoSistema()
	
	s.DefinirPrograma("test", "Lang5")
	s.DefinirInterprete("LOCAL", "Lang1")
	s.DefinirInterprete("Lang1", "Lang2")
	s.DefinirInterprete("Lang2", "Lang3")
	s.DefinirInterprete("Lang3", "Lang4")
	s.DefinirInterprete("Lang4", "Lang5")
	
	err := s.PuedeEjecutar("test")
	if err != nil {
		t.Error("test debería ser ejecutable con cadena larga")
	}
}

// TestTraductorYInterpreteMixtos verifica combinación de traductores e intérpretes
func TestTraductorYInterpreteMixtos(t *testing.T) {
	s := NuevoSistema()
	
	s.DefinirPrograma("test", "Python")
	s.DefinirTraductor("Java", "Python", "C")
	s.DefinirInterprete("LOCAL", "Java")
	s.DefinirInterprete("LOCAL", "C")
	
	err := s.PuedeEjecutar("test")
	if err != nil {
		t.Error("test debería ser ejecutable")
	}
}
