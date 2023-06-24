package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	// Mapeia os arquivos hmtls criando templates
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"João", "joao.pedro@gmail.com"}

		// Utiliza um dos templates, sendo u um 'parâmetro enviado' ao html
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
