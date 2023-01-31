package repositorios

import (
	"DevBook/api/src/modelos"
	"database/sql"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um reposítorio de usuários
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Criar insere um usuário no banco de dados
func (u usuarios) Criar(usuario modelos.Usuario) (uint64 error) {
	return 0, nil
}
