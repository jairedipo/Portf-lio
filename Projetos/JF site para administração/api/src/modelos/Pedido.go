package modelos

import (
	"time"
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
