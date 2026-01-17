package memory

import "vendas/internal/domain"

type CustomerRepo struct {
	data []domain.Customer
}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{}
}

func (r *CustomerRepo) Save(c domain.Customer) (domain.Customer, error) {
	c.ID = int64(len(r.data) + 1)
	r.data = append(r.data, c)
	return c, nil
}

func (r *CustomerRepo) FindAll() ([]domain.Customer, error) {
	return r.data, nil
}
