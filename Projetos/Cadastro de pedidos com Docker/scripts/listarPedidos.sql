SELECT 
id_pedido as "#Pedido",
data_compra as "Data",
produto.nome as "produto",
pedido.quantidade,
(preco_venda * pedido.quantidade) as "valor",
ROUND(preco_venda * pedido.quantidade * 0.2, 2) as "taxa",
ROUND((preco_venda * pedido.quantidade) - ROUND(preco_venda * pedido.quantidade * 0.2, 2), 2) as "receita",
(estoque.custo_unitario * pedido.quantidade) as "custo",
(embalagem.custo_unitario) as "embalagem",
ROUND((ROUND((preco_venda * pedido.quantidade) - ROUND(preco_venda * pedido.quantidade * 0.2, 2), 2)) - (estoque.custo_unitario * pedido.quantidade) - (embalagem.custo_unitario), 2) as "lucro"

FROM pedido 
    INNER JOIN produto ON pedido.id_Produto = produto.id_Produto
    INNER JOIN embalagem ON pedido.id_Embalagem = embalagem.id_Embalagem
    INNER JOIN estoque ON produto.id_Produto = estoque.id_Produto