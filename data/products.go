package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Model     string  `json:"model"`
	Price     float64 `json:"price"`
	SKU       string  `json:"sku,omitempty"`
	CreatedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:        1,
		Name:      "Air Jordan 1",
		Model:     "Nike",
		Price:     95,
		SKU:       "qwe123",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID:        2,
		Name:      "Air Max 95",
		Model:     "Nike",
		Price:     120,
		SKU:       "tyj723",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}
