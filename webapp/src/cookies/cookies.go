package cookies

import (
	"devbook-golang-app/webapp/src/config"
	"github.com/gorilla/securecookie"
	"net/http"
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