package modelos

import (
	"devbook-golang-app/webapp/src/config"
	"devbook-golang-app/webapp/src/requisicoes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

//Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
	Nick string `json:"nick"`
	CriadoEm time.Time `json:"criado_em"`
	Seguidores []Usuario `json:"seguidores"`
	Seguindo []Usuario `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoes(canalPublicacoes, usuarioID, r)

	var (
		usuario Usuario
		seguidores []Usuario
		seguindo []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
			case usuarioCarregado := <-canalUsuario:
				if usuarioCarregado.ID == 0 {
					return Usuario{}, errors.New("Erro ao buscar o usuário")
				}

				usuario = usuarioCarregado
			case seguidoresCarregadores := <-canalSeguidores:
				if seguidoresCarregadores == nil {
					return Usuario{}, errors.New("Erro ao buscar os seguidores")
				}

				seguidores = seguidoresCarregadores
			case seguindoCarregados := <- canalSeguindo:
				if seguindoCarregados == nil {
					return Usuario{}, errors.New("Erro ao buscar quem o usuário está seguindo")
				}

				seguindo = seguindoCarregados
			case publicacoesCarregadas := <-canalPublicacoes:
				if publicacoesCarregadas == nil {
					return Usuario{}, errors.New("Erro ao buscar as publicações do usuário")
				}

				publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

//BuscarDadosDoUsuario chama a API para buscar os dados base do usuário
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

//BuscarSeguidores chama a API para buscar os seguidores do usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	canal <- seguidores
}

//BuscarSeguindo chama a API para buscar os usuários seguidos por um usuário
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	canal <- seguindo
}

// BuscarPublicacoes chama a API para buscar as publicações de um usuário
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIUrl, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	canal <- publicacoes
}

