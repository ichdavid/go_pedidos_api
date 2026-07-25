package model

// Definição da estrutura de item de pedidos
type ItemPedido struct {
	Produto    Produto
	Preco      float64
	Quantidade int
}
