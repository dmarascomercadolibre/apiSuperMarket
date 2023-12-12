package domain

// Product represents a product in the super market.
type Product struct {
	ID int `json:"id"` // The unique identifier of the product.
	AtributtesProduct
}

// AtributtesProduct represents the attributes of a product.
type AtributtesProduct struct {
	Name        string  `json:"name"`         // Name of the product
	Quantity    int     `json:"quantity"`     // Quantity of the product
	Code_value  string  `json:"code_value"`   // Code value of the product
	IsPublished *bool   `json:"is_published"` // Indicates if the product is published
	Expiration  string  `json:"expiration"`   // Expiration date of the product
	Price       float64 `json:"price"`        // Price of the product
}
