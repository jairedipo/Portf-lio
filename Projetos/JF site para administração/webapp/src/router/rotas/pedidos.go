package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPedidos = []Rota{
	{
		URI:                "/pedidos",
		Metodo:             http.MethodPost,
		Funcao:             controllers.InserirPedido,
		RequerAutenticacao: true,
	},
	{
		URI:                "/pedidos/{pedidoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPedido,
		RequerAutenticacao: true,
	},
}
