package rotas

import (
	"devbook-golang-app/webapp/src/controllers"
	"net/http"
)

var rotaLogout = Rota{
	URI: "/logout",
	Metodo: http.MethodGet,
	Funcao: controllers.FazerLogout,
	RequerAutenticacao: true,
}
