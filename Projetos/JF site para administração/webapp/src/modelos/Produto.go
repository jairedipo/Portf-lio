package modelos

// Produto representa um produto cadastrado
type Produto struct {
	ProdutoID string `json:"produtoId,omitempty"`
	Nome      string `json:"nome,omitempty"`
	Categoria string `json:"categoria,omitempty"`
}
