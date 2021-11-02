package demo

import (
	"fmt"
	"sort"
	"strings"
)

type Clauses map[string]interface{} // collection of any type of thing, each one identified by a string

func (cls Clauses) String() string {
	lines := make([]string, 0, len(cls))

	for k, v := range cls {
		lines = append(lines, fmt.Sprintf("%q = %q", k, v))
	}

	sort.Strings(lines)
	return strings.Join(lines, " and ")
}

func (cls Clauses) Match(m Model) bool {
	for k, v := range cls {
		if m[k] != v {
			return false
		}
	}
	return true
}
