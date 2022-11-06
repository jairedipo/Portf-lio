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
    data_compra timestamp not null,
    id_Produto varchar(6) not null,
    quantidade integer not null,
    valor_compra decimal not null,
    id_embalagem varchar(6) not null,
    FOREIGN KEY (id_Produto) REFERENCES produto (id_Produto),
    FOREIGN KEY (id_Embalagem) REFERENCES embalagem (id_Embalagem)
);