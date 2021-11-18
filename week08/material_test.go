package week08

import (
	"testing"
	"time"
)

func TestMaterial(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name		string
		material 	Material
		str 		string
	}{
		{name: "string metal", material: Metal, str: "metal"},
		{name: "string oil", material: Oil, str: "oil"},
		{name: "string plastic", material: Plastic, str: "plastic"},
		{name: "string wood", material: Wood, str: "wood"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.material.String()
			exp := tc.str

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}

func TestMaterialDuration(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name		string
		material 	Material
		duration 	time.Duration
	}{
		{name: "duration metal", material: Metal, duration: 5000000},
		{name: "duration oil", material: Oil, duration: 3000000},
		{name: "duration plastic", material: Plastic, duration: 7000000},
		{name: "duration wood", material: Wood, duration: 4000000},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.material.Duration()
			exp := tc.duration

			if exp != act {
				t.Fatalf("expected %d, got %d", exp, act)
			}
		})
	}
}

func TestMaterialsDuration(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name		string
		mats 		Materials
		duration 	time.Duration
	}{
		{name: "materials duration simple", mats: Materials{Metal: 1}, duration: 5000000},
		{name: "materials duration many of one material", mats: Materials{Metal: 1250}, duration: 6250000000},
		{name: "materials duration one of all material", mats: Materials{Metal: 1, Oil: 1, Plastic: 1, Wood: 1}, duration: 19000000},
		{name: "materials duration many of all material", mats: Materials{Metal: 1250, Oil: 1250, Plastic: 1250, Wood: 1250}, duration: 23750000000},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.mats.Duration()
			exp := tc.duration

			if exp != act {
				t.Fatalf("expected %d, got %d", exp, act)
			}
		})
	}
}

func TestMaterialString(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name	string
		mats 	Materials
		str 	string
	}{
		{name: "materials string simple", mats: Materials{Metal: 1}, str: "[{metal:1x}]"},
		{name: "materials string many of one material", mats: Materials{Metal: 1250}, str: "[{metal:1250x}]"},
		{name: "materials string one of all material", mats: Materials{Metal: 1, Oil: 1, Plastic: 1, Wood: 1}, str: "[{metal:1x}, {oil:1x}, {plastic:1x}, {wood:1x}]"},
		{name: "materials string many of all material", mats: Materials{Metal: 1250, Oil: 1250, Plastic: 1250, Wood: 1250}, str: "[{metal:1250x}, {oil:1250x}, {plastic:1250x}, {wood:1250x}]"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.mats.String()
			exp := tc.str

			if exp != act {
				t.Fatalf("expected %s, got %s", exp, act)
			}
		})
	}
}