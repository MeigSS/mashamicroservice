package data

import "testing"

func TestProduct_Validate(t *testing.T) {
	p := &Product{
		Name:  "dadw",
		Price: -1,
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
