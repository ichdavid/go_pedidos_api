package service

import (
	"errors"

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
		return errors.New("Pedido ja existe!")
	}

	ps.repo.Create(*pd)
	return nil
}

func (ps *PedidoService) Get(id int) (model.Pedido, error) {
	pedido, ok := ps.repo.Get(id)

	if !ok {
		return model.Pedido{}, errors.New("Pedido não encontrado!")
	}

	return pedido, nil

}

func (ps *PedidoService) GetAll() ([]model.Pedido, error) {

	pedidos := ps.repo.GetAll()
	return pedidos, nil

}

func (ps *PedidoService) Update(pd *model.Pedido) error {
	_, ok := ps.repo.Get(pd.ID)

	if !ok {
		return errors.New("Pedido não encontrado")
	}

	ps.repo.Update(*pd)

	return nil
}

func (ps *PedidoService) Delete(id int) error {
	_, ok := ps.repo.Get(id)

	if !ok {
		return errors.New("Pedido não encontrado")
	}

	ps.repo.Delete(id)
	return nil
}
