package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import indireto
)

// Executa os comandos e para o código em caso de erro
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Conexão com o BD
	db, err := sql.Open("mysql", "curso:#curso123@/") // user:senha@/
	if err != nil {
		panic(err)
	}
	defer db.Close() // Encerrar a conexão ao terminar a main()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
