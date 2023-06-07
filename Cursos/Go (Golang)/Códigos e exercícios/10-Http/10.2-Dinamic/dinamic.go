package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	/* Formatação de data em Go:
	02 = Dia
	01 = Mês
	2006 = Ano com 4 dígitos (ou 06 para 2 dígitos)
	03 = Hora (ou 15 para formato 24h)
	04 = Minuto
	05 = Segundo
	*/
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
