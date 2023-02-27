package requisicoes

import (
	"devbook-golang-app/webapp/src/cookies"
	"fmt"
	"io"
	"net/http"
)

//FazerRequisicaoComAutenticacao é utilizado para colocar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer " + cookie["token"])
	fmt.Println(cookie["token"])

	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, erro
}
