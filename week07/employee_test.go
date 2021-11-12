package week07

import (
	"testing"

)

func TestEmployeeIsValid(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		id		int
		err 	error
	}{
		{name: "test invalid employee", id: -1, err: ErrInvalidEmployee(-1)},
		{name: "test invalid employee", id: 0, err: ErrInvalidEmployee(0)},
		{name: "test valid employee", id: 1, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			emp := Employee(tc.id)

			act := emp.IsValid()
			exp := tc.err

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestEmployeeWork(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		id		int
		err 	error
	}{
		{name: "test invalid employee", id: 0, err: ErrInvalidEmployee(0)},
		{name: "test valid employee", id: 1, err: nil},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := NewManager()
			defer m.Stop()

			emp := Employee(tc.id)
			go emp.work(m)

			p := &Product{Quantity: 1} // valid product
			err := m.Assign(p)

			if err != nil {
				t.Fatalf("error starting manager during test: %s", err)
			}

			exp := tc.err

			select {
			case act := <-m.Errors():
				if exp != act {
					t.Fatalf("expected %s, got %s", exp, act)
				}
			case cp := <-m.Completed():
				if cp.Product.Quantity != p.Quantity {
					t.Fatalf("got unexpected product")
				}
			}

			m.Stop()
		})
	}
}

func TestEmployeeCancel(t *testing.T) {
	t.Parallel()
	m := NewManager()
	defer m.Stop()

	emp := Employee(1)

	go m.Stop()

	emp.work(m)
}