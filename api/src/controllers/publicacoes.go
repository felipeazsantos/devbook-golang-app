package controllers

import "net/http"

//CriarPublicao permite que um usuário crie uma publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar Publicação"))
}

//BuscarPublicacoes retorna todas as publicações do feed do usuário
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar Publicações"))
}

//BuscarPublicacao retorna uma única publicação
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar Publicação"))
}

//AtualizarPublicao permite que um usuário atualize uma publicação sua
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Publicação"))
}

//DeletarPublicao permite que um usuário delete uma publicação sua
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Publicação"))
}