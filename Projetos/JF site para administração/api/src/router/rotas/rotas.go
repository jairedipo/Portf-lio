package rotas

import (
	"net/http"

	"api/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasPedidos
	rotas = append(rotas, rotasProdutos...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
	}

	return r
}
