package service

import (
	"github.com/ichdavid/go_pedidos_api/errors"
	"github.com/ichdavid/go_pedidos_api/model"
	"github.com/ichdavid/go_pedidos_api/repository"
)

type PedidoService struct {
	repo *repository.PedidoRepository
}

func NewPedidoService(repo *repository.PedidoRepository) *PedidoService {
	return &PedidoService{repo: repo}
}

func (ps *PedidoService) Create(pd *model.Pedido) error {

	_, ok := ps.repo.Get(pd.ID)

	if ok {
		return errors.ErrPedidoJaCadastrado
	}

	if len(pd.Itens) == 0 {
		return errors.ErrPedidoVazio
	}

	for _, itens := range pd.Itens {
		if itens.Quantidade <= 0 {
			return errors.ErrQuantidadeMinimaProduto
		}
		if itens.Produto.Estoque < itens.Quantidade {
			return errors.ErrEstoqueInsuficiente
		}

		itens.Produto.Estoque -= itens.Quantidade
	}

	pd.Status = model.StatusPending

	ps.repo.Create(*pd)
	return nil
}

func (ps *PedidoService) Get(id int) (model.Pedido, error) {
	pedido, ok := ps.repo.Get(id)

	if !ok {
		return model.Pedido{}, errors.ErrPedidoNaoEncontrado
	}

	return pedido, nil

}

func (ps *PedidoService) GetAll() ([]model.Pedido, error) {

	pedidos := ps.repo.GetAll()
	return pedidos, nil

}

func (ps *PedidoService) UpdateStatus(id int, novoStatus model.StatusPedido) error {

	pedido, ok := ps.repo.Get(id)

	if !ok {
		return errors.ErrPedidoNaoEncontrado
	}

	if pedido.Status == model.StatusPaid || pedido.Status == model.StatusCanceled {
		return errors.ErrMudancaStatusInvalida
	}

	if novoStatus == model.StatusCanceled {
		for _, itens := range pedido.Itens {
			itens.Produto.Estoque += itens.Quantidade
		}
	}

	pedido.Status = novoStatus
	ps.repo.Update(pedido)
	return nil

}

func (ps *PedidoService) Update(pd *model.Pedido) error {
	_, ok := ps.repo.Get(pd.ID)

	if !ok {
		return errors.ErrPedidoNaoEncontrado
	}

	ps.repo.Update(*pd)

	return nil
}

func (ps *PedidoService) Delete(id int) error {
	_, ok := ps.repo.Get(id)

	if !ok {
		return errors.ErrPedidoNaoEncontrado
	}

	ps.repo.Delete(id)
	return nil
}
