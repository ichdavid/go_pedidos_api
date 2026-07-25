package repository

//Estrutura do repository
type Repository struct {
	dados map[int]Produto
}

//Construtor
func NewRepository() *Repository {
	return &Repository{dados: make(map[int]Produto)}
}

//metodo create para adicionar as informaçõe do produto no banco
func (r *Repository) Create(p Produto) {
	r.dados[p.ID] = p
}

//metodo Get para buscar um unico produto por ID no banco
func (r *Repository) Get(id int) (Produto, bool) {
	p, ok := r.dados[id]
	return p, ok
}

//metodo para buscar todos os produtos do banco
func (r *Repository) GetAll() []Produto {
	produtos := make([]Produto, 0, len(r.dados))

	for _, p := range r.dados {
		produtos = append(produtos, p)
	}

	return produtos
}

//metodo para atualizar as informações do produto no banco de dados
func (r *Repository) Update(p Produto) {
	r.dados[p.ID] = p
}

//metodo para remover um produto do banco de dados
func (r *Repository) Delete(id int) {
	delete(r.dados, id)
}
