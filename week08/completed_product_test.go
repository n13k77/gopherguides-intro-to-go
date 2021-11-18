package week08

import (
	"context"
	"testing"
)

func TestCompletedProductIsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		emp 	Employee
		p 		Product
		err 	error
	}{
		{name: "validity of invalid product", emp: Employee(1), p: Product{Materials: Materials{}} , err: ErrInvalidMaterials(0)},
		{name: "validity of product built by invalid employee", emp: Employee(-1), p: Product{Materials: Materials{Metal: 1}}, err: ErrInvalidEmployee(-1)},
		{name: "validity of valid product", emp: Employee(1), p: Product{Materials: Materials{Metal: 1}}, err: nil},
	}
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			w := &Warehouse{}

			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)

			defer func(){
				t.Log("cleaning up context")
				w.Stop()
				cancel()
			}()

			_ = w.Start(ctx)

			_ = tc.p.Build(tc.emp, w)

			cp := CompletedProduct{
				Product: 	tc.p,
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
