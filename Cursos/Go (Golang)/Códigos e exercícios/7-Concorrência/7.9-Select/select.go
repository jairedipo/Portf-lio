package main

import (
	"fmt"
	"time"

	"github.com/cod3rcursos/html"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	// Estrutura de controle específica para concorrência
	// Executa de acordo com a informação que chegar primeiro
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"

		// Nesse caso, o default é retornado caso nenhuma das outras respostas estejam disponiveis
		// logo quando o select é executado
		// default:
		//     return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.youtube.com",
		"https://make.supercell.com/",
		"https://www.google.com",
	)
	fmt.Println(campeao)
}
