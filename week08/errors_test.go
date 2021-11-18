package week08

import (
	"testing"
)

func TestError(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		err  error
		exp  string
	}{
		{name: "ErrInvalidMaterials", err: ErrInvalidMaterials(-1), exp: "materials must be greater than 0, got -1"},
		{name: "ErrProductNotBuilt", err: ErrProductNotBuilt("teststring"), exp: "teststring"},
		{name: "ErrInvalidEmployee", err: ErrInvalidEmployee(-1), exp: "invalid employee number: -1"},
		{name: "ErrInvalidEmployeeCount", err: ErrInvalidEmployeeCount(0), exp: "invalid employee count: 0"},
		{name: "ErrManagerStopped", err: ErrManagerStopped{}, exp: "manager is stopped"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			act := tc.err.Error()
			exp := tc.exp

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}

		})
	}
}
