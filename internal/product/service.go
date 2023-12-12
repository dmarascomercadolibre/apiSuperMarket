package product

import (
	"app/internal/domain"
)

// ProductService represents a service for managing products.
type ProductService struct {
	ProductRepository *ProductRepository
}

// NewProductService creates a new instance of ProductService.
func NewProductService() *ProductService {
	return &ProductService{
		ProductRepository: NewProductRepository(),
	}
}

// GetAllProducts returns all products.
func (p *ProductService) GetAllProducts() []domain.Product {
	return p.ProductRepository.GetAll()
}

// GetByID returns the product with the specified ID.
func (p *ProductService) GetByID(id int) domain.Product {
	return p.ProductRepository.GetByID(id)
}

// CreateProduct creates a new product.
func (p *ProductService) CreateProduct(product domain.Product) {
	idGreater := p.getHighestID()
	product.ID = idGreater + 1
	p.ProductRepository.Create(product)
}

// getHighestID returns the highest ID among all products.
func (p *ProductService) getHighestID() int {
	products := p.ProductRepository.GetAll()
	highestID := 0

	for _, product := range products {
		if product.ID > highestID {
			highestID = product.ID
		}
	}

	return highestID
}
