package service

import (
	"github.com/ichdavid/go_pedidos_api/errors"
	"github.com/ichdavid/go_pedidos_api/model"
	"github.com/ichdavid/go_pedidos_api/repository"

	"time"
)

type ClienteService struct {
	repo *repository.ClienteRepository
}

func NewClienteService(repo *repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (cs *ClienteService) Create(c *model.Cliente) error {
	if c.Nome == "" {
		return errors.ErrNomeObrigatorio
	}

	_, ok := cs.repo.Get(c.ID)

	if ok {
		return errors.ErrClienteJaCadastrado
	}

	c.CreatedAt = time.Now()

	cs.repo.Create(*c)

	return nil
}

func (cs *ClienteService) Get(id int) (model.Cliente, error) {
	cliente, ok := cs.repo.Get(id)

	if !ok {
		return model.Cliente{}, errors.ErrClienteNaoEncontrado
	}

	return cliente, nil
}

func (cs *ClienteService) GetAll() ([]model.Cliente, error) {

	clientes := cs.repo.GetAll()
	return clientes, nil
}

func (cs *ClienteService) Update(c *model.Cliente) error {

	_, ok := cs.repo.Get(c.ID)

	if !ok {
		return errors.ErrClienteNaoEncontrado
	}

	cs.repo.Update(*c)

	return nil

}

func (cs *ClienteService) Delete(id int) error {
	_, ok := cs.repo.Get(id)

	if !ok {
		return errors.ErrClienteNaoEncontrado
	}

	cs.repo.Delete(id)

	return nil
}
