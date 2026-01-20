package repository

import "vendas/internal/models"

type CustomerRepository struct {
	data   []models.Customer
	nextID int64
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		data:   []models.Customer{},
		nextID: 1,
	}
}

func (r *CustomerRepository) Save(customer models.Customer) (models.Customer, error) {
	customer.ID = r.nextID
	r.nextID++
	r.data = append(r.data, customer)
	return customer, nil
}

func (r *CustomerRepository) FindAll() ([]models.Customer, error) {
	return r.data, nil
}

func (r *CustomerRepository) FindByID(id int64) (models.Customer, error) {
	for _, customer := range r.data {
		if customer.ID == id {
			return customer, nil
		}
	}
	return models.Customer{}, nil
}
