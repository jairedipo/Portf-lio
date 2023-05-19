package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

// Recebe como parâmetro uma função e dois ints e retorna um int
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 3, 4)

	fmt.Println(resultado)
}
