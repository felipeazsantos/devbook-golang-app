package autenticacao

import (
	"DevBook/api/src/config"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
	"errors"
)

//CriarToken gera uma string de autenticação com permissões para determinado usuário
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{} // jwt - JSON Web Token
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId
	token := jwt.NewWithClaims(jwt.SigningMethodES256, permissoes)
	return token.SignedString(config.SecretKey) //secret
}

// ValidarToken verifica o token passado na requisição é valido
func ValidarToken(r *http.Request) error {
	tokenSring := extrairToken(r)
	token, erro := jwt.Parse(tokenSring, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}