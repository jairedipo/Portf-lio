package modelos

import (
	"time"
)

// Busca representa atributos que podem ser usados na busca por pedidos cadastrados
type Busca struct {
	PedidoID     uint64    `json:"pedidoId,omitempty"`
	PlataformaID uint64    `json:"plataformaId,omitempty"`
	NomeProduto  string    `json:"nomeProduto,omitempty"`
	DataDaCompra time.Time `json:"dataDaCompra,omitempty"`
}
