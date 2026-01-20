package service

import (
	"errors"
	"vendas/internal/models"
	"vendas/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(name string, price float64) (models.Product, error) {
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

func (s *ProductService) List() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id int64) (models.Product, error) {
	return s.repo.FindByID(id)
}
