package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) // Gera um número "aleatório" usando o nanosegundo atual
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // Variável i definida apenas no escopo do if/else
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}

	// Aqui já não é mais possível acessá-la
}
