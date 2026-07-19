package main

// Definição da entidade pedidos
type Pedido struct {
	ID        int
	ClienteID int
	Item      []ItemPedido
	Status    StatusPedido
}
