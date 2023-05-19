package main

import "fmt"

// Nomeando os parâmetros de saída
func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return // Retorno limpo == 'return segundo, primeiro'
}

func main() {
	x, y := trocar(3, 7)
	fmt.Println(x, y)
}
