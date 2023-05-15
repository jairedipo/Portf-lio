package main

import (
	"fmt"
	"time"
)

func main() {

	n := 1
	for n <= 20 { // não tem while em Go, usar o for
		fmt.Printf("%d ", n)
		n++
	}

	fmt.Print("\nNúmeros pares: ")
	for n := 1; n <= 20; n++ {
		if n%2 == 0 {
			fmt.Printf("%d ", n)
		}
	}

	for { // laço infinito
		fmt.Print("\nPara sempre...")
		time.Sleep(time.Second)
	}
}
