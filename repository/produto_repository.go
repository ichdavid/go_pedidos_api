package repository

import "github.com/ichdavid/go_pedidos_api/model"

//Estrutura do repository
type Produtorepository struct {
	dados map[int]model.Produto
}

//Construtor
func NewProdutoRepository() *Produtorepository {
	return &Produtorepository{dados: make(map[int]model.Produto)}
}

//metodo create para adicionar as informaçõe do produto no banco
func (r *Produtorepository) Create(p model.Produto) {
	r.dados[p.ID] = p
}

//metodo Get para buscar um unico produto por ID no banco
func (r *Produtorepository) Get(id int) (model.Produto, bool) {
	p, ok := r.dados[id]
	return p, ok
}

//metodo para buscar todos os produtos do banco
func (r *Produtorepository) GetAll() []model.Produto {
	produtos := make([]model.Produto, 0, len(r.dados))

	for _, p := range r.dados {
		produtos = append(produtos, p)
	}

	return produtos
}

//metodo para atualizar as informações do produto no banco de dados
func (r *Produtorepository) Update(p model.Produto) {
	r.dados[p.ID] = p
}

//metodo para remover um produto do banco de dados
func (r *Produtorepository) Delete(id int) {
	delete(r.dados, id)
}
