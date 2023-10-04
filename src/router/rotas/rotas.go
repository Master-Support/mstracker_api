package rotas

import "net/http"

// Rota Representa todas as rotas da API
type Rota struct {
	URI string 
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool 
}