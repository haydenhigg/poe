package poe

import "testing"

func Test_CompareLessThan(t *testing.T) {
	// create Trials
	a := &Trial{Input: X{"a": 3, "b": 2}, Output: 3}
	b := &Trial{Input: X{"a": 3, "b": 2}, Output: 4}

	// Compare()
	comparison := a.Compare(b)

	// assert
	if comparison != -1 {
		t.Errorf("a.Compare(b) != -1: %d", comparison)
	}
}

func Test_CompareEqualTo(t *testing.T) {
	// create Trials
	a := &Trial{Input: X{"a": 3, "b": 2}, Output: 3}
	b := &Trial{Input: X{"a": 3, "b": 2}, Output: 3}

	// Compare()
	comparison := a.Compare(b)

	// assert
	if comparison != 0 {
		t.Errorf("a.Compare(b) != 0: %d", comparison)
	}
}

func Test_CompareGreaterThan(t *testing.T) {
	// create Trials
	a := &Trial{Input: X{"a": 3, "b": 2}, Output: 5}
	b := &Trial{Input: X{"a": 3, "b": 2}, Output: 4}

	// Compare()
	comparison := a.Compare(b)

	// assert
	if comparison != 1 {
		t.Errorf("a.Compare(b) != 1: %d", comparison)
	}
}
