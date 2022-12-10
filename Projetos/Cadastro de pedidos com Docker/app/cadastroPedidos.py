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

        self.route('/listarPedidos', method='GET', callback=self.listar)

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
        if (self.validar_SKU(sku) and self.validar_Embalagem(embalagem)):
            self.registrar_Pedido(dataCompra, sku, quantidade, embalagem)   
            return 'Pedido cadastrado! Data: {}, Produto: {}, Quantidade: {}, Embalagem: {}'.format(
                dataCompra, sku, quantidade, embalagem
            )
        elif (not self.validar_SKU(sku) and not self.validar_Embalagem(embalagem)):
            return f'SKU: {sku} não foi encontrado e a Embalagem: {embalagem} não está cadastrada.'  
        elif (not self.validar_SKU(sku)):
            return f'SKU: {sku} não foi encontrado.'
        elif (not self.validar_Embalagem(embalagem)):
            return f'Embalagem: {embalagem} não está cadastrada.'
    
    def validar_SKU(self, sku):
        SQL = f'SELECT count(*) FROM produto WHERE id_Produto = {sku!r}'
        cur = self.conn.cursor()
        cur.execute(SQL)
        recset = cur.fetchall()
        self.conn.commit()
        cur.close()
        return bool(recset[0][0])

    def validar_Embalagem(self, embalagem):
        SQL = f'SELECT count(*) FROM embalagem WHERE id_Embalagem = {embalagem!r}'
        cur = self.conn.cursor()
        cur.execute(SQL)
        recset = cur.fetchall()
        self.conn.commit()
        cur.close()
        return bool(recset[0][0])

    def listar(self):
        with open("scripts/listarPedidos.sql", "r") as arquivoSQL:
            SQL = arquivoSQL.read()
        cur = self.conn.cursor()
        cur.execute(SQL)
        recset = cur.fetchall()
        self.conn.commit()
        cur.close()
        print(recset)
        retorno = '<table class="table" id="listarTabela"><tr class="tableRow"><th class="tableItem">ID Pedido</th><th class="tableItem">Data</th><th class="tableItem">Produto</th><th class="tableItem">Quantidade</th><th class="tableItem">Valor</th><th class="tableItem">Taxa</th><th class="tableItem">Receita</th><th class="tableItem">Custo</th><th class="tableItem">Embalagem</th><th class="tableItem">Lucro</th></tr>'
        for rec in recset:
            retorno += '<tr>'
            for item in rec:
                retorno += '<td class="tableItem">' + str(item) + '</td>'
            retorno += '</tr>'
        retorno += '</table>'
        return retorno

if __name__ == '__main__':
    sender = Sender()
    sender.run(host='0.0.0.0', port=8080, debug=True)