package controllers

import (
	"bytes"
	"devbook-golang-app/webapp/src/config"
	"devbook-golang-app/webapp/src/requisicoes"
	"devbook-golang-app/webapp/src/respostas"
	"encoding/json"
	"fmt"
	"net/http"
)

//CriarPublicacao chama a API para cadastrar uma publicação no banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacao, erro := json.Marshal(map[string]string{
		"titulo": r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes", config.APIUrl)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
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

	respostas.JSON(w, response.StatusCode, nil)
}
