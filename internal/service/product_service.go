package service

import (
	"vendas/internal/models"
	"vendas/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(name string, price float64) (*models.Product, error) {
	product := &models.Product{
		Name:  name,
		Price: price,
	}
	err := s.repo.Save(product)
	return product, err
}

func (s *ProductService) List() ([]*models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id string) (*models.Product, error) {
	return s.repo.FindByID(id)
}
