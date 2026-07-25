package repository

import "github.com/ichdavid/go_pedidos_api/model"

//Estrutura do repository
type ClienteRepository struct {
	dados map[int]model.Cliente
}

//Construtor
func NewClienteRepository() *ClienteRepository {
	return &ClienteRepository{dados: make(map[int]model.Cliente)}
}

//metodo create para adicionar as informaçõe do usuario no banco
func (r *ClienteRepository) Create(c model.Cliente) {
	r.dados[c.ID] = c
}

//metodo Get para buscar um unico usuario por ID
func (r *ClienteRepository) Get(id int) (model.Cliente, bool) {
	c, ok := r.dados[id]
	return c, ok
}

//metodo para buscar todos os usuarios
func (r *ClienteRepository) GetAll() []model.Cliente {
	clientes := make([]model.Cliente, 0, len(r.dados))

	for _, c := range r.dados {
		clientes = append(clientes, c)
	}

	return clientes
}

//metodo para atualizar as informações do cliente no banco de dados
func (r *ClienteRepository) Update(c model.Cliente) {
	r.dados[c.ID] = c
}

//metodo para remover um usuario do banco de dados
func (r *ClienteRepository) Delete(id int) {
	delete(r.dados, id)
}
