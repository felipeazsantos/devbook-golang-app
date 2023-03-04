package cookies

import (
	"devbook-golang-app/webapp/src/config"
	"github.com/gorilla/securecookie"
	"net/http"
	"time"
)

var s *securecookie.SecureCookie

//Configurar utiliza as variáveis de ambiente para criar o SecureCookie
func Configurar() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

//Salvar registra as informações de login
func Salvar(w http.ResponseWriter, ID, Token string) error {
	dados := map[string]string{
		"id": ID,
		"token": Token,
	}

	dadosCodificados, erro := s.Encode("dados", dados)
	if erro != nil {
		return erro
	}

	http.SetCookie(w, &http.Cookie{
		Name: "dados",
		Value: dadosCodificados,
		Path: "/",
		HttpOnly: true,
	})

	return nil
}

// Ler retorna os valores armazenados em cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, erro := r.Cookie("dados")
	if erro != nil {
		return nil, erro
	}

	valores := make(map[string]string)
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil {
		return nil, erro
	}

	return valores, nil
}

//Deletar remove os valores de autenticação armazenados no cookie
func Deletar(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "dados",
		Value: "",
		Path: "/",
		HttpOnly: true,
		Expires: time.Unix(0, 0),
	})
}