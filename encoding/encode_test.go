package encoding

import (
	"fmt"
	"log"
	"testing"
)

func TestGetModeIndicator(t *testing.T) {
	got := GetModeIndicator()
	want := "0010"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
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

func TestGetEncodedDataStr(t *testing.T) {
	input := "HELLO WORLD"
	expectedOutput := "01100001011" + "01111000110" + "10001011100" + "10110111000" + "10011010100" + "001101"
	result := GetEncodedDataStr(&input)
	if result != expectedOutput {
		t.Errorf("Expected %s but got %s", expectedOutput, result)
	}
}

func TestEncode(t *testing.T) {
	inp := "HELLO WORLD"
	got := Encode(&inp, 1)
	expected := "00100000010110110000101101111000110100010111001011011100010011010100001101000000111011000001000111101100"
	if got != expected {
		log.Fatalf("Expected: %s, got: %s", expected, got)
	}
}
