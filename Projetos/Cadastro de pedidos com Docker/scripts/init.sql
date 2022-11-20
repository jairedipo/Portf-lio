create database jf_brinquedos;

\c jf_brinquedos

create table produto (
    id_Produto varchar(6) CONSTRAINT id_Produto PRIMARY KEY,
    nome varchar(100) not null,
    preco_venda decimal not null
);

create table estoque (
    id_Produto varchar(6) not null,
    quantidade integer not null,
    custo_unitario decimal not null,
    FOREIGN KEY (id_Produto) REFERENCES produto (id_Produto)
);

create table embalagem (
    id_Embalagem varchar(6) CONSTRAINT id_Embalagem PRIMARY KEY,
    tamanho varchar(50) not null, -- Dimensoes AxLxC
    quantidade integer not null,
    custo_unitario decimal not null
);

create table pedido (
    id_pedido serial not null,
    data_compra date not null,
    id_Produto varchar(6) not null,
    quantidade integer not null,
    valor_compra decimal not null,
    id_Embalagem varchar(6) not null,
    FOREIGN KEY (id_Produto) REFERENCES produto (id_Produto),
    FOREIGN KEY (id_Embalagem) REFERENCES embalagem (id_Embalagem)
);

insert into produto (id_Produto, nome, preco_venda) values ('JF0001', 'Boneca Barbie', 100.00);
insert into produto (id_Produto, nome, preco_venda) values ('JF0002', 'Pelúcia Dora Aventureira', 39.99);

insert into embalagem (id_Embalagem, tamanho, quantidade, custo_unitario) values ('JFE001', '15x21', 100, 0.17);

insert into estoque (id_Produto, quantidade, custo_unitario) values ('JF0001', 10, 60.00);
insert into estoque (id_Produto, quantidade, custo_unitario) values ('JF0002', 10, 19.90);