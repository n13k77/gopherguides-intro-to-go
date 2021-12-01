package week08

import (
	"context"
	"testing"
)

func TestProductBuiltBy(t *testing.T) {
	t.Parallel()
	tc := struct {
		name	string
		emp 	Employee
		p 		Product
	}{
		name: 	"product built by",
		emp: 	Employee(1),
		p:		*ProductA,
	}

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

		err := tc.p.Build(tc.emp, w)

		if err != nil {
			t.Fatalf("unexpected error while building product, got %s", err)
		}

		act := tc.p.BuiltBy()
		exp := tc.emp

		if exp != act {
			t.Fatalf("expected %d, got %d", exp, act)
		}
	})
}

func TestProductIsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		p 		Product
		err 	error
	}{
		{name: "valid product", p: Product{Materials: Materials{Metal: 1250}}, err: nil},
		{name: "tricky valid product", p: Product{Materials: Materials{Metal: 0}}, err: nil},
		{name: "invalid product", p: Product{Materials: Materials{}}, err: ErrInvalidMaterials(0)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			act := tc.p.IsValid() 
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestProductIsBuilt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		emp 	Employee
		p  		Product
		err 	error
	}{
		{name: "product is built of invalid product", emp: Employee(1), p: Product{Materials: Materials{}}, err: ErrInvalidMaterials(0)},
		{name: "product is built by invalid employee", emp: Employee(-1), p: Product{Materials: Materials{Metal: 1250}}, err: ErrProductNotBuilt("product is not built: [{metal:1250x}]")},
		{name: "valid product is built", emp: Employee(1), p: Product{Materials: Materials{Metal: 1250}}, err: nil},
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

			act := tc.p.IsBuilt()
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}
