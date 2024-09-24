package lib

import (
	"testing"
)

func TestGetNumRepresentation(t *testing.T) {
	type numbRepTest struct {
		arg1 rune
		exp  int
	}

	var numbRepTests = []numbRepTest{
		{'0', 0},
		{'1', 1},
		{'A', 10},
		{'-', 41},
	}
	for _, test := range numbRepTests {
		if output, _ := GetNumRepresentation(test.arg1); output != test.exp {
			t.Errorf("Output %q not equal to expected %q", output, test.exp)
		}
	}
}
