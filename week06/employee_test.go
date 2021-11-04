package week06

import (
	"testing"

)

func TestEmployeeIsValid(t *testing.T) {
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
