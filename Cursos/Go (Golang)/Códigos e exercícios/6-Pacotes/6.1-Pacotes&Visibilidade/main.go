package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2)) // É visível, pois está dentro do mesmo pacote
	fmt.Println(Distancia(p1, p2))
}
