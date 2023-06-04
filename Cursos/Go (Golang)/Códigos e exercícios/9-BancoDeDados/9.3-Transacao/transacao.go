package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "curso:#curso123@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin() // Inicia a transação
	smtm, _ := tx.Prepare("insert into usuarios(id, nome) values(?,?)")

	smtm.Exec(2000, "Bia")
	smtm.Exec(2001, "Carlos")
	_, err = smtm.Exec(1, "Tiago") // Erro de chave duplicada

	if err != nil {
		tx.Rollback() // Desfaz tudo que foi feito a partir do comando Begin
		log.Fatal(err)
	}

	tx.Commit() // Caso não faça rollback, esse comando efetiva as alterações realizadas
}
