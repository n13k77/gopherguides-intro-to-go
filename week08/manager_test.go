package week08

import (
	"context"
	"testing"
)

func TestManagerStart(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		emp  int
		err  error
	}{
		{name: "start manager with invalid employee number", emp: -1, err: ErrInvalidEmployeeCount(-1)},
		{name: "start manager with valid employee number", emp: 2, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			m := &Manager{}

			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)

			defer func() {
				t.Log("cleaning up context")
				cancel()
			}()

			_, act := m.Start(ctx, tc.emp)
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}

			// no error was returned, so manager was started properly
			// this means it has to be stopped
			if act == nil {
				m.Stop()
			}
		})
	}
}

func TestManagerAssignStopped(t *testing.T) {
	t.Parallel()
	tc := struct {
		name string
		err  error
	}{
		name: "assign to stopped manager",
		err:  ErrManagerStopped{},
	}

	t.Run(tc.name, func(t *testing.T) {

		m := &Manager{}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		_, err := m.Start(ctx, 1)

		defer func() {
			m.Stop()
			cancel()
		}()

		if err != nil {
			t.Fatalf("error starting manager during test")
		}

		m.Stop()

		act := m.Assign(&Product{Materials: Materials{Metal: 1250}})
		exp := tc.err

		if exp != act {
			t.Fatalf("expected %s, got %s", exp, act)
		}
	})
}

func TestManagerAssign(t *testing.T) {

	t.Parallel()
	testCases := []struct {
		name string
		p    []*Product
		err  error
	}{
		{
			name: "assign invalid product",
			p:    []*Product{{Materials: Materials{}}},
			err:  ErrInvalidMaterials(0),
		},
		{
			name: "assign one product",
			p:    []*Product{{Materials: Materials{Metal: 1250}}},
			err:  nil,
		},
		{
			name: "assign multiple products",
			p:    []*Product{{Materials: Materials{Metal: 1250}}, {Materials: Materials{Oil: 250}}},
			err:  nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := &Manager{}

			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
	
			ctx, err := m.Start(ctx, 5)

			if err != nil {
				t.Fatalf("unexpected error starting up manager")
			}

			defer func() {
				m.Stop()
				cancel()
			}()

			// start producing, assign products to manager
			go func() {
				err := m.Assign(tc.p...)
				if err != nil {
					m.Errors() <- err
				}
			}()

			var act []CompletedProduct

			// start receiving, we need to listen on the completed channel otherwise it blocks.
			// stop the manager when he has finished all products
			go func() {
				// here we need to wait until all is received, so a loop is put in place
				for product := range m.Completed() {
					act = append(act, product)
					if len(act) == len(tc.p) {
						m.Stop()
					}
				}
			}()

			// wait for one signal, not forever, so no loop is put in place
			select {
			case err, ok := <-m.Errors():
				if ok && err.Error() != tc.err.Error() {
					t.Fatal(err)
				}
			case <-ctx.Done(): 
				return
			}
		})
	}
}

func TestManagerComplete(t *testing.T) {
	t.Parallel()
	tc := struct {
		name		string
		product 	Product
	}{
		name: 		"complete a product",
		product: 	Product{Materials: Materials{Metal: 1250}},
	}

	t.Run(tc.name, func(t *testing.T) {
		m := &Manager{}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)

		ctx, err := m.Start(ctx, 5)

		if err != nil {
			t.Fatalf("unexpected error starting up manager")
		}

		defer func() {
			m.Stop()
			cancel()
		}()

		// start producing
		go func() {
			err := m.Assign(&tc.product)
			if err != nil {
				m.Errors() <- err
			}
		}()

		var cp []CompletedProduct

		// start receiving, stop the manager when he has finished all products
		go func() {
			for product := range m.Completed() {
				cp = append(cp, product)
				if len(cp) == 1 {
					m.Stop()
				}
			}
		}()

		select {
		case err, ok := <-m.Errors():
			if ok{
				t.Fatal(err)
			}
		case <-ctx.Done():
		}

		act := cp[0].Product.String()
		exp := tc.product.String()

		if exp != act {
			t.Fatalf("expected %s, got %s", exp, act)
		}

	})
}
