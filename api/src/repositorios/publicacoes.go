package repositorios

import (
	"DevBook/api/src/modelos"
	"database/sql"
)

//Publicações representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicaoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db: db}
}

//Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

//BuscarPorID traz uma única publicação
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		`select p,id, p.titulo, p.conteudo, p.autor_id, u.nick, p.curtidas, p.CriadoEm from usuarios
				inner join publicacoes p on u.id = p.author_id 
			    where p.id = ?`, publicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro = linha.Scan(&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

//Buscar traz as publicações dos usuários seguidos e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
    select distinct p.*, u.nick from publicacoes p 
    inner join usuarios u on u.id = p.author_id
    inner join seguidores s on p.author_id = s.usuario_id
    where p.author_id = ? or s.seguidor_id = ?
	order by 1 desc
   `, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick
			); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}