package data

import "testing"

func TestProduct_Validate(t *testing.T) {
	p := &Product{
		Name:  "dadw",
		Price: 1.00,
		SKU:   "adb-dasda-dad",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
