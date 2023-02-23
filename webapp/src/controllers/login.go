package controllers

import (
	"bytes"
	"devbook-golang-app/webapp/src/respostas"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	token, erro := ioutil.ReadAll(response.Body)

	fmt.Println(string(token), erro)
}
