# 🛒 Go Sales - Sistema de Vendas

Sistema simples de vendas construído em Go usando **Standard Layout Architecture**.

## 📁 Arquitetura

```
go-sales/
├── cmd/
│   └── api/
│       └── main.go              # Ponto de entrada da aplicação
├── internal/
│   ├── models/                  # Modelos de dados (Customer, Product)
│   │   ├── customer.go
│   │   └── product.go
│   ├── repository/              # Camada de persistência
│   │   ├── customer_repository.go
│   │   └── product_repository.go
│   ├── service/                 # Lógica de negócio
│   │   ├── customer_service.go
│   │   └── product_service.go
│   └── handler/                 # Handlers HTTP (controllers)
│       ├── customer_handler.go
│       └── product_handler.go
├── go.mod
├── go.sum
└── README.md
```

## 🏗️ Camadas da Aplicação

### **Models** (`internal/models/`)
Estruturas de dados que representam as entidades do negócio.

### **Repository** (`internal/repository/`)
Responsável pelo armazenamento e recuperação de dados. 
Atualmente usa armazenamento em memória (dados são perdidos ao reiniciar).

### **Service** (`internal/service/`)
Contém a lógica de negócio e validações.
- Validação de dados
- Regras de negócio
- Orquestração de operações

### **Handler** (`internal/handler/`)
Lida com requisições HTTP e respostas.
- Parse de requisições
- Validação de entrada
- Formatação de respostas JSON

## 🚀 Como Rodar

### Pré-requisitos
- Go 1.21 ou superior
- Git

### Instalação

1. Clone o repositório:
```bash
git clone <seu-repositório>
cd go-sales
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Execute a aplicação:
```bash
go run cmd/api/main.go
```

A API estará disponível em `http://localhost:8080`

## 📡 Endpoints da API

### Clientes

#### Criar Cliente
```bash
POST /customers
Content-Type: application/json

{
  "name": "João Silva"
}
```

**Resposta:**
```json
{
  "id": 1,
  "name": "João Silva"
}
```

#### Listar Clientes
```bash
GET /customers
```

**Resposta:**
```json
[
  {
    "id": 1,
    "name": "João Silva"
  },
  {
    "id": 2,
    "name": "Maria Santos"
  }
]
```

#### Buscar Cliente por ID
```bash
GET /customers/:id
```

**Resposta:**
```json
{
  "id": 1,
  "name": "João Silva"
}
```

### Produtos

#### Criar Produto
```bash
POST /products
Content-Type: application/json

{
  "name": "Notebook",
  "price": 2999.90
}
```

**Resposta:**
```json
{
  "id": 1,
  "name": "Notebook",
  "price": 2999.90
}
```

#### Listar Produtos
```bash
GET /products
```

**Resposta:**
```json
[
  {
    "id": 1,
    "name": "Notebook",
    "price": 2999.90
  },
  {
    "id": 2,
    "name": "Mouse",
    "price": 49.90
  }
]
```

#### Buscar Produto por ID
```bash
GET /products/:id
```

**Resposta:**
```json
{
  "id": 1,
  "name": "Notebook",
  "price": 2999.90
}
```

## 🧪 Testando a API

### Usando cURL

**Criar um cliente:**
```bash
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva"}'
```

**Listar clientes:**
```bash
curl http://localhost:8080/customers
```

**Criar um produto:**
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Notebook", "price": 2999.90}'
```

**Listar produtos:**
```bash
curl http://localhost:8080/products
```

**Buscar cliente por ID:**
```bash
curl http://localhost:8080/customers/1
```

## 📝 Validações

### Cliente
- Nome deve ter no mínimo 3 caracteres

### Produto
- Nome deve ter no mínimo 3 caracteres
- Preço não pode ser negativo

## 🔄 Fluxo de Dados

```
HTTP Request
    ↓
Handler (valida entrada)
    ↓
Service (aplica regras de negócio)
    ↓
Repository (persiste dados)
    ↓
Service (retorna resultado)
    ↓
Handler (formata resposta JSON)
    ↓
HTTP Response
```

## 🛠️ Tecnologias Utilizadas

- **Go 1.21+** - Linguagem de programação
- **Gin** - Framework web HTTP
- **Standard Layout** - Arquitetura de organização de código

## 📦 Dependências

```
github.com/gin-gonic/gin v1.10.0
```

## 🚧 Próximos Passos

- [ ] Implementar persistência em banco de dados (PostgreSQL)
- [ ] Adicionar testes unitários
- [ ] Implementar autenticação JWT
- [ ] Adicionar módulo de vendas
- [ ] Criar documentação Swagger
- [ ] Adicionar logging estruturado
- [ ] Implementar tratamento de erros customizado

## 💡 Por que Standard Layout?

Esta arquitetura foi escolhida por ser:
- ✅ **Simples** - Fácil de entender e manter
- ✅ **Organizada** - Separação clara de responsabilidades
- ✅ **Escalável** - Pode crescer conforme necessário
- ✅ **Padrão de mercado** - Usado em 70% dos projetos Go

## 📚 Aprendizado

Este projeto demonstra:
- Organização de código em Go
- Injeção de dependências manual
- Separação de camadas (Handler → Service → Repository)
- API RESTful
- Validações de negócio

## 🤝 Contribuindo

Sinta-se livre para abrir issues ou pull requests!

## 📄 Licença

MIT
