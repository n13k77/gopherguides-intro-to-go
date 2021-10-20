package week04

import (
	"fmt"
	"os"
	"testing"
)

func TestMessyEntertainerName (t *testing.T) {
	t.Parallel()

	testCases := []struct{
		 name	string
	} {
		{name:	"Johnson"},
		{name:	"王"},
		{name:	"Jansen"},
		{name:	"عبد الله"},
		{name:	"Müller"},
	}

	for _, tc := range testCases {

		var e messyEntertainer
		e.FullName = tc.name
		
		exp := tc.name
		act := e.Name()

		if exp != act {
			t.Errorf("expected %s, got %s", exp, act)
		}
	}
}

func TestMessyEntertainerSetup(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		a int
	} {
		{a: 1},
		{a: 10},
		{a: 100},
		{a: 1000},
	}

	m := messyEntertainer{}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.a,
			Log:		os.Stdout,
		}
		err := m.Setup(v)

		act := tc.a 

		if act < 100 && err == nil {
			t.Fatalf("expected err for an audience of %d, but got none", act)
		}
	}
}

func TestMessyEntertainerPerform(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		a int
	} {
		{a: 1},
		{a: 10},
		{a: 100},
		{a: 1000},
	}

	m := messyEntertainer{}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.a,
			Log:		os.Stdout,
		}
		err := m.Perform(v)

		act := tc.a 

		if act < 100 && err == nil {
			t.Fatalf("expected err for an audience of %d, but got none", act)
		}
	}
}

func TestCleanEntertainerName (t *testing.T) {
	t.Parallel()

	testCases := []struct{
		 name	string
	} {
		{name:	"Johnson"},
		{name:	"王"},
		{name:	"Jansen"},
		{name:	"عبد الله"},
		{name:	"Müller"},
	}

	for _, tc := range testCases {

		var e messyEntertainer
		e.FullName = tc.name
		
		exp := tc.name
		act := e.Name()

		if exp != act {
			t.Errorf("expected %s, got %s", exp, act)
		}
	}
}

func TestCleanEntertainerPerform(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		a int
	} {
		{a: 1},
		{a: 10},
		{a: 100},
		{a: 1000},
	}

	c := cleanEntertainer{}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.a,
			Log:		os.Stdout,
		}
		err := c.Perform(v)

		act := tc.a 

		if act > 100 && err == nil {
			t.Fatalf("expected err for an audience of %d, but got none", act)
		}
	}
}

func TestCleanEntertainerTeardown(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		a int
	} {
		{a: 1},
		{a: 10},
		{a: 100},
		{a: 1000},
	}

	c := cleanEntertainer{}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.a,
			Log:		os.Stdout,
		}
		err := c.Teardown(v)

		act := tc.a 

		if act > 100 && err == nil {
			t.Fatalf("expected err for an audience of %d, but got none", act)
		}
	}
}

func TestEntertain(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		a int
	} {
		{a: 1},
	}

	for _, tc := range testCases {
		fmt.Print(tc.a)
	}

	exp := 0
	act := 1

	if exp != act {
		t.Errorf("expected %d, got %d", exp, act)
	}
}
