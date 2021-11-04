package week06

import (
	"testing"
)

func TestProductBuiltBy(t *testing.T) {
	tc := struct {
		name	string
		emp 	Employee
		product Product
	}{
		name: 		"product built by", 
		emp: 		Employee(1), 
		product: 	Product{Quantity: 1}, 
	}

	t.Run(tc.name, func(t *testing.T) {

		if err := tc.product.Build(tc.emp); err != nil {
			t.Fatalf("error building product during test: %s", err)
		}

		act := tc.product.BuiltBy()
		exp := tc.emp

		if exp != act {
			t.Fatalf("expected %d, got %d", exp, act)
		}
	})
}

// Godoc reads:
// Build builds the product by the given employee.
// Returns an error if the product has already been built.
// Returns an error if the employee ID <= 0.
// Returns an error if the quantity <= 0.
// ---
// From the implementation of func (p *Product) Build(e Employee) 
// I do not see how the last condition is checked in that function
func TestProductBuild(t *testing.T) {
	testCases := []struct {
		name	string
		emp 	Employee
		product Product
		err 	error
	}{
		{name: "product build by invalid employee", emp: Employee(-1), product: Product{Quantity: 1}, err: ErrInvalidEmployee(-1)},
		{name: "product build of invalid product", emp: Employee(1), product: Product{Quantity: -1} , err: ErrInvalidQuantity(-1)},
		{name: "valid product build", emp: Employee(1), product: Product{Quantity: 1}, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.product.Build(tc.emp)
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestProductIsValid(t *testing.T) {
	testCases := []struct {
		name	string
		qty 	int
		err 	error
	}{
		{name: "product invalid quantity", qty: -1, err: ErrInvalidQuantity(-1)},
		{name: "product zero quantity", qty: 0, err: ErrInvalidQuantity(0)},
		{name: "product valid quantity", qty: 1, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			p := Product{Quantity: tc.qty}
			act := p.IsValid() 
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestProductIsBuilt(t *testing.T) {
	testCases := []struct {
		name	string
		emp 	Employee
		product Product
		err 	error
	}{
		{name: "product is built of invalid product", emp: Employee(1), product: Product{Quantity: -1}, err: ErrInvalidQuantity(-1)},
		{name: "product is built by invalid employee", emp: Employee(-1), product: Product{Quantity: 1}, err: ErrProductNotBuilt("product is not built: {1 0}")},
		{name: "valid product is built", emp: Employee(1), product: Product{Quantity: 1}, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_ = tc.product.Build(tc.emp)

			act := tc.product.IsBuilt()
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}
