package utils

import "testing"

func TestGetErrCorrVals(t *testing.T) {
	tests := []struct {
		version   int
		expected  *QRCodeData
		expectErr bool
	}{
		{version: 1, expected: &QRCodeData{Version: 1, ECCodewordsPerBlock: 13, Group1Blocks: 1, Group1DataCodewordsPerBlock: 13, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 13}, expectErr: false},
		{version: 2, expected: &QRCodeData{Version: 2, ECCodewordsPerBlock: 22, Group1Blocks: 1, Group1DataCodewordsPerBlock: 22, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 22}, expectErr: false},
		{version: 3, expected: &QRCodeData{Version: 3, ECCodewordsPerBlock: 18, Group1Blocks: 2, Group1DataCodewordsPerBlock: 17, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 34}, expectErr: false},
		{version: 40, expected: &QRCodeData{TotalDataCodewords: 1666}, expectErr: false},
		{version: 0, expected: nil, expectErr: true},
		{version: 41, expected: nil, expectErr: true},
	}

	for _, tt := range tests {
		t.Run(
			"Version_"+string(rune(tt.version)),
			func(t *testing.T) {
				result, err := GetErrCorrVals(tt.version)
				if (err != nil) != tt.expectErr {
					t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				}
				if !tt.expectErr && result.TotalDataCodewords*8 != tt.expected.TotalDataCodewords*8 {
					t.Errorf("Expected TotalDataCodewords: %d, got: %d", tt.expected.TotalDataCodewords*8, result.TotalDataCodewords*8)
				}
			})
	}
}
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
