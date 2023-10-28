package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPedidos = []Rota{
	{
		URI:                "/pedidos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPedido,
		RequerAutenticacao: true,
	},
	{
		URI:                "/pedidos/{pedidoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPedido,
		RequerAutenticacao: true,
	},
	{
		URI:                "/pedidos/buscar-id",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/pedidos/listar-pedidos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarPedidosSemFiltro,
		RequerAutenticacao: true,
	},
}
