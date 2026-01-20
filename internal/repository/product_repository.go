package repository

import "vendas/internal/models"

type ProductRepository struct {
	data   []models.Product
	nextID int64
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		data:   []models.Product{},
		nextID: 1,
	}
}

func (r *ProductRepository) Save(product models.Product) (models.Product, error) {
	product.ID = r.nextID
	r.nextID++
	r.data = append(r.data, product)
	return product, nil
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	return r.data, nil
}

func (r *ProductRepository) FindByID(id int64) (models.Product, error) {
	for _, product := range r.data {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, nil
}
