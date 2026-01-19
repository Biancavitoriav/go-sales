package service

import (
	"errors"
	"vendas/internal/models"
	"vendas/internal/repository"
)

// CustomerService contém a lógica de negócio para clientes
type CustomerService struct {
	repo *repository.CustomerRepository
}

// NewCustomerService cria uma nova instância do serviço
func NewCustomerService(repo *repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

// Create cria um novo cliente com validações
func (s *CustomerService) Create(name string) (models.Customer, error) {
	// Validações de negócio
	if len(name) < 3 {
		return models.Customer{}, errors.New("customer name must have at least 3 characters")
	}

	customer := models.Customer{Name: name}
	return s.repo.Save(customer)
}

// List retorna todos os clientes
func (s *CustomerService) List() ([]models.Customer, error) {
	return s.repo.FindAll()
}

// GetByID busca um cliente por ID
func (s *CustomerService) GetByID(id int64) (models.Customer, error) {
	return s.repo.FindByID(id)
}
