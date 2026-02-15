package poe

import (
	"math"
	"testing"
)

func Test_normPDF(t *testing.T) {
	// normPDF()
	p := normPDF(1, 0, 1)

	// assert
	if math.Abs(p-0.2419707) > 1e-6 {
		t.Errorf("normPDF(1, 0, 1) != 0.2419707: %f", p)
	}
}
