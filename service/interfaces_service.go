// service/interfaces.go
package service

import "github.com/ichdavid/go_pedidos_api/model"

type ClienteServiceI interface {
	Create(c *model.Cliente) error
	Get(id int) (model.Cliente, error)
	GetAll() ([]model.Cliente, error)
	Update(c *model.Cliente) error
	Delete(id int) error
}

type PedidoServiceI interface {
	Create(c *model.Pedido) error
	Get(id int) (model.Pedido, error)
	GetAll() ([]model.Pedido, error)
	Update(c *model.Pedido) error
	Delete(id int) error
}

type ProdutoServiceI interface {
	Create(c *model.Produto) error
	Get(id int) (model.Produto, error)
	GetAll() ([]model.Produto, error)
	Update(c *model.Produto) error
	Delete(id int) error
}
