package rotas

import (
	"net/http"

	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da Aplicação Web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasPedidos
	rotas = append(rotas, rotasPaginas...)
	rotas = append(rotas, rotasProdutos...)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI,
			middlewares.Logger(rota.Funcao),
		).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
