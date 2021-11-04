package week06

import (
	"testing"
)

func TestManagerStart(t *testing.T) {
	testCases := []struct {
		name	string
		emp 	int
		err 	error
	}{
		{name: "start manager with invalid employee number", emp: -1, err: ErrInvalidEmployeeCount(-1)},
		{name: "start manager with valid employee number", emp: 2, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			m := NewManager()
			act := m.Start(tc.emp)
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestManagerAssignStopped(t *testing.T) {
	tc := struct {
		name	string
		err 	error
	}{
		name: 		"assign to stopped manager", 
		err:		ErrManagerStopped{},
	}

	t.Run(tc.name, func(t *testing.T) {

		mgr := NewManager()
		mgr.Stop()

		act := mgr.Assign(&Product{Quantity: 1})
		exp := tc.err

		if exp != act {
			t.Fatalf("expected %s, got %s", exp, act)
		}
	})
}

func TestManagerAssign(t *testing.T) {
	testCases := []struct {
		name		string
		products 	[]*Product
		err 		error
	}{
		{name: "assign one product", products: []*Product{{Quantity: 1}}, err: nil},
		{name: "assign invalid product", products: []*Product{{Quantity: -1}}, err: ErrInvalidQuantity(-1)},
		{name: "assign multiple products", products: []*Product{{Quantity: 1}, {Quantity: 1}}, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			// start manager
			m := NewManager()

			// cannot put this block inside the starting block because then 
			// t.Fatalf would be called from inside a non-testing function
			err := m.Start(5)
			if err != nil {
				t.Fatalf("error starting manager during test: %s", err)
			}

			// start producing, assign products to manager
			go func() {
				err := m.Assign(tc.products...)
				if err != nil {
					m.Errors() <- err
				}
			}()
			
			// this snippet is exactly like in the example, but in my defense, I 
			// use `act` all over the place in tests :)
			var act []CompletedProduct

			// start receiving, we need to listen on the completed channel otherwise it blocks.
			// stop the manager when he has finished all products
			go func() {
				// here we need to wait until all is received, so a loop is put in place
				for product := range m.Completed() {
					act = append(act, product)
					if len(act) == len(tc.products) {
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
			case <-m.Done():
			}
		})
	}
}

 func TestManagerComplete(t *testing.T) {
	tc := struct {
		name		string
		product 	Product
		err 		error
	}{
		name: 		"complete a product", 
		product: 	Product{Quantity: 1}, 
		err: 		nil,
	}

	t.Run(tc.name, func(t *testing.T) {

		// start manager
		m := NewManager()
		err := m.Start(5)
		if err != nil {
			t.Fatalf("error starting manager during test: %s", err)
		}

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
		case <-m.Done():
		}
		
		act := cp[0].Product.Quantity
		exp := tc.product.Quantity

		if exp != act {
			t.Fatalf("expected %d, got %d", exp, act)
		}

	})
}


// func TestManagerDone(t *testing.T) {
// 	testCases := []struct {
// 		name	string
		
// 	}{
// 		{name: ""},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
			
// 		})
// 	}
// }

// func TestManagerJobs(t *testing.T) {
// 	testCases := []struct {
// 		name	string
		
// 	}{
// 		{name: ""},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
			
// 		})
// 	}
// }

// func TestManagerErrors(t *testing.T) {
// 	testCases := []struct {
// 		name	string
		
// 	}{
// 		{name: ""},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
			
// 		})
// 	}
// }

// func TestManagerStop(t *testing.T) {
// 	testCases := []struct {
// 		name	string
		
// 	}{
// 		{name: ""},
// 	}
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
			
// 		})
// 	}
// }
