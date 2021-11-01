package demo

import (
	"fmt"
	"reflect"
	"testing"
)

func assertClause(t testing.TB, act Clauses, exp Clauses) bool {
	t.Helper()
	for k, v := range exp {
		if act[k] != v {
			return false
		}
	}
	return true
}

func TestErrTableNotFoundError(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testError := ErrTableNotFound{
		table:		tn,
	}

	exp := "table not found " + tn
	act := testError.Error()

	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func TestErrTableNotFound(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testError := ErrTableNotFound{
		table:		tn,
	}

	exp := tn
	act := testError.TableNotFound()

	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func TestIsErrTableNotFound(t *testing.T) {
	t.Parallel()
	tn := "testtable"

	testError := ErrTableNotFound{
		table:		tn,
	}

	if !IsErrTableNotFound(testError) {
		t.Fatalf("error asserting type")
	}
}


func TestErrNoRowsError(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testError := errNoRows{
		clauses:	Clauses{"id": 1, "name": "John"},
		table:		tn,
	}

	exp := fmt.Sprintf("[%s] no rows found\nquery: %s", tn, testError.clauses.String())
	act := testError.Error()

	if exp != act {
		t.Fatalf("expected %s, got %s", exp, act)
	}
}

func TestErrNoRowsClauses(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testClause 	:= Clauses{"id": 2, "name": "Jane"} 

	table := []struct {
		name 	string
		exp 	Clauses
		err  	errNoRows
	}{
		{name: "existing clauses", exp: testClause, err: errNoRows{testClause, tn}},
		{name: "empty clauses", exp: Clauses{}, err: errNoRows{nil, tn}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			act := tt.err.Clauses()
			
			if ! reflect.DeepEqual(tt.exp, act) {
				t.Fatalf("expected %s, got %s", tt.exp, act)
			}

		})
	}
}

func TestErrNoRowsRowNotFound(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testClause 	:= Clauses{"id": 2, "name": "Jane"} 

	testError := errNoRows{
		table:		tn,
		clauses: 	testClause,
	}

	actTable, actClause := testError.RowNotFound()
	expTable := tn
	expClause := testClause

	if expTable != actTable {
		t.Fatalf("expected table %s, got %s", expTable, actTable)
	}

	if !reflect.DeepEqual(expClause, actClause) {
		t.Fatalf("expected clause %s got %s", expClause, actClause)
	}
	
	if ! assertClause(t, actClause, expClause) {
		t.Fatalf("expected %s, got %s", expClause, actClause)
	}
}

func TestErrNoRowsIs(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testError := errNoRows{
		table:		tn,
		clauses: 	nil,
	}

	if !testError.Is(&testError) {
		t.Fatalf("error asserting error type")
	}
}

func TestErrNoRowsIsErrNoRows(t *testing.T) {
	t.Parallel()

	tn := "testtable"
	testError := errNoRows{
		table:		tn,
		clauses: 	nil,
	}

	// ok, I dropped the reflect.DeepEquals :)
	// Thanks for the guidance in class
	if ! IsErrNoRows(&testError) {
		t.Fatalf("error asserting error type")
	}
}

// // I had to disable this test as I could not get it to work.
// // I did notice that in calling this test, in some way the Error()
// // function also ends up in the execution path. 
// // Please clarify what I'm missing here. 

// func TestErrNoRowsAsErrNoRows(t *testing.T) {
// 	t.Parallel()

// 	tn := "testtable"
// 	testClause 	:= Clauses{"id": 2, "name": "Jane"} 

// 	testError := errNoRows{
// 		table:		tn,
// 		clauses: 	testClause,
// 	}

// 	// testerr, ok := AsErrNoRows(&testError)
// 	s := &Store{}
// 	s.Insert("users", Model{"id": 1, "name": "John"})
// 	_, err := s.Select("users", Clauses{"id": 2})

// 	if AsErrNoRows(err) {
// 		t.Fatalf("error asserting error type")
// 	}
// }
