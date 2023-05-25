package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // Sem identificador, vai gerar campos anonimos
	turboLigado bool
}

func main() {
	f := ferrari{}
	// A atribuição/leitura dos campos funciona normalmente, baseada no identificador do tipo original
	f.nome = "F40"
	f.velocidadeAtual = 190
	f.turboLigado = true

	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboLigado)
	fmt.Println(f) // Aqui podemos ver a composição
}
