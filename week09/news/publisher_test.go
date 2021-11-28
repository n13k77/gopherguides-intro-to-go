package news

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc	string
		path    string
		err 	error
	}{
		{desc: "new publisher, no path given", path: "", err: nil},
		{desc: "new publisher, correct path", path: "", err: nil},
		{desc: "new publisher, incorrect path", path: ""},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			
		})
	}
}