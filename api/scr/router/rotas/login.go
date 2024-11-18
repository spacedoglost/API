package rotas

import (
	"api/scr/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	funcao:             controllers.Login,
	RequerAutenticacao: false,
}
