insert into usuarios(nome, nick, email, senha)
values
('Usuário 1', 'usuario_1', 'usuario@gmail.com', '123456'), -- usuário 1
('Usuário 2', 'usuario_2', 'usuario2@gmail.com', '123456'), -- usuário 2
('Usuário 3', 'usuario_3', 'usuario3@gmail.com', '123456'); -- usuário 3

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);