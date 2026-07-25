package repository

//Estrutura do repository
type Repository struct {
	dados map[int]Pedido
}

//Construtor
func NewRepository() *Repository {
	return &Repository{dados: make(map[int]Pedido)}
}

//metodo create para adicionar as informaçõe do pedido no banco
func (r *Repository) Create(pd Pedido) {
	r.dados[pd.ID] = pd
}

//metodo Get para buscar um unico pedido por ID
func (r *Repository) Get(id int) (Pedido, bool) {
	pd, ok := r.dados[id]
	return pd, ok
}

//metodo para buscar todos os pedido
func (r *Repository) GetAll() []Pedido {
	pedidos := make([]Pedido, 0, len(r.dados))

	for _, pd := range r.dados {
		pedidos = append(pedidos, pd)
	}

	return pedidos
}

//metodo para atualizar as informações do pedido no banco de dados
func (r *Repository) Update(pd Pedido) {
	r.dados[pd.ID] = pd
}

//metodo para remover um pedido do banco de dados
func (r *Repository) Delete(id int) {
	delete(r.dados, id)
}
