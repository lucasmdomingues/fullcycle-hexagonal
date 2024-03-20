package dto

import (
	"github.com/lucasmdomingues/hexagonal/application"
)

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func (p *Product) Bind(product *application.Product) error {
	_, err := product.IsValid()
	if err != nil {
		return err
	}

	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	return nil
}

func NewProduct() *Product {
	return &Product{}
}
