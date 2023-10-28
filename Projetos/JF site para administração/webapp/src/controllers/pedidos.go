package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// InserirPedido chama a api para inserir um pedido na base
func InserirPedido(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	modeloPedido := modelos.Pedido{}

	erro := modeloPedido.Preparar(r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	pedido, erro := json.Marshal(modeloPedido)

	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/pedidos", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(pedido))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarPedido chama a api para deletar um pedido que foi cadastrado na base
func DeletarPedido(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	pedidoID, erro := strconv.ParseUint(parametros["pedidoId"], 10, 64)

	url := fmt.Sprintf("%s/pedidos/%d", config.APIURL, pedidoID)
	response, erro := requisicoes.FazerRequisicao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
