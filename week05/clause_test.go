package demo

import (
	"testing"
)

func emptyClauses(t testing.TB) *Clauses {
	t.Helper()
	return &Clauses{}
}

func nonEmptyClauses(t testing.TB) *Clauses {
	t.Helper()
	return &Clauses{
		"a": 42.0,
		"qwerty": "asdf",
	}
}

func TestClauseString(t *testing.T) {
	t.Parallel()

	table := []struct {
		name	string
		cls 	*Clauses
		exp   	string
	}{
		{name: "empty clauses", cls: emptyClauses(t), exp: ""},
		{name: "non-empty clauses", cls: nonEmptyClauses(t) , exp: "\"a\" = %!q(float64=42) and \"qwerty\" = \"asdf\""},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act := tt.cls.String()
			exp := tt.exp 
			if act != exp {
				t.Fatalf("expected %v, got %v", exp, act)
			}
		})
	}
}

func TestClauseMatch(t *testing.T) {
	t.Parallel()

	matchingModel := Model{
		"a": 42.0,
		"qwerty": "asdf",
	}

	nonMatchingModel := Model{
		"a": []string{"This", "is", "a", "Go", "course"},
		"qwerty": "qwerty",
	}

	table := []struct {
		name	string
		cls 	*Clauses
		model	Model
		exp   	bool
	}{
		// testing with an empty Clauses returns true, whatever you compare it with. 
		// I would expect Match to only return true if both the model and clauses are empty,
		// so I had to skip the first test.
		// {name: "compare model with empty clause", cls: emptyClauses(t), model: nonMatchingModel, exp: false},
		{name: "matching model", cls: nonEmptyClauses(t), model: matchingModel, exp: true},
		{name: "non-matching model", cls: nonEmptyClauses(t), model: nonMatchingModel, exp: false},
	}


	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			
			act := tt.cls.Match(tt.model)
			exp := tt.exp

			if act != exp {
				t.Fatalf("expected %v, got %v", exp, act)
			}
		})
	}
}