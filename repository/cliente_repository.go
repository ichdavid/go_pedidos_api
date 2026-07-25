package repository

//Estrutura do repository
type Repository struct {
	dados map[int]Cliente
}

//Construtor
func NewRepository() *Repository {
	return &Repository{dados: make(map[int]Cliente)}
}

//metodo create para adicionar as informaçõe do usuario no banco
func (r *Repository) Create(c Cliente) {
	r.dados[c.ID] = c
}

//metodo Get para buscar um unico usuario por ID
func (r *Repository) Get(id int) (Cliente, bool) {
	c, ok := r.dados[id]
	return c, ok
}

//metodo para buscar todos os usuarios
func (r *Repository) GetAll() []Cliente {
	clientes := make([]Cliente, 0, len(r.dados))

	for _, c := range r.dados {
		clientes = append(clientes, c)
	}

	return clientes
}

//metodo para atualizar as informações do cliente no banco de dados
func (r *Repository) Update(c Cliente) {
	r.dados[c.ID] = c
}

//metodo para remover um usuario do banco de dados
func (r *Repository) Delete(id int) {
	delete(r.dados, id)
}
