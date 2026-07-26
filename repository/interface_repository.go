package repository

import "github.com/ichdavid/go_pedidos_api/model"

type ClienteRepositoryI interface {
	Create(c model.Cliente)
	Get(id int) (model.Cliente, bool)
	GetAll() []model.Cliente
	Update(c model.Cliente)
	Delete(id int)
}

type PedidoRepositoryI interface {
	Create(c model.Pedido)
	Get(id int) (model.Pedido, bool)
	GetAll() []model.Pedido
	Update(c model.Pedido)
	Delete(id int)
}

type ProdutoRepositoryI interface {
	Create(c model.Produto)
	Get(id int) (model.Produto, bool)
	GetAll() []model.Produto
	Update(c model.Produto)
	Delete(id int)
}
