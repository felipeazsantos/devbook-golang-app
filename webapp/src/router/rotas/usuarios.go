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
		URI: "/usuarios/{usuarioId}",
		Metodo: http.MethodGet,
		Funcao: controllers.CarregarPerfilDoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI: "/usuarios/{usuarioId}/seguir",
		Metodo: http.MethodPost,
		Funcao: controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
}
