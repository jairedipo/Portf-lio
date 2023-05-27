package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id_produto"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1) // Método que faz a conversão
	fmt.Println(string(p1Json))

	// Json para struct
	var p2 produto
	jsonString := `{"id_produto":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2) // Método que faz a conversão
	fmt.Println(p2.Tags[1])
	fmt.Println(p2)
}
