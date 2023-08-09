INSERT INTO usuarios(nome, nick, email, senha) VALUES
("Thiago Alves", "Thi", "thiaguito@gmail.com", "$2a$10$kOPDvOqC.hgXKjF1Eb72/.kPlbdNtlwCgZ/eWPm/B5csuMA5dMr4q"),
("Sandy Pires", "Sandy", "sandy.pires@gmail.com", "$2a$10$kOPDvOqC.hgXKjF1Eb72/.kPlbdNtlwCgZ/eWPm/B5csuMA5dMr4q"),
("Terry Brot", "Terry", "terry.brot@outlook.com", "$2a$10$kOPDvOqC.hgXKjF1Eb72/.kPlbdNtlwCgZ/eWPm/B5csuMA5dMr4q");

INSERT INTO seguidores(usuario_id, seguidor_id) VALUES
(1, 2),
(1, 3),
(3, 1);

INSERT INTO publicacoes(titulo, conteudo, autor_id) VALUES
("Amo tirar férias!", "Estou curtindo muito essa pousada na praia!", 1),
("Dia especial", "Dia de ir as compraaaas! Amo demais!", 2),
("Com fome", "Aguardando a entrega da pizza que já está 30 min atrasada :(", 3);