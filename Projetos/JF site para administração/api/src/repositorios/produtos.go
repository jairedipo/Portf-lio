package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Produto representa um repositório de produtos
type Produto struct {
	db *sql.DB
}

// NovoRepositorioDeProduto cria um repositório de produtos
func NovoRepositorioDeProdutos(db *sql.DB) *Produto {
	return &Produto{db}
}

// BuscarProdutos busca o nome de todos os produtos cadastrados
func (repositorio Produto) BuscarProdutos() ([]modelos.Produto, error) {
	linhas, erro := repositorio.db.Query("select nome from produtos")
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var produtos []modelos.Produto

	for linhas.Next() {
		var produto modelos.Produto

		if erro = linhas.Scan(
			&produto.Nome,
		); erro != nil {
			return nil, erro
		}

		produtos = append(produtos, produto)
	}

	return produtos, nil
}

// BuscarIDProduto busca o id de um produto através do nome informado
func (repositorio Produto) BuscarIDProduto(nome string) (string, error) {
	linha, erro := repositorio.db.Query("select id_produto from produtos where nome = ?", nome)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var id_produto string

	linha.Next()
	if erro = linha.Scan(&id_produto); erro != nil {
		return "", erro
	}

	return id_produto, nil
}
