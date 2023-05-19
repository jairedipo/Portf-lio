package main

import "fmt"

var soma = func(a, b int) int { // Definição da função atribuida a variável soma
	return a + b
}

func main() {
	fmt.Println(soma(2, 3))

	sub := func(a, b int) int { // Definição da função atribuida a variável sub
		return a - b
	}

	fmt.Println(sub(2, 3))
}
