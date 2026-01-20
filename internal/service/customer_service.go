package service

import (
	"errors"
	"strings"
	"unicode"
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
	if len(name) < 3 {
		return models.Customer{}, errors.New("customer name must have at least 3 characters")
	}

	cleanPhone := cleanPhoneNumber(phone)

	if !isValidPhone(cleanPhone) {
		return models.Customer{}, errors.New("invalid phone number")
	}

	customer := models.Customer{
		Name:  name,
		Phone: cleanPhone,
	}
	return s.repo.Save(customer)
}

func cleanPhoneNumber(phone string) string {
	var result strings.Builder

	for _, char := range phone {
		if unicode.IsDigit(char) {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func isValidPhone(phone string) bool {
	length := len(phone)
	return length >= 10 && length <= 11
}

func (s *CustomerService) List() ([]models.Customer, error) {
	return s.repo.FindAll()
}

func (s *CustomerService) GetByID(id string) (models.Customer, error) {
	return s.repo.FindByID(id)
}
