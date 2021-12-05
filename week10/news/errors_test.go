package news

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
		{name: "ErrArticleNotFound", err: ErrArticleNotFound(int(123)), exp: "article 123 not found"},
		{name: "ErrAlreadySubscribed", err: ErrAlreadySubscribed("teststring"), exp: "teststring"},
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
