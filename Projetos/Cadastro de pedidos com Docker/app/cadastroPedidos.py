from bottle import route, run, request

@route('/', method='POST')
def send():
    dataCompra = request.forms.get('dataCompra')
    embalagem = request.forms.get('embalagem')
    sku = request.forms.get('SKU')
    quantidade = request.forms.get('quantidade')    
    return 'Pedido cadastrado! Data: {}, Produto: {}, Quantidade: {}, Embalagem: {}'.format(
        dataCompra, sku, quantidade, embalagem
    )

if __name__ == '__main__':
    run(host='0.0.0.0', port=8080, debug=True)