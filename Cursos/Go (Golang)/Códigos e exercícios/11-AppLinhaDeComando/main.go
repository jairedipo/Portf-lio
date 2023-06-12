package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}

// Pare executar o arquivo .exe basta passar o parâmetro 'ip' ou 'servidores' e,
// opcionalmente, '--host {URL}'
