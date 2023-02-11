package controllers

import (
	"DevBook/api/autenticacao"
	"DevBook/api/src/banco"
	"DevBook/api/src/modelos"
	"DevBook/api/src/repositorios"
	"DevBook/api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CriarPublicao permite que um usuário crie uma publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoDaRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
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