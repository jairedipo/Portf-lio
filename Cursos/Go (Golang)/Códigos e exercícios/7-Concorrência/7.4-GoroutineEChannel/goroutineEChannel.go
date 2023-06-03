package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre goroutines

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) // Cria um assinc, pois chama a função e continua executando a main
	fmt.Println("Aguardando...")

	a, b := <-c, <-c // sincroniza novamente, aguardando o retorno do preenchimento do canal
	fmt.Println(a, b)

	fmt.Println(<-c)

	// Erro de deadlock, tentar ler mais uma vez, mas não ter mais goroutines em execução
	// fmt.Println(<-c)
}
