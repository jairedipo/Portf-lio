package main

import "fmt"

// A variável funcao gravará o contexto atual (x=10) e levará consigo essa informação
func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}
	return funcao
}

func main() {
	x := 20
	fmt.Println(x)

	imprimeX := closure()
	imprimeX() // Aqui ela irá imprimir 10 devido ao contexto em que foi criada
}
