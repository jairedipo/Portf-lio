package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// // Modo normal
	// fale("Maria", "Pq vc não fala comigo?", 3)
	// fale("João", "Só posso falar depois de você!", 1)

	// // O comando go cria uma Goroutine, que permite a execução de forma concorrente,
	// // nesse caso, vai executar ambas as chamadas até o programa finalizar e não as 500 vezes
	// go fale("Maria", "Ei...", 500)
	// go fale("João", "Opa...", 500)
	// time.Sleep(time.Second * 5)

	go fale("Maria", "Entendi!", 10)
	fale("João", "Parabéns!", 5)
}
