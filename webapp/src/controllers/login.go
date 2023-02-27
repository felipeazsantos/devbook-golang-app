package controllers

import (
	"bytes"
	"devbook-golang-app/webapp/src/config"
	"devbook-golang-app/webapp/src/cookies"
	"devbook-golang-app/webapp/src/modelos"
	"devbook-golang-app/webapp/src/respostas"
	"encoding/json"
	"fmt"
	"net/http"
)

// FazerLogin utiliza o email e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		fmt.Println(erro)
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/%s", config.APIUrl, "login")
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		fmt.Println(erro)
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		fmt.Println(response.StatusCode)
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		fmt.Println(erro)
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro = cookies.Salvar(w, dadosAutenticacao.ID, dadosAutenticacao.Token); erro != nil {
		fmt.Println(erro)
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
