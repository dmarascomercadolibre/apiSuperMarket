package domain

type Product struct {
	ID int `json:"id"`
	AtributtesProduct
}

type AtributtesProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	Code_value  string  `json:"code_value"`
	IsPublished bool    `json:"is_published`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
