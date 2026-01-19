package service

import (
	"errors"
	"vendas/internal/models"
	"vendas/internal/repository"
)

// ProductService contém a lógica de negócio para produtos
type ProductService struct {
	repo *repository.ProductRepository
}

// NewProductService cria uma nova instância do serviço
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// Create cria um novo produto com validações
func (s *ProductService) Create(name string, price float64) (models.Product, error) {
	// Validações de negócio
	if len(name) < 3 {
		return models.Product{}, errors.New("product name must have at least 3 characters")
	}
	
	if price < 0 {
		return models.Product{}, errors.New("price cannot be negative")
	}

	product := models.Product{
		Name:  name,
		Price: price,
	}
	return s.repo.Save(product)
}

// List retorna todos os produtos
func (s *ProductService) List() ([]models.Product, error) {
	return s.repo.FindAll()
}

// GetByID busca um produto por ID
func (s *ProductService) GetByID(id int64) (models.Product, error) {
	return s.repo.FindByID(id)
}
