package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int -> float
	x := 2.4
	y := 2
	fmt.Println(x / float64(y)) // 1.2

	// float -> int (não arredonda)
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// tabela ascii
	fmt.Println("Teste " + string(97)) // 97 == 'a'

	// int -> string
	fmt.Println("Teste " + strconv.Itoa(123)) // I to a == int to string

	// string -> int
	num, _ := strconv.Atoi("123") // Retorna o valor convertido e um erro (ignorado pelo uso do '_')
	fmt.Println(num - 122)        // 123 - 122 = 1

	// string -> bool
	b, _ := strconv.ParseBool("true") // true: "1" ou "true", false: qualquer outra coisa
	if b {
		fmt.Println("Verdadeiro")
	}
}
