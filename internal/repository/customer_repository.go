package repository

import "vendas/internal/models"

// CustomerRepository gerencia o armazenamento de clientes
type CustomerRepository struct {
	data   []models.Customer
	nextID int64
}

// NewCustomerRepository cria uma nova instância do repositório
func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		data:   []models.Customer{},
		nextID: 1,
	}
}

// Save salva um cliente e retorna com o ID atribuído
func (r *CustomerRepository) Save(customer models.Customer) (models.Customer, error) {
	customer.ID = r.nextID
	r.nextID++
	r.data = append(r.data, customer)
	return customer, nil
}

// FindAll retorna todos os clientes
func (r *CustomerRepository) FindAll() ([]models.Customer, error) {
	return r.data, nil
}

// FindByID busca um cliente por ID
func (r *CustomerRepository) FindByID(id int64) (models.Customer, error) {
	for _, customer := range r.data {
		if customer.ID == id {
			return customer, nil
		}
	}
	return models.Customer{}, nil
}
