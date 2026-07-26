package model

// Definição da entidade pedidos
type Pedido struct {
	ID        int
	ClienteID int
	Itens     []ItemPedido
	Status    StatusPedido
}
