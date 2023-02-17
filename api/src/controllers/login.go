package controllers

import (
	"devbook-golang-app/api/autenticacao"
	"devbook-golang-app/api/seguranca"
	"devbook-golang-app/api/src/banco"
	"devbook-golang-app/api/src/modelos"
	"devbook-golang-app/api/src/repositorios"
	"devbook-golang-app/api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Login é responsável por autenticar um usuáiro na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisição, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisição, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))
}
