package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasProdutos = []Rota{
	{
		URI:                "/buscar-produtos",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarProdutos,
		RequerAutenticacao: true,
	},
}
