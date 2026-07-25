package service

import (
	"errors"

	"github.com/ichdavid/go_pedidos_api/model"
	"github.com/ichdavid/go_pedidos_api/repository"
)

type ProdutoService struct {
	repo *repository.ProdutoRepository
}

func NewProdutoService(repo *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (cs *ProdutoService) Create(p *model.Produto) error {

	if p.Nome == "" {
		return errors.New("Obrigatorio o preenchimento do nome")
	}

	if p.Preco == 0.0 {
		return errors.New("Produto não pode ter valor 0,0")
	}

	_, ok := cs.repo.Get(p.ID)

	if ok {
		return errors.New("Produto já cadastrado!")
	}

	cs.repo.Create(*p)

	return nil
}

func (cs *ProdutoService) Get(id int) (model.Produto, error) {
	produto, ok := cs.repo.Get(id)

	if !ok {
		return model.Produto{}, errors.New("Produto não encontrado na base de dados!")
	}

	return produto, nil
}

func (cs *ProdutoService) GetAll() ([]model.Produto, error) {

	Produtos := cs.repo.GetAll()
	return Produtos, nil
}

func (cs *ProdutoService) Update(p *model.Produto) error {

	_, ok := cs.repo.Get(p.ID)

	if !ok {
		return errors.New("Produto não encontrado na base de dados")
	}

	cs.repo.Update(*p)

	return nil

}

func (cs *ProdutoService) Delete(id int) error {
	_, ok := cs.repo.Get(id)

	if !ok {
		return errors.New("Produto não encontrado na base de dados")
	}

	cs.repo.Delete(id)

	return nil
}
