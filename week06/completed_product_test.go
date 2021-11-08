package week06

import (
	"testing"
)

func TestCompletedProductIsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		emp 	Employee
		product Product
		err 	error
	}{
		{name: "validity of invalid product", emp: Employee(1), product: Product{Quantity: -1} , err: ErrInvalidQuantity(-1)},
		{name: "validity of product built by invalid employee", emp: Employee(-1), product: Product{Quantity: 1}, err: ErrInvalidEmployee(-1)},
		{name: "validity of valid product", emp: Employee(1), product: Product{Quantity: 1}, err: nil},
	}
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			_ = tc.product.Build(tc.emp)

			cp := CompletedProduct{
				Product: 	tc.product,
				Employee: 	tc.emp,
			}

			act := cp.IsValid()
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}
