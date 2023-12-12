package product

import (
	"app/internal/domain"
)

type ProductService struct {
	ProductRepository *ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		ProductRepository: NewProductRepository(),
	}
}

func (p *ProductService) GetAllProducts() []domain.Product {
	return p.ProductRepository.GetAll()
}
func (p *ProductService) GetByID(id int) domain.Product {
	return p.ProductRepository.GetByID(id)
}

func (p *ProductService) CreateProduct(product domain.Product) {

	idGreater := p.getHighestID()
	product.ID = idGreater + 1

	p.ProductRepository.Create(product)
}

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
