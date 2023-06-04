package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "curso:#curso123@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// criando um statement (template) para insert
	smtm, _ := db.Prepare("insert into usuarios(nome) values(?)")
	smtm.Exec("Maria")
	smtm.Exec("João")

	// Capturando a resposta da execução do comando
	res, _ := smtm.Exec("Pedro")

	id, _ := res.LastInsertId() // Retorna o id do elemento inserido
	fmt.Println("ID:", id)

	linhas, _ := res.RowsAffected() // Retorna o número de linhas afetadas
	fmt.Println("Linhas afetadas:", linhas)
}
