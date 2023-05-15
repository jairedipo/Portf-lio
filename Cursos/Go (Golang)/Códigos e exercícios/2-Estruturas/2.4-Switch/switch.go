package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough // Execute o comando do próximo case
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "F"
	default:
		return "Nota inválida"
	}
}

func periodoDoDia() {
	t := time.Now()
	switch { // 'switch true' com condicionais
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}

func tipo(i interface{}) string { // interface{} permite receber diferentes tipos
	switch i.(type) { // switch baseado no tipo da variável
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(notaParaConceito(10.0))
	fmt.Println(notaParaConceito(8.5))
	fmt.Println(notaParaConceito(2.4))
	fmt.Println(notaParaConceito(5.3))
	fmt.Println(notaParaConceito(-1.0))
	fmt.Println(notaParaConceito(11.5))

	fmt.Println()

	periodoDoDia()

	fmt.Println()

	fmt.Println(tipo(2.3))
	fmt.Println(tipo(5))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
}
