package week08

import (
	"context"
	"testing"
)

func TestWarehouseRetrieve(t *testing.T) {
	tc := struct {
		name	string
		m		Material
		q 		int
	}{
		name: "test warehouse retrieve",
		m: Plastic, 
		q: 1,
	}
	t.Run(tc.name, func(t *testing.T) {
		
		w := &Warehouse{}

		ctx := context.Background() 
		ctx, cancel := context.WithCancel(ctx)
		defer func(){
			w.Stop()
			cancel()
		}()

		ctx = w.Start(ctx)

		m, err := w.Retrieve(tc.m, tc.q)
		if err != nil {
			t.Fatalf("error retrieving product")
		}

		act := m
		exp := tc.m
		if exp != act {
			t.Fatalf("expected %s, got %s", exp, act)
		}
	})
}