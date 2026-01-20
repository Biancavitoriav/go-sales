# 🛒 Go Sales - Sistema de Vendas com MongoDB

Sistema de vendas em Go usando Gin Framework e MongoDB.

## 📁 Arquitetura

```
go-sales/
├── cmd/api/main.go
├── internal/
│   ├── models/
│   ├── repository/
│   ├── service/
│   ├── handler/
│   └── database/
├── .env.example
└── go.mod
```

## 🚀 Como Rodar

### 1. Subir MongoDB com Docker

```bash
docker run -d --name mongodb -p 27017:27017 mongo:latest
```

### 2. Configurar variáveis de ambiente

```bash
cp .env.example .env
```

### 3. Instalar dependências

```bash
go mod tidy
```

### 4. Rodar aplicação

```bash
go run cmd/api/main.go
```

## 📡 Endpoints

### Clientes

```bash
# Criar
POST /customers
{
  "name": "Alice",
  "phone": "(11) 98765-4321"
}

# Listar
GET /customers

# Buscar por ID
GET /customers/:id
```

### Produtos

```bash
# Criar
POST /products
{
  "name": "Notebook",
  "price": 2999.90
}

# Listar
GET /products

# Buscar por ID
GET /products/:id
```

## 🧪 Testes

```bash
# Criar cliente
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice", "phone": "11987654321"}'

# Listar clientes
curl http://localhost:8080/customers

# Criar produto
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{"name": "Mouse", "price": 49.90}'
```

## 🛠️ Stack

- Go 1.21+
- Gin Framework
- MongoDB
- Docker

## 📝 Validações

- Nome do cliente: mínimo 3 caracteres
- Telefone: 10-11 dígitos numéricos
- Nome do produto: mínimo 3 caracteres
- Preço: não pode ser negativo
