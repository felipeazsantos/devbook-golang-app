insert into usuarios(nome, nick, email, senha)
values
('Usuário 1', 'usuario_1', 'usuario@gmail.com', '$2a$10$wtBM6SWUzuqz.kHKdNBrP.IQZNqLU4Ye8eb80nOi0ypkFWq0c376q'), -- usuário 1
('Usuário 2', 'usuario_2', 'usuario2@gmail.com', '$2a$10$wtBM6SWUzuqz.kHKdNBrP.IQZNqLU4Ye8eb80nOi0ypkFWq0c376q'), -- usuário 2
('Usuário 3', 'usuario_3', 'usuario3@gmail.com', '$2a$10$wtBM6SWUzuqz.kHKdNBrP.IQZNqLU4Ye8eb80nOi0ypkFWq0c376q'); -- usuário 3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publicacoes(titulo, conteudo, autor_id)
values
("publicação do Usuário 1", "Essa é a publicação do usuário 1", 1),
("publicação do Usuário 2", "Essa é a publicação do usuário 2", 2),
("publicação do Usuário 3", "Essa é a publicação do usuário 3", 3);