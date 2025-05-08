# CRUD API em Go (Golang)

Este é um projeto simples de uma API REST em Go que implementa operações básicas de CRUD (Create, Read, Update e Delete) usando o banco de dados MySQL. O projeto utiliza os pacotes `net/http` e `github.com/gorilla/mux` para roteamento e `database/sql` com `go-sql-driver/mysql` para conexão com o banco de dados.

## Tecnologias Utilizadas

- Golang 1.22.2
- MySQL
- Gorilla Mux
- go-sql-driver/mysql

## Estrutura do Projeto

```
.
├── database
│   └── database.go        # Lógica de conexão com o banco de dados
├── server
│   └── server.go          # Implementação das rotas e funções CRUD
├── go.mod
├── go.sum
└── main.go                # Inicialização da aplicação e configuração das rotas
```

## Rotas Disponíveis

- `POST /user` – Cria um novo usuário
- `GET /users` – Lista todos os usuários
- `GET /user/{id}` – Busca um usuário específico
- `PUT /user/{id}` – Atualiza um usuário
- `DELETE /user/{id}` – Deleta um usuário

## Exemplo de Requisição

### POST /user

```json
{
  "nome": "João da Silva",
  "email": "joao@email.com"
}
```

### Resposta:

```
Status: 201 Created
Usuário inserido com sucesso - ID: 1
```

## Configuração do Banco de Dados

Certifique-se de que você tenha um banco de dados MySQL rodando com a seguinte configuração no arquivo `database.go`:

```go
stringConnection := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
```

> **Nota:** Altere o usuário, senha e nome do banco conforme sua configuração local.

## Como Rodar o Projeto

1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Execute o servidor:
   ```bash
   go run main.go
   ```

A aplicação irá rodar em `http://localhost:5000`.

## Autor

Francisco Neto

---

Projeto de estudo prático em Golang com foco em APIs REST e MySQL.
