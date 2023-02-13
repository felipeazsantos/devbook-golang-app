package rotas

import (
	"github.com/gorilla/mux"
	"net/http"
)

//Rota representa todas as rotas da aplicação web
type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configrar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
