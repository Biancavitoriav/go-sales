package service

import (
	"vendas/internal/models"
	"vendas/internal/repository"
)

type CustomerService struct {
	repo *repository.CustomerRepository
}

func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) Create(name, phone string) (models.Customer, error) {
	customer := models.Customer{
		Name:  name,
		Phone: phone,
	}
	return s.repo.Save(customer)
}

func (s *CustomerService) List() ([]models.Customer, error) {
	return s.repo.FindAll()
}

func (s *CustomerService) GetByID(id string) (models.Customer, error) {
	return s.repo.FindByID(id)
}
