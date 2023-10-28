package modelos

import (
	"time"
)

// ProdutoVendido representa um produto que faz parte de um pedido
type ProdutoVendido struct {
	NomeProduto  string  `json:"nomeProduto"`
	Quantidade   uint64  `json:"quantidade"`
	ValorProduto float64 `json:"valorProduto"`
	ValorTaxa    float64 `json:"valortaxa"`
}

// PedidoCompleto representa um pedido feito contendo uma lista de todos os produtos vendidos
type PedidoCompleto struct {
	ID                    uint64           `json:"id"`
	DataDaCompra          time.Time        `json:"dataDaCompra"`
	ListaProdutosVendidos []ProdutoVendido `json:"listaProdutosVendidos"`
	EmbalagemID           uint64           `json:"embalagemId"`
	PlataformaID          uint64           `json:"plataformaId"`
}
