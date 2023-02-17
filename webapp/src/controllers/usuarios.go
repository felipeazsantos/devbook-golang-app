package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(usuario))
}
