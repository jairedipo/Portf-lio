CREATE DATABASE IF NOT EXISTS jf;
USE jf;

create table produtos (
    id_produto varchar(6) PRIMARY KEY,
    nome varchar(100) not null,
    categoria varchar(100) not null
);

create table precos (
    id_Preco INT AUTO_INCREMENT PRIMARY KEY,
    id_produto varchar(6) not null,
    valor decimal(18,2) not null,
    FOREIGN KEY (id_produto) REFERENCES produtos (id_produto)
);

create table estoque (
	id_Estoque INT AUTO_INCREMENT PRIMARY KEY,
    id_produto varchar(6) not null,
    quantidade_inicial integer not null,
    quantidade_disponivel integer not null,
    custo_unitario decimal(18,2) not null,
    FOREIGN KEY (id_produto) REFERENCES produtos (id_produto)
);

create table embalagens (
    id_Embalagem INT AUTO_INCREMENT PRIMARY KEY,
    tamanho varchar(50) not null, -- Dimensoes AxLxC
    quantidade_inicial integer not null,
    quantidade_utilizada integer not null,
    custo_unitario decimal(18,2) not null
);

create table plataformas_de_venda (
	id_plataforma INT AUTO_INCREMENT PRIMARY KEY,
    nome varchar(100) not null
);

create table pedidos (
    id_Pedido int not null,
    data_compra date not null,
    id_produto varchar(6) not null,
    quantidade integer not null,
    valor_produto decimal(18,2) not null,
    valor_taxa decimal(18,2) not null,
    id_Estoque int not null,
    id_Embalagem int not null,
    id_Plataforma int not null,
    FOREIGN KEY (id_produto) REFERENCES produtos (id_produto),
    FOREIGN KEY (id_Estoque) REFERENCES estoque (id_Estoque),
    FOREIGN KEY (id_Embalagem) REFERENCES embalagens (id_Embalagem),
    FOREIGN KEY (id_Plataforma) REFERENCES plataformas_de_venda (id_Plataforma),
	CONSTRAINT pk_PP PRIMARY KEY(id_pedido, id_produto, id_estoque) 
);
