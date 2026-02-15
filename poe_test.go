package poe

import (
	"math"
	"testing"
)

// helpers
func assertXsEqual(a, b X, t *testing.T) {
	for k, va := range a {
		if vb, ok := b[k]; !ok {
			t.Errorf(`b["%s"] does not exist`, k)
		} else if math.Abs(va-vb) > 1e-6 {
			t.Errorf(`a["%s"] != b["%s"]: %v != %v`, k, k, va, vb)
		}
	}

	for k := range b {
		if _, ok := a[k]; !ok {
			t.Errorf(`a["%s"] does not exist`, k)
		}
	}
}

// tests
func Test_XEmpty(t *testing.T) {
	// create Poe
	opt := New(func(_ X) float64 { return 0 })

	// assert
	assertXsEqual(opt.X(), X{}, t)
}

func Test_X(t *testing.T) {
	// create Poe
	opt := New(func(_ X) float64 { return 0 })
	opt.Trials = Trials{
		&Trial{Input: X{"a": 1, "b": 2}, Output: 3},
		&Trial{Input: X{"a": 2.5, "b": 1.5}, Output: 4},
	}

	// assert
	assertXsEqual(opt.X(), X{"a": 1, "b": 2}, t)
}

func Test_Minimize(t *testing.T) {
	// create Poe
	opt := New(func(x X) float64 { return x["a"] + x["b"] })
	opt.Trials = Trials{
		&Trial{Input: X{"a": 1.5, "b": 2.5}, Output: 4},
		&Trial{Input: X{"a": 1, "b": 1}, Output: 2},
	}

	// Minimize()
	space := NewSpace(Bounds{"a": {1, 3}, "b": {1, 3}})
	opt.Minimize(NewRandomSampler(space), 5)

	// assert
	if len(opt.Trials) != 7 {
		t.Errorf("len(opt.Trials) != 7: %d", len(opt.Trials))
	}
	assertXsEqual(opt.X(), X{"a": 1, "b": 1}, t)
}

func Test_Maximize(t *testing.T) {
	// create Poe
	opt := New(func(x X) float64 { return x["a"] + x["b"] })
	opt.Trials = Trials{
		&Trial{Input: X{"a": 1.5, "b": 2.5}, Output: 4},
		&Trial{Input: X{"a": 3, "b": 3}, Output: 6},
	}

	// Maximize()
	space := NewSpace(Bounds{"a": {1, 3}, "b": {1, 3}})
	opt.Maximize(NewRandomSampler(space), 5)

	// assert
	if len(opt.Trials) != 7 {
		t.Errorf("len(opt.Trials) != 7: %d", len(opt.Trials))
	}
	assertXsEqual(opt.X(), X{"a": 3, "b": 3}, t)
}
