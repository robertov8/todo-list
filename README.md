# API Todo List em Golang

Esta é uma API RESTful para gerenciamento de tarefas (todo list) desenvolvida em Golang, seguindo boas práticas de organização de código e arquitetura de software.

## Tecnologias Utilizadas

* Go (Golang)
* Gorilla Mux (para gerenciamento de rotas)
* UUID (para geração de identificadores únicos)

## Estrutura do Projeto
```
todo-list/
├── cmd/
│   └── api/
│       └── main.go         # Ponto de entrada da aplicação
├── internal/
│   ├── handlers/
│   │   └── task_handler.go # Handlers HTTP para as requisições
│   ├── middleware/
│   │   └── logging.go      # Middlewares da aplicação
│   ├── models/
│   │   └── task.go         # Definição dos modelos de dados
│   └── repository/
│       ├── repository.go        # Interface do repositório
│       └── memory_repository.go # Implementação em memória
├── go.mod                  # Dependências do Go
├── Dockerfile              # Configuração do Docker
└── README.md               # Documentação
```

## Instalação e Execução

### Pré-requisitos
* Go 1.24 ou superior instalado
* Git instalado

### Passos para Executar

1. Clone o repositório:
```bash
git clone https://github.com/robertov8/todo-list.git
cd todo-list
```

2. Instale as dependências:
```bash
go mod download
```

3. Execute a aplicação:
```bash
go run cmd/api/main.go
```

O servidor será iniciado na porta 4000 por padrão.

## Usando Docker
Alternativamente, você pode usar Docker para executar a aplicação:

```bash
docker build -t todo-list .
docker run -p 4000:4000 todo-list
```

## Rotas da API

### Obter todas as tarefas

````
GET /api/tasks
````

Parâmetros de consulta opcionais:

* done: filtrar por tarefas concluídas (true) ou não concluídas (false)

### Obter uma tarefa específica

```
GET /api/tasks/{id}
```

### Criar uma nova tarefa

```
POST /api/tasks
```
Corpo da requisição:

```json
{
  "title": "Nome da tarefa",
  "description": "Descrição da tarefa"
}
```

### Atualizar uma tarefa existente
```
PUT /api/tasks/{id}
```
Corpo da requisição:

```json
{
  "title": "Novo título",
  "description": "Nova descrição",
  "done": true
}
```

### Excluir uma tarefa
```
DELETE /api/tasks/{id}
```

## Arquitetura
A aplicação segue os princípios de arquitetura limpa e separação de responsabilidades:

1. Modelos (models): Definem as estruturas de dados utilizadas na aplicação.
2. Repositório (repository): Responsável pelo acesso e manipulação dos dados.
    * Interface: Define os métodos para interagir com os dados.
    * Implementação em memória: Implementa a interface usando um mapa para armazenamento em memória.

3. Handlers: Responsáveis por lidar com as requisições HTTP, validar entradas e formatar respostas.
Middleware: Funções que são executadas antes de cada requisição, como logging.

## Próximos Passos

Para tornar esta aplicação pronta para produção, considere implementar:

1. Persistência de dados em um banco de dados real (PostgreSQL, MongoDB, etc.)
2. Autenticação e autorização de usuários
3. Paginação para listagem de tarefas
4. Testes automatizados
5. CI/CD para implantação contínua
6. Documentação da API com Swagger