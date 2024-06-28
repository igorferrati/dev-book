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
