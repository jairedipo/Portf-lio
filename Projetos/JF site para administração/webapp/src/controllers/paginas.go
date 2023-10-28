package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

// CarregarPaginaCadastrarPedidos redenriza a página de cadastro de pedidos
func CarregarPaginaCadastrarPedidos(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/pedidos/buscar-id", config.APIURL)
	response, erro := requisicoes.FazerRequisicao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var pedido modelos.Pedido
	if erro = json.NewDecoder(response.Body).Decode(&pedido); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url = fmt.Sprintf("%s/buscar-produtos", config.APIURL)
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

	utils.ExecutarTemplate(w, "cadastrar-pedidos.html", struct {
		Produtos []string
		PedidoID modelos.Pedido
	}{
		Produtos: nomeProdutos,
		PedidoID: pedido,
	})
}

// CarregarPaginaListarPedidos redenriza a página que lista os pedidos cadastrados
func CarregarPaginaListarPedidos(w http.ResponseWriter, r *http.Request) {
	modeloBusca := modelos.Busca{
		PedidoID:     0,
		PlataformaID: 0,
		NomeProduto:  "",
	}

	busca, erro := json.Marshal(modeloBusca)

	url := fmt.Sprintf("%s/pedidos/listar-pedidos", config.APIURL)
	response, erro := requisicoes.FazerRequisicao(r, http.MethodGet, url, bytes.NewBuffer(busca))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var pedidos []modelos.Pedido
	if erro = json.NewDecoder(response.Body).Decode(&pedidos); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "listar-pedidos.html", pedidos)
}
