package product

import (
	"app/internal/domain"
	"app/pkg"
)

// ProductRepository represents a repository for managing products.
type ProductRepository struct {
	ProductRepository []domain.Product
}

// NewProductRepository creates a new instance of ProductRepository.
// It initializes the repository with data from the products.json file.
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		ProductRepository: pkg.FullfilDBProduct("../products.json"),
	}
}

// GetAll returns all the products in the repository.
func (p *ProductRepository) GetAll() []domain.Product {
	return p.ProductRepository
}

// GetByID returns the product with the specified ID from the repository.
// If no product is found, it returns an empty product.
func (p *ProductRepository) GetByID(id int) domain.Product {
	for _, product := range p.ProductRepository {
		if product.ID == id {
			return product
		}
	}
	return domain.Product{}
}

// Create adds a new product to the repository.
func (p *ProductRepository) Create(product domain.Product) {
	p.ProductRepository = append(p.ProductRepository, product)
}

// Update updates the product with the given ID in the repository.
// It replaces the existing product with the new product provided.
// If the product with the given ID is not found, it returns an empty product and no error.
func (p *ProductRepository) Update(product domain.Product) (pr domain.Product, err error) {
	for i, prod := range p.ProductRepository {
		if prod.ID == product.ID {
			p.ProductRepository[i] = product
			prod = product
			return
		}
	}
	return
}
