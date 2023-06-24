package rotas

import "net/http"

// ROta representa todas as rotas da API
type Rota struct {
	URI    string
	Metodo string
	Funcao func(http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}