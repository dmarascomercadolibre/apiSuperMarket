package product

import (
	"app/internal/domain"
	"app/pkg"
)

type ProductRepository struct {
	ProductRepository []domain.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		ProductRepository: pkg.FullfilDBProduct("../products.json"),
	}
}

func (p *ProductRepository) GetAll() []domain.Product {
	return p.ProductRepository
}

func (p *ProductRepository) GetByID(id int) domain.Product {
	for _, product := range p.ProductRepository {
		if product.ID == id {
			return product
		}
	}
	return domain.Product{}
}

func (p *ProductRepository) Create(product domain.Product) {
	p.ProductRepository = append(p.ProductRepository, product)
}
