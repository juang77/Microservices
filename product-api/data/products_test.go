package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "juancho",
		Price: 1.00,
		SKU:   "abc-abcd-12345",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
