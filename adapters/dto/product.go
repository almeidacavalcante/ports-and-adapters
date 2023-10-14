package dto

import "github.com/almeidacavalcante/ports-and-adapters/application"

type Product struct {
	ID     string  `json:"id" valid:"uuidv4"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	p.ID = product.ID
	p.Name = product.Name
	p.Price = product.Price
	p.Status = product.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	return product, nil
}
