package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0: // Caso inválido
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0: // Caso base
		return 1, nil
	default: // Chamada recursiva
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
