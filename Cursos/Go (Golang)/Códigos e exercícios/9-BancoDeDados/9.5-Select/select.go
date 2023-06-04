package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "curso:#curso123@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuarios where id > ?", 3)
	defer rows.Close() // Necessário fechar esse retorno também

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome) // Faz a leitura e atribuição
		fmt.Println(u)
	}
}
