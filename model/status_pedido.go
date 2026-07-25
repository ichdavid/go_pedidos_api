package model

// Definição das constantes dos status do pedido.
type StatusPedido string

const (
	StatusPending  StatusPedido = "PENDING"
	StatusPaid     StatusPedido = "PAID"
	StatusCanceled StatusPedido = "CANCELED"
)
