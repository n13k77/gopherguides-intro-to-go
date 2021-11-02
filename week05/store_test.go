package demo

import (
	"errors"
	"reflect"
	"testing"
)

func noData(t testing.TB) *Store {
	t.Helper()
	return &Store{}
}

func withData(t testing.TB) *Store {
	t.Helper()
	return &Store{
		data: data{},
	}
}

func withUsers(t testing.TB) *Store {
	t.Helper()

	users := Models{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}

	return &Store{
		data: data{
			"users": users,
		},
	}
}

func assertModel(t testing.TB, act Model, exp Model) bool {
	t.Helper()
	for k, v := range exp {
		if act[k] != v {
			return false
		}
	}
	return true
}

func TestStoreDb(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  string
		store *Store
		equal bool
	}{
		{name: "no data", store: noData(t), equal: false},
		{name: "with data, no users", store: withData(t), equal: true},
		{name: "with users", store: withUsers(t), equal: true},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			exp := tt.store.data
			act := tt.store.db()

			if reflect.DeepEqual(exp, act) != tt.equal {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

 func TestStoreAll(t *testing.T) {
	t.Parallel()

	tn := "users"

	table := []struct {
		name  	string
		store 	*Store
		err 	error
	}{
		{name: "no data", store: noData(t), err: ErrTableNotFound{}},
		{name: "with data, no users", store: withData(t), err: ErrTableNotFound{}},
		{name: "with users", store: withUsers(t), err: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			exp := tt.store.data[tn]
			act, err := tt.store.All(tn)

			if ! errors.Is(err, tt.err) {
				t.Fatalf("expected %v, got %v", tt.err, err)
			} 
			
			if err == nil && !reflect.DeepEqual(exp, act) {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestStoreAllErrors(t *testing.T) {
	t.Parallel()

	tn := "users"

	table := []struct {
		name  string
		store *Store
		exp   error
	}{
		{name: "no data", store: noData(t), exp: ErrTableNotFound{}},
		{name: "with data, no users", store: withData(t), exp: ErrTableNotFound{}},
		{name: "with users", store: withUsers(t), exp: nil},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.store.All(tn)

			ok := errors.Is(err, tt.exp)

			if !ok {
				t.Fatalf("expected error %v, got %v", tt.exp, err)
			}
		})
	}
}

func TestStoreLen(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  	string
		store 	Store
		table 	string
		err 	error
	}{
		{name: "length store", store: *withUsers(t), table: "users", err: nil},
		{name: "length store for non-existing table", store: *withUsers(t), table: "notexistingtable", err: ErrTableNotFound{}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			exp := 2
			act, err := tt.store.Len(tt.table)

			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("did not expect error, got %v", err)
			}

			if err == nil && exp != act {
				t.Fatalf("expected %d, got %d", exp, act)
			}
		})
	}
}

func TestStoreInsert(t *testing.T) {
	t.Parallel()

	table := []struct {
		name   	string
		models 	int
	}{
		{name: "10 models", models: 10},
		{name: "100 models", models: 100},
		{name: "0 models", models: 0},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			var m Models
			// initialize a store with an empty but existing data field
			store := Store{
				data{
					"table": []Model{},
				},
			}

			for i := 0; i < tt.models; i++ {
				newModel := Model{
					"id": "id",
				}
				m = append(m, newModel)
			}

			store.Insert("table", m...)

			act, err := store.Len("table")
			exp := tt.models

			if err != nil {
				t.Fatalf("did not expect error, got %v", err)
			}

			if exp != act {
				t.Fatalf("expected %d, got %d", exp, act)
			}
		})
	}
}

func TestStoreSelect(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  	string
		cls   	Clauses
		result	Models
	}{
		{name: "matching result", cls: Clauses{"id": 1, "name": "John"}, result: Models{{"id": 1, "name": "John"}},}, 
		{name: "empty clause", cls: Clauses{}, result: Models{}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			store := withUsers(t)
			tn := "users"

			act, err := store.Select(tn, tt.cls)
			var exp Models

			if len(tt.cls) == 0 {
				var err error
				exp, err = store.All("users")
				if err != nil {
					t.Fatalf("error running test, got error %v", err)
				}
			} else {
				exp = tt.result
			}

			if err != nil {
				t.Fatalf("did not expect error, got %v", err)
			}

			// ok, I dropped the reflect.DeepEquals :)
			// Thanks for the guidance in class
			for i, m := range exp {
				if ! assertModel(t, act[i], m) {
					t.Fatalf("expected %s, got %s", m, act[i])
				}
			}
		})
	}
}

func TestStoreSelectError(t *testing.T) {
	t.Parallel()

	table := []struct {
		name  	string
		table 	string
		cls   	Clauses
		err		error
	}{
		{name: "non existing table", table: "nonexistingtable", cls: Clauses{"id": 1, "name": "John"}, err: ErrTableNotFound{}},
		{name: "no matching result", table: "users", cls: Clauses{"id": 1, "name": "Pete"}, err: &errNoRows{}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			store := withUsers(t)

			exp := tt.err
			_, err := store.Select(tt.table, tt.cls)

			if err != nil && !errors.Is(err, tt.err) {
				t.Fatalf("expected error %v, got %v", exp, err)
			}
	 	})
	}
}

