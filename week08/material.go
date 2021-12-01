package week08

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const (
	// some handy constants
	Metal   Material = "metal"
	Oil     Material = "oil"
	Plastic Material = "plastic"
	Wood    Material = "wood"
)

// Material is the type of material
// that a product is made of.
type Material string

func (m Material) String() string {
	return string(m)
}

// Duration is the amount of time it takes
// to produce the given material.
// This is based on the length of the
// material string.
// 	Material("metal").Duration() == 5ms
// 	Material("oil").Duration() == 3ms
// 	Material("plastic").Duration() == 7ms
// 	Material("wood").Duration() == 4ms
func (m Material) Duration() time.Duration {
	i := len(m)
	return time.Duration(i) * time.Millisecond
}

// Materials is a list of materials
// and their quantities.
type Materials map[Material]int

// Duration is the amount of time it takes
// to produce the given materials.
func (mats Materials) Duration() time.Duration {
	var d time.Duration
	for m, q := range mats {
		d += m.Duration() * time.Duration(q)
	}
	return d
}

func (mats Materials) String() string {
	lines := make([]string, 0, len(mats))

	for m, q := range mats {
		lines = append(lines, fmt.Sprintf("{%s:%dx}", m, q))
	}
	sort.Strings(lines)
	s := strings.Join(lines, ", ")
	return fmt.Sprintf("[%s]", s)
}
