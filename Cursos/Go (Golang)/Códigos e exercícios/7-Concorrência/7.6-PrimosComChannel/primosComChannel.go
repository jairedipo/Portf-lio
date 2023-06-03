package main

import (
	"fmt"
	"time"
)

func isPrimo(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if isPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) // Fecha o canal
}

func main() {
	c := make(chan int, 30)
	go primos(cap(c), c)   // Cap -> retorna a capacidade do buffer do canal
	for primo := range c { // Usando o range do canal
		fmt.Printf("%d ", primo)
	}
	fmt.Println("Fim!")
}
