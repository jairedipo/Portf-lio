package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// A mesma função pode ser chamada com quantidade de parâmetros diferentes
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1, 5.9))
	fmt.Printf("Média: %.2f\n", media(7.7, 8.1, 5.9, 9.9))
	fmt.Printf("Média: %.2f\n", media(7.0))

	// Passando um slice como parâmetro
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)
}
