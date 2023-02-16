package rotas

import (
	"devbook-golang-app/webapp/src/controllers"
	"net/http"
)

var rotaUsuarios = []Rota{
	{
		URI: "/criar-usuario",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeCadastrarUsuario,
		RequerAutenticacao: false,
	},
}
