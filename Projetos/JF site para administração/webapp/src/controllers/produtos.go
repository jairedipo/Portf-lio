package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
)

// BuscarProdutos busca o nome de todos os produtos cadastrados na base
func BuscarProdutos(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/buscar-produtos", config.APIURL)
	responseProduto, erro := requisicoes.FazerRequisicao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer responseProduto.Body.Close()

	if responseProduto.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, responseProduto)
		return
	}

	var produtos []modelos.Produto
	if erro = json.NewDecoder(responseProduto.Body).Decode(&produtos); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	var nomeProdutos []string

	for _, produto := range produtos {
		nomeProdutos = append(nomeProdutos, produto.Nome)
	}

	respostas.JSON(w, http.StatusOK, nomeProdutos)
}
