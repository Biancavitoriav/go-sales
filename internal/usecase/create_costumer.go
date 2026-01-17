package usecase

import (
	"vendas/internal/domain"
	"vendas/internal/repository"
)

type CreateCustomer struct {
	repo repository.CustomerRepository
}

func NewCreateCustomer(r repository.CustomerRepository) *CreateCustomer {
	return &CreateCustomer{repo: r}
}

func (uc *CreateCustomer) Execute(name string) (domain.Customer, error) {
	customer := domain.Customer{Name: name}
	return uc.repo.Save(customer)
}
