package erros

import "errors"

var (
	//generico
	ErrNomeObrigatorio = errors.New("nome obrigatório")

	// Cliente
	ErrClienteNaoEncontrado = errors.New("cliente não encontrado")
	ErrClienteJaCadastrado  = errors.New("cliente já cadastrado")

	// Produto
	ErrProdutoNaoEncontrado    = errors.New("produto não encontrado")
	ErrProdutoJaCadastrado     = errors.New("produto já cadastrado")
	ErrPrecoInvalido           = errors.New("preço inválido")
	ErrQuantidadeMinimaProduto = errors.New("mínimo 1 unidade do produto")

	// Pedido
	ErrPedidoVazio           = errors.New("pedido não pode estar vazio")
	ErrPedidoNaoEncontrado   = errors.New("pedido não encontrado")
	ErrPedidoJaCadastrado    = errors.New("pedido já cadastrado")
	ErrEstoqueInsuficiente   = errors.New("quantidade insuficiente no estoque")
	ErrMudancaStatusInvalida = errors.New("Mudança de status invalida")
)
