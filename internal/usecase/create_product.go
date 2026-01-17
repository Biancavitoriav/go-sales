package usecase

import (
	"vendas/internal/domain"
	"vendas/internal/repository"
)

type CreateProduct struct {
	repo repository.ProductRepository
}

func NewCreateProduct(r repository.ProductRepository) *CreateProduct {
	return &CreateProduct{repo: r}
}

func (uc *CreateProduct) Execute(name string, price float64) (domain.Product, error) {
	product := domain.Product{
		Name:  name,
		Price: price,
	}
	return uc.repo.Save(product)
}
