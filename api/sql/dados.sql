insert into usuarios (nome, nick, email, senha)
values 
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$IGT9rjCLfWawiRA1HHBQK.lB4OT3oR2upyxG8r9xHJxXvKN59PeUO" ), 
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$IGT9rjCLfWawiRA1HHBQK.lB4OT3oR2upyxG8r9xHJxXvKN59PeUO" ),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$IGT9rjCLfWawiRA1HHBQK.lB4OT3oR2upyxG8r9xHJxXvKN59PeUO" );

insert into seguidores(usuario_id, seguidor_id)
values 
(1,2),
(3,1),
(1,3);


insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacao do usuário 1", "Essa é a publicacao do usuário 1! Oba", 1),
("Publicacao do usuário 2", "Essa é a publicacao do usuário 2! Oba", 2),
("Publicacao do usuário 3", "Essa é a publicacao do usuário 3! Oba", 3);
