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
	{
		URI: "/usuarios",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/buscar-usuarios",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPaginaDeUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuario/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
}
