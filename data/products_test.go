package data

import "testing"

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "react",
		Price: 200,
		SKU:   "iab990",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal((err))
	}
}
