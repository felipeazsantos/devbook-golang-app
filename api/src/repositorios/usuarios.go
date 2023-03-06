package repositorios

import (
	"devbook-golang-app/api/src/modelos"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuario cria um reposítorio de usuários
func NovoRepositorioDeUsuario(db *sql.DB) *usuarios {
	return &usuarios{db}
}

//Criar insere um usuário no banco de dados
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

// Buscar tráz todos os usuários que atendem um filtro de nome ou nick
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, CriadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuariosResult []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuariosResult = append(usuariosResult, usuario)
	}

	return usuariosResult, nil
}

// BuscarPorId tráz um usuário do banco de dados
func (repositorio usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, CriadoEm from usuarios where id = ?", id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm);
		   erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//Atualizar altera as informações de um usuário
func (repositorio usuarios) Atualizar(id uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id)
	if erro != nil {
		return erro
	}

	return nil
}

//Deletar exclui um usuário no banco de dados
func (repositorio usuarios) Deletar(id uint64) error {
	statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(id)
	if erro != nil {
		return erro
	}

	return nil
}

//BuscarPorEmail busca um usuário por email e o retorna o seu id e senha com hash
func (repositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro:= repositorio.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()


	var usuario modelos.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//Seguir permite que um usuário siga outro
func (repositorio usuarios) Seguir(usuarioId, SeguidorId uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores(usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, SeguidorId); erro != nil {
		return erro
	}

	return nil
}


//PararDeSeguir permite um usuário parar de seguir outro
func (repositorio usuarios) PararDeSeguir(usuarioId, SeguidorId uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuarioId, SeguidorId); erro != nil {
		return erro
	}

	return nil
}

//BuscarSeguidores retorna todos os seguidores de um usuário
func (repositorio usuarios) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
    	select u.id, u.nome, u.nick, u.email, u.CriadoEm from usuarios u
		inner join seguidores s on 
		u.id = s.seguidor_id
		where s.usuario_id = ?
    `, usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuariosResult []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuariosResult = append(usuariosResult, usuario)
	}

	return usuariosResult, nil
}

//BuscarSeguindo retorna todoso os usuários que um determinado usuário está seguindo
func (repositorio usuarios) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.CriadoEm from usuarios u
		inner join seguidores s on u.id = s.usuario_id
        where s.seguidor_id = ?
	`, usuarioId)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuariosResult []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuariosResult = append(usuariosResult, usuario)
	}

	return usuariosResult, nil
}

//BuscarSenha retorna a senha do usuário pelo id
func (repositorio usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioId)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if (linha.Next()) {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

//AtualizarSenha altera a senha de um usuário
func (repositorio usuarios) AtualizarSenha(usuarioId uint64, senhaComHash string) error {
	statement, erro := repositorio.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _,erro = statement.Exec(senhaComHash, usuarioId); erro != nil {
		return erro
	}

	return nil
}