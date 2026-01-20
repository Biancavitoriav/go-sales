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

### 1. Configurar MongoDB

Escolha uma das opções:

**Opção A: MongoDB com Docker (Local)**
```bash
docker run -d --name mongodb -p 27017:27017 mongo:latest
```

**Opção B: MongoDB instalado localmente**
```bash
# Se já tem MongoDB instalado, apenas inicie o serviço
mongod
```

**Opção C: MongoDB Atlas (Remoto/Cloud)**
- Crie uma conta em [mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas)
- Crie um cluster gratuito
- Copie a connection string

### 2. Configurar variáveis de ambiente

```bash
cp .env.example .env
```

Edite o `.env` conforme seu ambiente:

```env
# Local (Docker ou instalado)
MONGODB_URI=mongodb://localhost:27017

# Remoto (MongoDB Atlas)
MONGODB_URI=mongodb+srv://usuario:senha@cluster.mongodb.net/?retryWrites=true&w=majority

# Com autenticação local
MONGODB_URI=mongodb://admin:senha@localhost:27017

MONGODB_DATABASE=go_sales
PORT=8080
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

## 🌐 MongoDB Atlas (Cloud Gratuito)

Para usar MongoDB na nuvem gratuitamente:

1. Acesse [mongodb.com/cloud/atlas](https://www.mongodb.com/cloud/atlas/register)
2. Crie uma conta gratuita
3. Crie um cluster (Free tier M0 - 512MB grátis)
4. Em "Database Access", crie um usuário e senha
5. Em "Network Access", adicione seu IP (ou `0.0.0.0/0` para aceitar qualquer IP)
6. Clique em "Connect" → "Connect your application"
7. Copie a connection string e cole no `.env`:

```env
MONGODB_URI=mongodb+srv://seu_usuario:sua_senha@cluster0.xxxxx.mongodb.net/?retryWrites=true&w=majority
```
