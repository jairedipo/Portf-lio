package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Pedido representa um repositório de publicações
type Pedido struct {
	db *sql.DB
}

// NovoRepositorioDePedido cria um repositório de publicações
func NovoRepositorioDePedidos(db *sql.DB) *Pedido {
	return &Pedido{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Pedido) Criar(pedido modelos.Pedido) (uint64, error) {
	linha, erro := repositorio.db.Prepare(
		`insert into pedidos (id_Pedido, data_compra, id_produto, quantidade, valor_produto, valor_taxa,
			id_Estoque, id_Embalagem, id_Plataforma) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer linha.Close()

	resultado, erro := linha.Exec(
		pedido.ID,
		pedido.DataDaCompra.String()[:10], // Enviar apenas o dia/mes/ano
		pedido.ProdutoID,
		pedido.Quantidade,
		pedido.ValorProduto,
		pedido.ValorTaxa,
		pedido.EstoqueID,
		pedido.EmbalagemID,
		pedido.PlataformaID,
	)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// Deletar exclui as informações de um pedido no banco de dados
func (repositorio Pedido) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from pedidos where id_Pedido = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarId busca o próximo id que deve ser utilizado ao cadastrar um pedido
func (repositorio Pedido) BuscarId() (modelos.Pedido, error) {
	linha, erro := repositorio.db.Query("select ifnull(max(id_Pedido) + 1, 1) as maiorId from pedidos")
	if erro != nil {
		return modelos.Pedido{}, erro
	}
	defer linha.Close()

	var maiorId modelos.Pedido

	linha.Next()
	if erro = linha.Scan(
		&maiorId.ID,
	); erro != nil {
		return modelos.Pedido{}, erro
	}

	return maiorId, nil
}

// ListarPedidosSemFiltro retorna uma lista com todos os pedidos cadastrados
func (repositorio Pedido) ListarPedidosSemFiltro() ([]modelos.PedidoCompleto, error) {
	linhas, erro := repositorio.db.Query(`select id_pedido, data_compra, id_embalagem, id_plataforma from pedidos
	group by id_pedido, data_compra, id_embalagem, id_plataforma
	order by id_pedido desc`)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var pedidos []modelos.PedidoCompleto

	for linhas.Next() {
		var pedido modelos.PedidoCompleto

		if erro = linhas.Scan(
			&pedido.ID,
			&pedido.DataDaCompra,
			&pedido.EmbalagemID,
			&pedido.PlataformaID,
		); erro != nil {
			return nil, erro
		}

		pedidos = append(pedidos, pedido)
	}

	return pedidos, nil
}

// BuscarProdutosPedidoCompleto preenche os produtos vendidos de um pedido completo
func (repositorio Pedido) BuscarProdutosPedidoCompleto(pedido modelos.PedidoCompleto) (modelos.PedidoCompleto, error) {
	linhas, erro := repositorio.db.Query(`select nome, quantidade, valor_produto, valor_taxa from pedidos
	inner join produtos on pedidos.id_produto = produtos.id_produto
	where id_pedido = ?`, pedido.ID)
	if erro != nil {
		return pedido, erro
	}
	defer linhas.Close()

	var produtos []modelos.ProdutoVendido

	for linhas.Next() {
		var produto modelos.ProdutoVendido

		if erro = linhas.Scan(
			&produto.NomeProduto,
			&produto.Quantidade,
			&produto.ValorProduto,
			&produto.ValorTaxa,
		); erro != nil {
			return pedido, erro
		}

		produtos = append(produtos, produto)
	}

	pedido.ListaProdutosVendidos = produtos

	return pedido, nil
}
