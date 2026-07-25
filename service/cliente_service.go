package service

import (
	"errors"

	"github.com/ichdavid/go_pedidos_api/model"
	"github.com/ichdavid/go_pedidos_api/repository"
)

type ClienteService struct {
	repo *repository.ClienteRepository
}

func NewClienteService(repo *repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (cs *ClienteService) Create(c *model.Cliente) error {
	if c.Nome == "" {
		return errors.New("Obrigatorio o preenchimento do nome")
	}

	_, ok := cs.repo.Get(c.ID)

	if ok {
		return errors.New("Usuario já cadastrado!")
	}

	cs.repo.Create(*c)

	return nil
}

func (cs *ClienteService) Get(id int) (model.Cliente, error) {
	cliente, ok := cs.repo.Get(id)

	if !ok {
		return model.Cliente{}, errors.New("Usuario não encontrado na base de dados!")
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
		return errors.New("Usuario não encontrado na base de dados")
	}

	cs.repo.Update(*c)

	return nil

}

func (cs *ClienteService) Delete(id int) error {
	_, ok := cs.repo.Get(id)

	if !ok {
		return errors.New("Usuario não encontrado na base de dados")
	}

	cs.repo.Delete(id)

	return nil
}
