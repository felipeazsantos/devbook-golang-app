package rotas

import (
	"devbook-golang-app/webapp/src/controllers"
	"net/http"
)

var rotaPaginaPrincipal = Rota{
	URI: "/home",
	Metodo: http.MethodGet,
	Funcao: controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
