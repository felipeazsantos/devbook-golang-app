package autenticacao

import (
	"DevBook/api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
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