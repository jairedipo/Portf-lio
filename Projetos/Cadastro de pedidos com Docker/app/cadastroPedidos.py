import psycopg2
from bottle import route, run, request

DSN = 'dbname=jf_brinquedos user=postgres password=teste123 host=db'
SQL = 'INSERT INTO pedido (data_compra, id_Produto, quantidade, valor_compra, id_embalagem) VALUES (%s, %s, %s, %s, %s)'

def registrar_Pedido(dataCompra, sku, quantidade, embalagem):
    conn = psycopg2.connect(DSN)
    cur = conn.cursor()
    print('parametros: {} {} {} {} {}'.format(dataCompra, sku, quantidade, 0, embalagem))
    cur.execute(SQL, (dataCompra, sku, quantidade, 0, embalagem))
    conn.commit()
    cur.close()
    conn.close()

    print('Pedido gravado!')

@route('/', method='POST')
def send():
    dataCompra = request.forms.get('dataCompra')
    embalagem = request.forms.get('embalagem')
    sku = request.forms.get('SKU')
    quantidade = request.forms.get('quantidade') 
    registrar_Pedido(dataCompra, sku, quantidade, embalagem)   
    return 'Pedido cadastrado! Data: {}, Produto: {}, Quantidade: {}, Embalagem: {}'.format(
        dataCompra, sku, quantidade, embalagem
    )

if __name__ == '__main__':
    run(host='0.0.0.0', port=8080, debug=True)