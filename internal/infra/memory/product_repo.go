package memory

import "vendas/internal/domain"

type ProductRepo struct {
	data []domain.Product
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

func (r *ProductRepo) Save(p domain.Product) (domain.Product, error) {
	p.ID = int64(len(r.data) + 1)
	r.data = append(r.data, p)
	return p, nil
}

func (r *ProductRepo) FindAll() ([]domain.Product, error) {
	return r.data, nil
}
