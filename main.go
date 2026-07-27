package gopedidosapi

import (
	"github.com/ichdavid/go_pedidos_api/model"
	"github.com/ichdavid/go_pedidos_api/repository"
	"github.com/ichdavid/go_pedidos_api/service"
)

func main() {

	clienteRepo := repository.NewClienteRepository()
	produtoRepo := repository.NewProdutoRepository()
	pedidoRepo := repository.NewPedidoRepository()

	clienteService := service.NewClienteService()
	produtoService := service.NewProdutoService()
	pedidoService := service.NewPedidoService()

	cliente1 := model.Cliente{ID: 1, Nome: "David", PasswordHash: "123456"}
	cliente2 := model.Cliente{ID: 1, Nome: "Alexandre", PasswordHash: "987654"}

	produto1 := model.Produto{ID: 1, Nome: "Notebook", Preco: 2175.99, Estoque: 20}
	produto2 := model.Produto{ID: 1, Nome: "Mouse", Preco: 80.00, Estoque: 50}

	pedido := model.Pedido{ID: 1, Cliente: cliente1}

}
