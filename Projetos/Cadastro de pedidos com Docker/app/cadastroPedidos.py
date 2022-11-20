import psycopg2
import os
from bottle import Bottle, request

class Sender(Bottle):
    def __init__(self):
        super().__init__()
        self.route('/', method='POST', callback=self.cadastrar)
        db_host = os.getenv('DB_HOST', 'db')
        db_user = os.getenv('DB_USER', 'postgres')
        db_pass = os.getenv('DB_PASS', 'password')
        db_name = os.getenv('DB_NAME', 'sender')
        dsn = f'dbname={db_name} user={db_user} password={db_pass} host={db_host}'
        self.conn = psycopg2.connect(dsn)

    def registrar_Pedido(self, dataCompra, sku, quantidade, embalagem):
        SQL = 'INSERT INTO pedido (data_compra, id_Produto, quantidade, valor_compra, id_embalagem) VALUES (%s, %s, %s, %s, %s)'
        cur = self.conn.cursor()
        cur.execute(SQL, (dataCompra, sku, quantidade, 0, embalagem))
        self.conn.commit()
        cur.close()

        print('Pedido gravado!')

    def cadastrar(self):
        dataCompra = request.forms.get('dataCompra')
        embalagem = request.forms.get('embalagem')
        sku = request.forms.get('SKU')
        quantidade = request.forms.get('quantidade') 
        self.registrar_Pedido(dataCompra, sku, quantidade, embalagem)   
        return 'Pedido cadastrado! Data: {}, Produto: {}, Quantidade: {}, Embalagem: {}'.format(
            dataCompra, sku, quantidade, embalagem
        )

if __name__ == '__main__':
    sender = Sender()
    sender.run(host='0.0.0.0', port=8080, debug=True)