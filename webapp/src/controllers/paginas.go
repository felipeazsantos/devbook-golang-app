package controllers

import (
	"devbook-golang-app/webapp/src/utils"
	"net/http"
)

//CarregarTelaDeLogin vai carregar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

//CarregarPaginaDeCadastrarUsuario vai carregar a tela de cadastro de usuário
func CarregarPaginaDeCadastrarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}