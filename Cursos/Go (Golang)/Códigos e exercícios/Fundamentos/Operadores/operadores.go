package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	/* básico: soma (+), subtração (-), divisão (/)
	multiplicacão (*) e resto da divisão (%) */

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01

	// usando o import math
	c := 3.0
	d := 2.0

	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
	fmt.Println("Raiz quadrada =>", math.Sqrt(c))
	fmt.Println("Logaritmo natural =>", math.Log(10))
	fmt.Println("Logaritmo base 10 =>", math.Log10(10))

	// atribuições com operação
	// +=, -=, *=, /=, %=

	// trocar variáveis sem auxiliar
	a, b = b, a
	fmt.Println(a, b)

	// operadores relacionais
	// == ou .Equal(), !=, >, <, >=, <=
	// == serve até para structs
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Fulano"}
	p2 := Pessoa{"Fulano"}
	fmt.Println("Pessoas iguais:", p1 == p2)
}
