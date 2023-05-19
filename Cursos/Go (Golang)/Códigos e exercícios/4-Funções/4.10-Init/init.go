package main

import "fmt"

// Terminal: go run init.go init2.go
// Função que executa antes da main e pode ser criada um para cada arquivo .go
func init() {
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Main...")
}
