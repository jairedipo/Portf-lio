package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPaginas = []Rota{
	{
		URI:                "/pedidos/cadastrar-pedidos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaCadastrarPedidos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/pedidos/listar-pedidos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaListarPedidos,
		RequerAutenticacao: true,
	},
}
