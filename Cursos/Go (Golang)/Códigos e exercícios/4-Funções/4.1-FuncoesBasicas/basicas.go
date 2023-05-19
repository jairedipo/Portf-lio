package main

import "fmt"

func f1(s1, s2 string) {
	fmt.Printf("F1: %s %s\n", s1, s2)
}

func f2(s1 string) string {
	return fmt.Sprintf("F2: %s", s1)
}

func f3() (string, string) {
	return "retorno 1", "retorno 2"
}

func main() {
	f1("A", "B")

	r1, r2 := f2("C"), f2("D")
	fmt.Println(r1, r2)

	r3, r4 := f3() // Função que retorna mais de 1 parâmetro
	fmt.Println(r3, r4)
}
