insert into produtos(id_produto, categoria, nome) values
('JF001', 'Material Escolar', 'Borracha'),
('JF002', 'Material Escolar', 'Carimbos One Piece');

select * from produtos;

insert into precos(id_produto, valor) values
('JF001', 7.00),
('JF001', 12.50),
('JF002', 20.00);

select * from precos;

insert into estoque(id_produto, quantidade_inicial, quantidade_disponivel, custo_unitario) values
('JF001', 15, 15, 1.80),
('JF001', 20, 20, 3.60),
('JF002', 28, 18, 12.37);

select * from estoque;

insert into embalagens(tamanho, quantidade_inicial, quantidade_utilizada, custo_unitario) values
('Saco Plástico P', 250, 20, 0.17),
('Saco Plástico M', 150, 10, 0.28),
('Saco Plástico G', 50, 2, 0.46);

select * from embalagens;

insert into plataformas_de_venda(nome) values
('Shopee'),
('ML'), 
('Outros');

select * from plataformas_de_venda;

insert into pedidos values
(40, '20231001', 'JF001', 1, 30.50, 10.10, 1, 1, 1),
(1, '20230909', 'JF002', 2, 30.00, 10.00, 1, 1, 1),
(2, '20230909', 'JF002', 1, 15.00, 5.00, 1, 1, 1);

select * from pedidos;

select pe.id_Pedido, pe.data_compra, pe.quantidade, pe.valor_produto, pe.valor_taxa, 
pe.id_Embalagem, pe.id_Plataforma, prod.nome from pedidos pe 
inner join produtos prod on pe.id_produto = prod.id_produto where
pe.data_compra > '20230901' and pe.data_compra < '20231001'
order by pe.id_pedido

select max(id_Pedido) + 1 as maiorId from pedidos

drop table precos;
drop table pedidos;
drop table estoque;
drop table embalagens;
drop table produtos;