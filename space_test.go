package poe

import "testing"

// tests -> Domain
func Test_NewDomain(t *testing.T) {
	// NewDomain()
	domain := NewDomain([2]float64{5, 2})

	// assert
	if domain.Min != 2 {
		t.Errorf("Min != 2: %f", domain.Min)
	}
	if domain.Max != 5 {
		t.Errorf("Max != 5: %f", domain.Max)
	}
}

func Test_DomainClipIn(t *testing.T) {
	// NewDomain()
	domain := NewDomain([2]float64{5, 2})

	// assert
	v := domain.Clip(4)
	if v != 4 {
		t.Errorf("Clip(4) != 4: %f", v)
	}
}

func Test_DomainClipOutHigh(t *testing.T) {
	// NewDomain()
	domain := NewDomain([2]float64{5, 2})

	// assert
	v := domain.Clip(6)
	if v != 5 {
		t.Errorf("Clip(6) != 5: %f", v)
	}
}

func Test_DomainClipOutLow(t *testing.T) {
	// NewDomain()
	domain := NewDomain([2]float64{5, 2})

	// assert
	v := domain.Clip(1)
	if v != 2 {
		t.Errorf("Clip(1) != 2: %f", v)
	}
}
