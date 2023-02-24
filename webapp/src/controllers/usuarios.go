package controllers

import (
	"bytes"
	"devbook-golang-app/webapp/src/config"
	"devbook-golang-app/webapp/src/respostas"
	"encoding/json"
	"fmt"
	"net/http"
)

//CriarUsuario chama a API para cadastrar um usuário no banco
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"nome": r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick": r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro : erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios", config.APIUrl)
	response,erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro : erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
