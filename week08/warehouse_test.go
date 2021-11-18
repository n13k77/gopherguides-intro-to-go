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
			t.Log("cleaning up context")
			w.Stop()
			cancel()
		}()

		t.Log("start warehouse")

		// TODO
		// I do not get what I could do with the returned context
		_ = w.Start(ctx)

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