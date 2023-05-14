package main

import "fmt"

// trab1 bool, trab2 bool == trab1, trab2 bool
func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // && -> e
	comprarTv32 := trab1 != trab2    // != -> ou exclusivo
	comprarSorvete := trab1 || trab2 // || -> ou

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("Tv50 = %t, Tv32 = %t, Sorvete: %t, Saldável: = %t\n",
		tv50, tv32, sorvete, !sorvete) // ! -> negação

	x := 1
	y := 1

	x++ // incremento -> x + 1
	y-- // decremento -> y - 1

	fmt.Printf("x=%d e y=%d", x, y)

	// Go não possui operador ternário
}
