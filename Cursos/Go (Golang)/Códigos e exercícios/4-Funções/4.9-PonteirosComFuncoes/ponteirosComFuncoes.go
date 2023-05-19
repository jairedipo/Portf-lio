package main

import "fmt"

// Recebe uma cópia do parâmetro e não altera o original - função pura
func inc1(n int) {
	n++
}

// Recebe o endereço do parâmetro original, permitindo a sua alteração - função impura
func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n) // Por valor

	inc2(&n)
	fmt.Println(n) // Por referência
}
