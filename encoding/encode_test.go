package encoding

import (
	"fmt"
	"testing"
)

func TestGetModeIndicator(t *testing.T) {
	got := GetModeIndicator()
	want := "0010"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetNumRepresentation(t *testing.T) {
	type numbRepTest struct {
		arg1 rune
		exp  int
	}

	var numbRepTests = []numbRepTest{
		numbRepTest{'0', 0},
		numbRepTest{'1', 1},
		numbRepTest{'A', 10},
		numbRepTest{'-', 41},
	}
	for _, test := range numbRepTests {
		if output, _ := GetNumRepresentation(test.arg1); output != test.exp {
			t.Errorf("Output %q not equal to expected %q", output, test.exp)
		}
	}
}

func TestGetCharCountIndicator(t *testing.T) {
	tests := []struct {
		version   int
		charLen   int
		expected  string
		expectErr bool
	}{
		{version: 1, charLen: 5, expected: "000000101", expectErr: false},
		{version: 9, charLen: 123, expected: "001111011", expectErr: false},
		{version: 26, charLen: 2047, expected: "11111111111", expectErr: false},
		{version: 40, charLen: 8191, expected: "1111111111111", expectErr: false},
		{version: 0, charLen: 5, expected: "", expectErr: true},
		{version: 41, charLen: 5, expected: "", expectErr: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Version%d_CharLen%d", tt.version, tt.charLen), func(t *testing.T) {
			result, err := GetCharCountIndicator(tt.version, tt.charLen)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
			if result != tt.expected {
				t.Errorf("Expected result: %s, got: %s", tt.expected, result)
			}
		})
	}
}

func TestGetCharCapacityBits(t *testing.T) {
	tests := []struct {
		version   int
		expected  int
		expectErr bool
	}{
		{version: 1, expected: 13 * 8, expectErr: false},
		{version: 9, expected: 132 * 8, expectErr: false},
		{version: 20, expected: 485 * 8, expectErr: false},
		{version: 40, expected: 1666 * 8, expectErr: false},
		{version: 0, expected: 0, expectErr: true},
		{version: 41, expected: 0, expectErr: true},
	}

	for _, tt := range tests {
		t.Run(
			"Version_"+string(rune(tt.version)),
			func(t *testing.T) {
				result, err := GetCharCapacityBits(tt.version)
				if (err != nil) != tt.expectErr {
					t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				}
				if result != tt.expected {
					t.Errorf("Expected result: %d, got: %d", tt.expected, result)
				}
			})
	}
}

func TestGetTerminator(t *testing.T) {
	tests := []struct {
		diff     int
		expected string
	}{
		{diff: 0, expected: ""},
		{diff: 2, expected: "00"},
		{diff: 3, expected: "000"},
		{diff: 4, expected: "0000"},
		{diff: 5, expected: "0000"},
		{diff: 10, expected: "0000"},
	}

	for _, tt := range tests {
		t.Run(
			"Diff_"+string(rune(tt.diff)),
			func(t *testing.T) {
				result := GetTerminator(tt.diff)
				if result != tt.expected {
					t.Errorf("Expected result: %s, got: %s", tt.expected, result)
				}
			})
	}
}

func TestGenPadding(t *testing.T) {
	tests := []struct {
		length   int
		capacity int
		expected string
	}{
		{length: 80, capacity: 104, expected: "111011000001000111101100"},
	}

	for _, tt := range tests {
		t.Run(
			"Length_"+string(rune(tt.length))+"_Capacity_"+string(rune(tt.capacity)),
			func(t *testing.T) {
				result := GenPadding(tt.length, tt.capacity)
				if result != tt.expected {
					t.Errorf("Expected result: %s, got: %s", tt.expected, result)
				}
			})
	}
}

func TestGetQrVersion(t *testing.T) {
	tests := []struct {
		val       int
		expected  int
		expectErr bool
	}{
		{val: 16, expected: 1, expectErr: false},
		{val: 29, expected: 2, expectErr: false},
		{val: 47, expected: 3, expectErr: false},
		{val: 67, expected: 4, expectErr: false},
		{val: 150, expected: 8, expectErr: false},
		{val: 2500, expected: 0, expectErr: true},
		{val: 0, expected: 0, expectErr: true},
	}

	for _, tt := range tests {
		t.Run(
			"Val_"+string(rune(tt.val)),
			func(t *testing.T) {
				result, err := GetQrVersion(tt.val)
				if (err != nil) != tt.expectErr {
					t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				}
				if result != tt.expected {
					t.Errorf("Expected result: %d, got: %d", tt.expected, result)
				}
			})
	}
}
