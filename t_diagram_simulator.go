package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Programa representa un programa escrito en algún lenguaje
type Programa struct {
	nombre   string
	lenguaje string
}

// Interprete representa un intérprete para un lenguaje
type Interprete struct {
	lenguajeBase     string // lenguaje en el que está escrito el intérprete
	lenguajeInterpretado string // lenguaje que interpreta
}

// Traductor representa un traductor de un lenguaje a otro
type Traductor struct {
	lenguajeBase   string // lenguaje en el que está escrito el traductor
	lenguajeOrigen string // lenguaje fuente
	lenguajeDestino string // lenguaje destino
}

// Sistema mantiene el estado del simulador
type Sistema struct {
	programas    map[string]Programa
	interpretes  []Interprete
	traductores  []Traductor
}

// NuevoSistema crea un nuevo sistema vacío
func NuevoSistema() *Sistema {
	return &Sistema{
		programas:   make(map[string]Programa),
		interpretes: make([]Interprete, 0),
		traductores: make([]Traductor, 0),
	}
}

// DefinirPrograma define un nuevo programa
func (s *Sistema) DefinirPrograma(nombre, lenguaje string) error {
	if _, existe := s.programas[nombre]; existe {
		return fmt.Errorf("ERROR: Ya existe un programa con el nombre '%s'", nombre)
	}
	s.programas[nombre] = Programa{nombre: nombre, lenguaje: lenguaje}
	fmt.Printf("Se definió el programa '%s', ejecutable en '%s'\n", nombre, lenguaje)
	return nil
}

// DefinirInterprete define un nuevo intérprete
func (s *Sistema) DefinirInterprete(lenguajeBase, lenguajeInterpretado string) {
	s.interpretes = append(s.interpretes, Interprete{
		lenguajeBase:     lenguajeBase,
		lenguajeInterpretado: lenguajeInterpretado,
	})
	fmt.Printf("Se definió un intérprete para '%s', escrito en '%s'\n", 
		lenguajeInterpretado, lenguajeBase)
}

// DefinirTraductor define un nuevo traductor
func (s *Sistema) DefinirTraductor(lenguajeBase, lenguajeOrigen, lenguajeDestino string) {
	s.traductores = append(s.traductores, Traductor{
		lenguajeBase:   lenguajeBase,
		lenguajeOrigen: lenguajeOrigen,
		lenguajeDestino: lenguajeDestino,
	})
	fmt.Printf("Se definió un traductor de '%s' hacia '%s', escrito en '%s'\n",
		lenguajeOrigen, lenguajeDestino, lenguajeBase)
}

// PuedeEjecutar verifica si un programa puede ejecutarse
func (s *Sistema) PuedeEjecutar(nombre string) error {
	programa, existe := s.programas[nombre]
	if !existe {
		return fmt.Errorf("ERROR: No existe un programa con el nombre '%s'", nombre)
	}
	
	// Usamos un conjunto de lenguajes ejecutables para búsqueda BFS
	ejecutables := make(map[string]bool)
	ejecutables["LOCAL"] = true
	
	// Iteramos hasta que no haya cambios (punto fijo)
	cambio := true
	for cambio {
		cambio = false
		
		// Agregar lenguajes que pueden interpretarse
		for _, interp := range s.interpretes {
			if ejecutables[interp.lenguajeBase] && !ejecutables[interp.lenguajeInterpretado] {
				ejecutables[interp.lenguajeInterpretado] = true
				cambio = true
			}
		}
		
		// Agregar lenguajes a los que podemos traducir
		for _, trad := range s.traductores {
			if ejecutables[trad.lenguajeBase] && ejecutables[trad.lenguajeOrigen] {
				if !ejecutables[trad.lenguajeDestino] {
					ejecutables[trad.lenguajeDestino] = true
					cambio = true
				}
			}
		}
	}
	
	if ejecutables[programa.lenguaje] {
		fmt.Printf("Si, es posible ejecutar el programa '%s'\n", nombre)
		return nil
	}
	
	fmt.Printf("No es posible ejecutar el programa '%s'\n", nombre)
	return nil
}

// ProcesarComando procesa un comando del usuario
func (s *Sistema) ProcesarComando(comando string) bool {
	partes := strings.Fields(comando)
	if len(partes) == 0 {
		return true
	}
	
	accion := strings.ToUpper(partes[0])
	
	switch accion {
	case "SALIR":
		return false
		
	case "DEFINIR":
		if len(partes) < 3 {
			fmt.Println("ERROR: Comando DEFINIR incompleto")
			return true
		}
		
		tipo := strings.ToUpper(partes[1])
		switch tipo {
		case "PROGRAMA":
			if len(partes) != 4 {
				fmt.Println("ERROR: DEFINIR PROGRAMA requiere <nombre> <lenguaje>")
				return true
			}
			if err := s.DefinirPrograma(partes[2], partes[3]); err != nil {
				fmt.Println(err)
			}
			
		case "INTERPRETE":
			if len(partes) != 4 {
				fmt.Println("ERROR: DEFINIR INTERPRETE requiere <lenguaje_base> <lenguaje>")
				return true
			}
			s.DefinirInterprete(partes[2], partes[3])
			
		case "TRADUCTOR":
			if len(partes) != 5 {
				fmt.Println("ERROR: DEFINIR TRADUCTOR requiere <lenguaje_base> <lenguaje_origen> <lenguaje_destino>")
				return true
			}
			s.DefinirTraductor(partes[2], partes[3], partes[4])
			
		default:
			fmt.Printf("ERROR: Tipo desconocido '%s'\n", tipo)
		}
		
	case "EJECUTABLE":
		if len(partes) != 2 {
			fmt.Println("ERROR: EJECUTABLE requiere <nombre>")
			return true
		}
		if err := s.PuedeEjecutar(partes[1]); err != nil {
			fmt.Println(err)
		}
		
	default:
		fmt.Printf("ERROR: Comando desconocido '%s'\n", accion)
	}
	
	return true
}

func main() {
	sistema := NuevoSistema()
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Simulador de Diagramas T")
	fmt.Println("Comandos disponibles:")
	fmt.Println("  DEFINIR PROGRAMA <nombre> <lenguaje>")
	fmt.Println("  DEFINIR INTERPRETE <lenguaje_base> <lenguaje>")
	fmt.Println("  DEFINIR TRADUCTOR <lenguaje_base> <lenguaje_origen> <lenguaje_destino>")
	fmt.Println("  EJECUTABLE <nombre>")
	fmt.Println("  SALIR")
	fmt.Println()
	
	for {
		fmt.Print("$> ")
		if !scanner.Scan() {
			break
		}
		
		comando := scanner.Text()
		if !sistema.ProcesarComando(comando) {
			break
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error leyendo entrada: %v\n", err)
	}
}