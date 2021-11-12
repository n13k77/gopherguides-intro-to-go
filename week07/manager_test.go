package week07

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestManagerStart(t *testing.T) {
	t.Parallel()
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
			defer m.Stop()

			act := m.Start(tc.emp)
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestManagerAssignStopped(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
			defer m.Stop()

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
	t.Parallel()
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
		defer m.Stop()

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

func TestRun(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name		string
		employees 	int
		products 	[]*Product
		err 		error
	}{
		{name: "run correct product", employees: 1, products: []*Product{{Quantity: 1}}, err: nil,},
		{name: "run incorrect product", employees: 1, products: []*Product{{Quantity: -1}}, err: ErrInvalidQuantity(-1)},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			ctx := context.Background()
			cps, err := Run(ctx, tc.employees, tc.products...)

			if err != tc.err {
				t.Fatalf("unexpected error: expected %s, got %s", tc.err, err)
			}

			if err == nil {
				// Check whether all products are completed
				expLen := len(tc.products)
				actLen := len(cps)
				if actLen != expLen {
					t.Fatalf("error: expected %d products, got %d", expLen, actLen)
				}
			
				// Check whether all products are correct
				for i := 0; i < len(cps); i++ {
					act := cps[i].Product.Quantity
					exp := tc.products[i].Quantity

					if exp != act {
						t.Fatalf("expected product quantity %d, got %d", exp, act)
					}
				}
			}
		})
	}
}
 
func TestRunCancel(t *testing.T) {
	//t.Parallel()
	tc := struct {
		name	string
	}{
		name: "run with cancellation",
	}
	t.Run(tc.name, func(t *testing.T) {

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		sigCtx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()

		// goroutine that will interrupt after 10 milliseconds (ctrl-c)
		go func() {
			time.Sleep(10 * time.Millisecond)
			t.Log("syscall triggered")
			syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		}()

		// 1 piece of product takes 1 millisecond to be produced
		// so we assign 10000 products, to make sure that we are longer
		// producing than sleeping, allowing the Run process to be
		// interrupted by the SIGINT
		_, err := Run(sigCtx, 1, &Product{Quantity: 10000})
	
		if err != nil {
			t.Fatalf("unexpected error while testing for cancellation")
		}
		
		select {
		case <-ctx.Done():
			t.Log("context finished")
		case <-sigCtx.Done():
			t.Log("sigCtx context finished")
			return
		}
		
		err = ctx.Err()
		if err == nil {
			return
		}
		
		if err == context.DeadlineExceeded {
			t.Fatal("unexpected error", err)
		}
	})
}

func TestRunTimeout(t *testing.T) {
	t.Parallel()
	tc := struct {
		name	string
	}{
		name: "run with timeout",
	}
	t.Run(tc.name, func(t *testing.T) {

		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		// Again, we assign 10000 products, to allow the Run process to be
		// interrupted by the timeout
		_, err := Run(ctx, 1, &Product{Quantity: 10000})

		if err != nil {
			t.Fatalf("unexpected error while testing for cancellation")
		}
		
		// wait for the context to finish
		<-ctx.Done()

		err = ctx.Err()

		if err != context.DeadlineExceeded {
			t.Fatal("unexpected error", err)
		}
	})
}