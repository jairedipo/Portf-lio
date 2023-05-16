package main

import "fmt"

func main() {
	var notas [3]float64 // vetor de 3 posições do tipo float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)

	// Outra forma de criar e usar o for:
	numeros := [...]int{8, 3, 1, 4, 6}
	for i, numero := range numeros { // Semelhante a um foreach
		fmt.Printf("numeros[%d] = %d\n", i, numero)
	}

	for _, num := range numeros { // Nesse caso o índice é ignorado
		fmt.Println(num)
	}
}
