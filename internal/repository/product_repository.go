package repository

import "vendas/internal/domain"

type ProductRepository interface {
	Save(product domain.Product) (domain.Product, error)
	FindAll() ([]domain.Product, error)
}
