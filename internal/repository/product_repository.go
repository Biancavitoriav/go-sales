package repository

import "vendas/internal/models"

// ProductRepository gerencia o armazenamento de produtos
type ProductRepository struct {
	data   []models.Product
	nextID int64
}

// NewProductRepository cria uma nova instância do repositório
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		data:   []models.Product{},
		nextID: 1,
	}
}

// Save salva um produto e retorna com o ID atribuído
func (r *ProductRepository) Save(product models.Product) (models.Product, error) {
	product.ID = r.nextID
	r.nextID++
	r.data = append(r.data, product)
	return product, nil
}

// FindAll retorna todos os produtos
func (r *ProductRepository) FindAll() ([]models.Product, error) {
	return r.data, nil
}

// FindByID busca um produto por ID
func (r *ProductRepository) FindByID(id int64) (models.Product, error) {
	for _, product := range r.data {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, nil
}
