package modelos

import (
	"errors"
	"strings"
	"time"
)

//Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID uint64 `json:"id,omitempty"`
	Nome string `json:"nome,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

// Preparar valida e formata as informações do usuário
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O Nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O E-mail é obrigatório e não pode estar em branco")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
			return errors.New("A Senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}