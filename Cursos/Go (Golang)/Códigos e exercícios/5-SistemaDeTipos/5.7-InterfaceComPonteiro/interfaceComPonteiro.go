package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

type fusca struct {
	turboLigado bool
}

// Para esse tipo de implementação, a função será impura, alterando o valor original
func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func (f fusca) ligarTurbo() {
	f.turboLigado = true
	fmt.Println("Fusca não possui turbo")
}

func main() {
	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()

	// Necessário usar '&', pois o método é implementado usando ponteiro
	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)

	var carro3 esportivo = fusca{false}
	carro3.ligarTurbo()

	fmt.Println(carro3)
}
