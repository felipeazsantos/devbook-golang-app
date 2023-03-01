package rotas

import (
	"devbook-golang-app/webapp/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI: "/publicacoes",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
}
