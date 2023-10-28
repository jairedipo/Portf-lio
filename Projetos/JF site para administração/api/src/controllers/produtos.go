package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"net/http"

	"github.com/gorilla/mux"
)

// BuscarProdutos busca todos os produtos cadastrados
func BuscarProdutos(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	produtos, erro := repositorio.BuscarProdutos()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, produtos)
}

// BuscarIDProduto busca o id de um produto usando o nome dele
func BuscarIDProduto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	nomeProduto := parametros["nome"]

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeProdutos(db)
	IDProduto, erro := repositorio.BuscarIDProduto(nomeProduto)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, IDProduto)
}
