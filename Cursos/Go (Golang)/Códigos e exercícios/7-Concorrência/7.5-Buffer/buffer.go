package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("Li o valor 1")
	ch <- 2
	fmt.Println("Li o valor 2")
	ch <- 3
	fmt.Println("Li o valor 3")
	ch <- 4
	fmt.Println("Li o valor 4") // Leitura feita até aqui
	ch <- 5
	fmt.Println("Li o valor 5")
	ch <- 6
	fmt.Println("Li o valor 6")
}

func main() {
	ch := make(chan int, 3) // Canal com buffer de tamanho 3
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	time.Sleep(time.Second)
	fmt.Println("Finalizando...")
}
