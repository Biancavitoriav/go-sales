package service

import (
	"errors"
	"vendas/internal/models"
	"vendas/internal/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Create(name string) (models.Customer, error) {
	if len(name) < 3 {
		return models.Customer{}, errors.New("customer name must have at least 3 characters")
	}

	customer := models.Customer{Name: name}
	return s.repo.Save(customer)
}

func (s *CustomerService) List() ([]models.Customer, error) {
	return s.repo.FindAll()
}

func (s *CustomerService) GetByID(id int64) (models.Customer, error) {
	return s.repo.FindByID(id)
}
