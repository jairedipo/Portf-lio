package main

import (
	"fmt"
	m "math" // Renomear um import
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// Forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// Criando constantes e variáveis em lote
	const (
		a = "a"
		b = "b"
	)

	var (
		c = "c"
		d = "d"
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false // Criar mais de uma variável em 1 linha
	fmt.Println(e, f)

	g, h, i := 2, false, "Epa!" // Criar várias variáveis de tipos diferentes
	fmt.Println(g, h, i)
}
