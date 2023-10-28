package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarPedido adiciona um novo pedido no banco de dados
func CriarPedido(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var pedido modelos.Pedido
	if erro = json.Unmarshal(corpoRequisicao, &pedido); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePedidos(db)
	pedido.ID, erro = repositorio.Criar(pedido)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, pedido)
}

// DeletarPedido exclui um pedido do banco de dados
func DeletarPedido(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	pedidoID, erro := strconv.ParseUint(parametros["pedidoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePedidos(db)
	if erro = repositorio.Deletar(pedidoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarId busca o id do próximo pedido a ser cadastrado
func BuscarId(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePedidos(db)
	maiorId, erro := repositorio.BuscarId()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, maiorId)
}

// ListarPedidosSemFiltro retorna todos os pedidos cadastrados agrupando os itens por pedido
func ListarPedidosSemFiltro(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePedidos(db)
	pedidos, erro := repositorio.ListarPedidosSemFiltro()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	for i, pedido := range pedidos {
		pedidos[i], erro = repositorio.BuscarProdutosPedidoCompleto(pedido)
		if erro != nil {
			respostas.Erro(w, http.StatusInternalServerError, erro)
			return
		}
	}

	respostas.JSON(w, http.StatusOK, pedidos)
}
