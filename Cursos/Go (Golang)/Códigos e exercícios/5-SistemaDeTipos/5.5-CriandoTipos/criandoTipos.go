package main

import "fmt"

type nota float64

// Usando o tipo nota ao invés do tipo float64
func (n nota) entre(inicio, fim nota) bool {
	return n >= inicio && n <= fim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 6.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "F"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
}
