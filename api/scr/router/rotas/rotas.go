package rotas

import (
	"api/scr/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas do API
type Rota struct {
	URI                string
	Metodo             string
	funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.funcao))).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, middlewares.Logger(rota.funcao)).Methods(rota.Metodo)
		}
	}

	return r
}
