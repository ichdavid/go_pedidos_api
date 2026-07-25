package repository

import "github.com/ichdavid/go_pedidos_api/model"

//Estrutura do repository
type ProdutoRepository struct {
	dados map[int]model.Produto
}

//Construtor
func NewProdutoRepository() *ProdutoRepository {
	return &ProdutoRepository{dados: make(map[int]model.Produto)}
}

//metodo create para adicionar as informaçõe do produto no banco
func (r *ProdutoRepository) Create(p model.Produto) {
	r.dados[p.ID] = p
}

//metodo Get para buscar um unico produto por ID no banco
func (r *ProdutoRepository) Get(id int) (model.Produto, bool) {
	p, ok := r.dados[id]
	return p, ok
}

//metodo para buscar todos os produtos do banco
func (r *ProdutoRepository) GetAll() []model.Produto {
	produtos := make([]model.Produto, 0, len(r.dados))

	for _, p := range r.dados {
		produtos = append(produtos, p)
	}

	return produtos
}

//metodo para atualizar as informações do produto no banco de dados
func (r *ProdutoRepository) Update(p model.Produto) {
	r.dados[p.ID] = p
}

//metodo para remover um produto do banco de dados
func (r *ProdutoRepository) Delete(id int) {
	delete(r.dados, id)
}
