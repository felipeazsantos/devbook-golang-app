package rotas

import (
	"devbook-golang-app/api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota {
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI: "/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoes,
		RequerAutenticacao: false,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodPut,
		Funcao: controllers.AtualizarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI: "/publicacoes/{publicacaoId}",
		Metodo: http.MethodDelete,
		Funcao: controllers.DeletarPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI: "/usuarios/{usuarioId}/publicacoes",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: false,
	},
	{
		URI: "/publicacoes/{publicacaoId}/curtir",
		Metodo: http.MethodPost,
		Funcao: controllers.CurtirPublicacao,
		RequerAutenticacao: false,
	},
	{
		URI: "/publicacoes/{publicacaoId}/descurtir",
		Metodo: http.MethodPost,
		Funcao: controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
