package week04

import (
	"bytes"
	"fmt"
	"strings"
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
		viewers int
	} {
		{viewers: 1},
		{viewers: 10},
		{viewers: 100},
		{viewers: 1000},
	}

	m := messyEntertainer{}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.viewers,
			Log:		nil,
		}

		err := m.Setup(v)
		act := tc.viewers 

		if act < 100 && err == nil {
			t.Fatalf("expected err for an audience of %d, but got none", act)
		}
	}
}

func TestMessyEntertainerPerform(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		viewers int
	} {
		{viewers: 1},
		{viewers: 10},
		{viewers: 100},
		{viewers: 1000},
	}

	m := messyEntertainer{}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.viewers,
			Log:		nil,
		}

		err := m.Perform(v)
		act := tc.viewers 

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
			Log:		nil,
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
			Log:		nil,
		}
		err := c.Teardown(v)

		act := tc.a 

		if act > 100 && err == nil {
			t.Fatalf("expected err for an audience of %d, but got none", act)
		}
	}
}

func TestShow(t *testing.T) {
	t.Parallel()

	// test for cleanEntertainer
	// initialize variables
	cb := new(bytes.Buffer)
	cv := Venue {
		Audience:	10,
		Log:		cb,
	}
	c := cleanEntertainer{
		FullName: "Full Name",
	}

	// give show
	err := cv.show(c)
	if err != nil {
		t.Fatal(err)
	}

	// test the resulting log 
	cexp := [2]string{
		"has completed teardown",
		"has performed for",
	}

	for _, element := range(cexp) {
		if !strings.Contains(cb.String(),element) {
			t.Fatalf("expected log %s not found", element)
		}
	}

	// test for messyEntertainer
	// initialize variables
	mb := new(bytes.Buffer)
	mv := Venue {
		Audience:	200,
		Log:		mb,
	}

	m := messyEntertainer{
		FullName: "Full Name",
	}

	// give show
	err = mv.show(m)
	if err != nil {
		t.Fatal(err)
	}

	mexp := [2]string{
		"has completed setup",
		"has performed for",
	}

	for _, element := range(mexp) {
		if !strings.Contains(mb.String(),element) {
			t.Fatalf("expected log %s not found", element)
		}
	}
}

func TestEntertain(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		viewers 		int
		entertainers 	int
	}{
		{viewers: 1, 	entertainers: 0},
		{viewers: 100, 	entertainers: 0},
		{viewers: 3000,	entertainers: 0},
		{viewers: 10, 	entertainers: 10},
		{viewers: 15, 	entertainers: 20},
		{viewers: 1000,	entertainers: 200},
	}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.viewers,
			Log:		&bytes.Buffer{},
		}

		e := []Entertainer{}
		c := cleanEntertainer{
			FullName: "Full Name",
		}
		m := messyEntertainer{
			FullName: "Full Name",
		}

		// fill the entertainers slice with both types of entertainers in order
		// to include them both in the test
		for i := 0; i < tc.entertainers; i++ {
			e = append(e, c)
			e = append(e, m)
		}

		err := v.Entertain(tc.viewers, e...)

		// checking a very specific error condition; whether Entertain returns an error
		// for 0 viewers. All other errors conditions coming from the show function 
		// through the Entertain function are tested elsewhere
		if tc.entertainers == 0 && !strings.Contains(fmt.Sprint(err), "no entertainers to perform") {
			t.Fatalf("expected err for %d entertainers", tc.entertainers)
		}
	}
}


func TestMessyEntertain(t *testing.T) {
	t.Parallel()

	testCases := []struct{
		viewers 		int
		entertainers 	int
	}{
		{viewers: 110, 	entertainers: 0},
		{viewers: 200, 	entertainers: 10},
	}

	for _, tc := range testCases {

		v := Venue {
			Audience:	tc.viewers,
			Log:		&bytes.Buffer{},
		}

		me := []Entertainer{}
		m := messyEntertainer{}

		for i := 0; i < tc.entertainers; i++ {
			me = append(me, m)
		}

		err := v.Entertain(tc.viewers, me...)

		if tc.entertainers <= 0 && err == nil {
			t.Fatalf("expected err for %d viewers and %d entertainers", tc.viewers, tc.entertainers)
		}
	}
}
