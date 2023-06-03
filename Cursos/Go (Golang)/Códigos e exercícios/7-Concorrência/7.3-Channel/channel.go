package main

import "fmt"

func main() {
	// Tipo canal = chan
	ch := make(chan int, 1)

	ch <- 1 // Enviar dados para o canal (escrita)
	<-ch    // Receber dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
