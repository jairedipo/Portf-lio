package main

import "fmt"

func main() {
	n := 1

	var p *int = nil // p é um ponteiro de int
	p = &n           // atribuição do endereço de memória da variável
	*p++             // desreferenciando (pegando o valor da variável)
	n++

	fmt.Println(p, &n, *p, n)
}
