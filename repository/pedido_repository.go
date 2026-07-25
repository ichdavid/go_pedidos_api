package repository

import "github.com/ichdavid/go_pedidos_api/model"

//Estrutura do repository
type PedidoRepository struct {
	dados map[int]model.Pedido
}

//Construtor
func NewPedidoRepository() *PedidoRepository {
	return &PedidoRepository{dados: make(map[int]model.Pedido)}
}

//metodo create para adicionar as informaçõe do pedido no banco
func (r *PedidoRepository) Create(pd model.Pedido) {
	r.dados[pd.ID] = pd
}

//metodo Get para buscar um unico pedido por ID
func (r *PedidoRepository) Get(id int) (model.Pedido, bool) {
	pd, ok := r.dados[id]
	return pd, ok
}

//metodo para buscar todos os pedido
func (r *PedidoRepository) GetAll() []model.Pedido {
	pedidos := make([]model.Pedido, 0, len(r.dados))

	for _, pd := range r.dados {
		pedidos = append(pedidos, pd)
	}

	return pedidos
}

//metodo para atualizar as informações do pedido no banco de dados
func (r *PedidoRepository) Update(pd model.Pedido) {
	r.dados[pd.ID] = pd
}

//metodo para remover um pedido do banco de dados
func (r *PedidoRepository) Delete(id int) {
	delete(r.dados, id)
}
