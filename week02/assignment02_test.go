package assigment02_test

import (
	"testing"
)

func TestArray(t *testing.T) {
	exp := [4]string{"John", "Paul", "George", "Ringo"}
	act := [4]string{}

	// Iterate through the exp variable and copy its values into the act variable.
	for i, n := range exp {
		act[i] = n
	}

	// Iterate through the act variable and assert that its contents match that of the exp variable.
	for i := range act {
		if act[i] != exp[i] {
			t.Errorf("element %d of arrays do not match", i)
		}
	}
}

func TestSlice(t *testing.T) {
	exp := []string{"John", "Paul", "George", "Ringo"}
	act := []string{}

	// Iterate through the exp variable and copy its values into the act variable.
	for _, n := range exp {
		act = append(act, n)
	}

	// Iterate through the act variable and assert that its contents match that of the exp variable.
	for i := range act {
		if act[i] != exp[i] {
			t.Errorf("element %d of slices do not match", i)
		}
	}

	// Assert that the length of act and exp are the same.
	if len(act) != len(exp) {
		t.Errorf("length of slices is not the same")
	}
}

func TestMap(t *testing.T) {
	exp := map[int]string{
		1: "John",
		2: "Paul",
		3: "George",
		4: "Ringo",
	}
	act := map[int]string{}

	// Iterate through the exp variable and copy its values into the act variable.
	for key, value := range exp {
		act[key] = value
	}

	// Iterate through the act variable and assert that its contents match that of the exp variable.
	// Assert that the key being requested from exp exists
	for key, value := range act {
		expectedValue, ok := exp[key]
		if !ok {
			t.Errorf("key %d does not exist in exp", key)
		}
		if value != expectedValue {
			t.Errorf("values do not match between maps for key %d", key)
		}
	}

	// Assert that the length of act and exp are the same.
	if len(act) != len(exp) {
		t.Errorf("length is not the same")
	}
}
