package modelos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Pedido representa um pedido feito
type Pedido struct {
	ID           uint64    `json:"id,omitempty"`
	DataDaCompra time.Time `json:"dataDaCompra,omitempty"`
	ProdutoID    string    `json:"produtoId,omitempty"`
	Quantidade   uint64    `json:"quantidade,omitempty"`
	ValorProduto float64   `json:"valorProduto,omitempty"`
	ValorTaxa    float64   `json:"valortaxa,omitempty"`
	EstoqueID    uint64    `json:"estoqueId,omitempty"`
	EmbalagemID  uint64    `json:"embalagemId,omitempty"`
	PlataformaID uint64    `json:"plataformaId,omitempty"`
}

// Preparar valida e transforma os dados de um pedido feito
func (pedido *Pedido) Preparar(r *http.Request) error {
	id, erro := strconv.ParseUint(r.FormValue("id_pedido"), 10, 64)
	if erro != nil {
		return erro
	}

	dataDaCompra, erro := time.Parse("2006-01-02", r.FormValue("dataDaCompra"))
	if erro != nil {
		return erro
	}

	produto := r.FormValue("id_produto")
	produtoId := BuscarProdutoID(produto, r)
	if erro != nil {
		return erro
	}

	quantidade, erro := strconv.ParseUint(r.FormValue("quantidade"), 10, 64)
	if erro != nil {
		return erro
	}

	valorProduto, erro := strconv.ParseFloat(r.FormValue("valor"), 64)
	if erro != nil {
		return erro
	}

	valorTaxa, erro := strconv.ParseFloat(r.FormValue("valor_taxa"), 64)
	if erro != nil {
		return erro
	}

	estoqueId, erro := strconv.ParseUint(r.FormValue("estoque"), 10, 64)
	if erro != nil {
		return erro
	}

	embalagemId, erro := strconv.ParseUint(r.FormValue("embalagem"), 10, 64)
	if erro != nil {
		return erro
	}

	plataformaId, erro := strconv.ParseUint(r.FormValue("plataforma"), 10, 64)
	if erro != nil {
		return erro
	}

	pedido.ID = id
	pedido.DataDaCompra = dataDaCompra
	pedido.ProdutoID = produtoId
	pedido.Quantidade = quantidade
	pedido.ValorProduto = valorProduto
	pedido.ValorTaxa = valorTaxa
	pedido.EstoqueID = estoqueId
	pedido.EmbalagemID = embalagemId
	pedido.PlataformaID = plataformaId
	return nil
}

// BuscarProdutoID busca o ID de um produto a partir do nome informado
func BuscarProdutoID(produto string, r *http.Request) string {
	url := fmt.Sprintf("%s/buscar-id-produto/%s", config.APIURL, produto)
	response, erro := requisicoes.FazerRequisicao(r, http.MethodGet, url, nil)
	if erro != nil {
		return ""
	}
	defer response.Body.Close()

	var id_produto string
	if erro = json.NewDecoder(response.Body).Decode(&id_produto); erro != nil {
		return ""
	}
	return id_produto
}
