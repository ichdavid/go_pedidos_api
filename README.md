# go_pedidos

Núcleo de um serviço de pedidos em Go, desenvolvido como desafio de domínio e regras de negócio.  
O foco inicial é a parte interna da aplicação — sem HTTP, banco de dados ou frameworks externos.

## 🎯 Objetivo
Construir um sistema capaz de:
- Representar produtos, pedidos e itens.
- Validar regras de negócio (estoque, cliente, quantidade, status).
- Manter repositórios em memória usando `map`.
- Coordenar casos de uso via services.
- Tratar erros de forma explícita (sem `panic`).

## 📦 Entidades
- **cliente**: ID, Name, Email, PasswordHash, CreatedAt.
- **Produto**: id, nome, preço, estoque.
- **Item de Pedido**: produto, preço no momento da compra, quantidade.
- **Pedido**: id, cliente, itens, statuspedido
- **statuspedido** (`PENDING`, `PAID`, `CANCELED`).

## 🔑 Regras principais
- Cliente precisa ter nome.
- Pedido precisa ter pelo menos um item.
- Quantidade deve ser maior que zero.
- Produto deve existir e ter estoque suficiente.
- Estoque é reduzido ao criar pedido e devolvido ao cancelar.
- Pedido nasce como `PENDING`.
- Pedido pode ser pago (`PAID`) ou cancelado (`CANCELED`).
- Pedidos pagos ou cancelados não podem mudar de status novamente.

## ⚠️ Erros esperados
- Produto não encontrado.
- Pedido não encontrado.
- Quantidade inválida.
- Estoque insuficiente.
- Cliente inválido.
- Pedido vazio.
- Mudança de status inválida.

## 🗂 Estrutura sugerida
pedidos/
  go.mod
  main.go
  product.go
  order.go
  errors.go
  repositories.go
  services.go

## 🚀 Como rodar

Pré-requisitos: Go instalado na máquina.

No terminal:

git clone https://github.com/ichdavid/go_pedidos_api.git
cd go_pedidos_api
go mod init pedidos
go run main.go
