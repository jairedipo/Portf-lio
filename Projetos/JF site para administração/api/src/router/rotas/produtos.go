package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasProdutos = []Rota{
	{
		URI:                "/buscar-produtos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarProdutos,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscar-id-produto/{nome}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarIDProduto,
		RequerAutenticacao: true,
	},
}
