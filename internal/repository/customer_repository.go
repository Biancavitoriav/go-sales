package repository

import "vendas/internal/domain"

type CustomerRepository interface {
	Save(customer domain.Customer) (domain.Customer, error)
	FindAll() ([]domain.Customer, error)
}
