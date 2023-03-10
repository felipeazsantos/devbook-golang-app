package controllers

import (
	"devbook-golang-app/api/autenticacao"
	"devbook-golang-app/api/seguranca"
	"devbook-golang-app/api/src/banco"
	"devbook-golang-app/api/src/modelos"
	"devbook-golang-app/api/src/repositorios"
	"devbook-golang-app/api/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
		fmt.Println("Teste")
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Println("Teste")
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		fmt.Println("Teste")
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		fmt.Println("Teste")
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		fmt.Println("Teste")
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{usuarioID, token})

}
