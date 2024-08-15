package utils

import (
	"errors"
	"fmt"
)

type QRCodeData struct {
	Version                     int
	TotalDataCodewords          int
	ECCodewordsPerBlock         int
	Group1Blocks                int
	Group1DataCodewordsPerBlock int
	Group2Blocks                int
	Group2DataCodewordsPerBlock int
}

func GetErrCorrVals(version int) (*QRCodeData, error) {
	var QRErrorCorrectionQ = []QRCodeData{
		{Version: 1, ECCodewordsPerBlock: 13, Group1Blocks: 1, Group1DataCodewordsPerBlock: 13, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 13},
		{Version: 2, ECCodewordsPerBlock: 22, Group1Blocks: 1, Group1DataCodewordsPerBlock: 22, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 22},
		{Version: 3, ECCodewordsPerBlock: 18, Group1Blocks: 2, Group1DataCodewordsPerBlock: 17, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 34},
		{Version: 4, ECCodewordsPerBlock: 26, Group1Blocks: 2, Group1DataCodewordsPerBlock: 24, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 48},
		{Version: 5, ECCodewordsPerBlock: 18, Group1Blocks: 2, Group1DataCodewordsPerBlock: 15, Group2Blocks: 2, Group2DataCodewordsPerBlock: 16, TotalDataCodewords: 62},
		{Version: 6, ECCodewordsPerBlock: 24, Group1Blocks: 4, Group1DataCodewordsPerBlock: 19, Group2Blocks: 0, Group2DataCodewordsPerBlock: 0, TotalDataCodewords: 76},
		{Version: 7, ECCodewordsPerBlock: 18, Group1Blocks: 2, Group1DataCodewordsPerBlock: 14, Group2Blocks: 4, Group2DataCodewordsPerBlock: 15, TotalDataCodewords: 88},
		{Version: 8, ECCodewordsPerBlock: 22, Group1Blocks: 4, Group1DataCodewordsPerBlock: 18, Group2Blocks: 2, Group2DataCodewordsPerBlock: 19, TotalDataCodewords: 110},
		{Version: 9, ECCodewordsPerBlock: 20, Group1Blocks: 4, Group1DataCodewordsPerBlock: 16, Group2Blocks: 4, Group2DataCodewordsPerBlock: 17, TotalDataCodewords: 132},
		{Version: 10, ECCodewordsPerBlock: 24, Group1Blocks: 6, Group1DataCodewordsPerBlock: 19, Group2Blocks: 2, Group2DataCodewordsPerBlock: 20, TotalDataCodewords: 154},
		{Version: 11, ECCodewordsPerBlock: 28, Group1Blocks: 4, Group1DataCodewordsPerBlock: 22, Group2Blocks: 4, Group2DataCodewordsPerBlock: 23, TotalDataCodewords: 180},
		{Version: 12, ECCodewordsPerBlock: 26, Group1Blocks: 4, Group1DataCodewordsPerBlock: 20, Group2Blocks: 6, Group2DataCodewordsPerBlock: 21, TotalDataCodewords: 206},
		{Version: 13, ECCodewordsPerBlock: 24, Group1Blocks: 8, Group1DataCodewordsPerBlock: 20, Group2Blocks: 4, Group2DataCodewordsPerBlock: 21, TotalDataCodewords: 244},
		{Version: 14, ECCodewordsPerBlock: 20, Group1Blocks: 11, Group1DataCodewordsPerBlock: 16, Group2Blocks: 5, Group2DataCodewordsPerBlock: 17, TotalDataCodewords: 261},
		{Version: 15, ECCodewordsPerBlock: 30, Group1Blocks: 5, Group1DataCodewordsPerBlock: 24, Group2Blocks: 7, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 295},
		{Version: 16, ECCodewordsPerBlock: 24, Group1Blocks: 15, Group1DataCodewordsPerBlock: 19, Group2Blocks: 2, Group2DataCodewordsPerBlock: 20, TotalDataCodewords: 325},
		{Version: 17, ECCodewordsPerBlock: 28, Group1Blocks: 1, Group1DataCodewordsPerBlock: 22, Group2Blocks: 15, Group2DataCodewordsPerBlock: 23, TotalDataCodewords: 367},
		{Version: 18, ECCodewordsPerBlock: 28, Group1Blocks: 17, Group1DataCodewordsPerBlock: 22, Group2Blocks: 1, Group2DataCodewordsPerBlock: 23, TotalDataCodewords: 397},
		{Version: 19, ECCodewordsPerBlock: 26, Group1Blocks: 17, Group1DataCodewordsPerBlock: 21, Group2Blocks: 4, Group2DataCodewordsPerBlock: 22, TotalDataCodewords: 445},
		{Version: 20, ECCodewordsPerBlock: 30, Group1Blocks: 15, Group1DataCodewordsPerBlock: 24, Group2Blocks: 5, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 485},
		{Version: 21, ECCodewordsPerBlock: 28, Group1Blocks: 17, Group1DataCodewordsPerBlock: 22, Group2Blocks: 6, Group2DataCodewordsPerBlock: 23, TotalDataCodewords: 512},
		{Version: 22, ECCodewordsPerBlock: 30, Group1Blocks: 7, Group1DataCodewordsPerBlock: 24, Group2Blocks: 16, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 568},
		{Version: 23, ECCodewordsPerBlock: 30, Group1Blocks: 11, Group1DataCodewordsPerBlock: 24, Group2Blocks: 14, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 614},
		{Version: 24, ECCodewordsPerBlock: 30, Group1Blocks: 11, Group1DataCodewordsPerBlock: 24, Group2Blocks: 16, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 664},
		{Version: 25, ECCodewordsPerBlock: 30, Group1Blocks: 7, Group1DataCodewordsPerBlock: 24, Group2Blocks: 22, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 718},
		{Version: 26, ECCodewordsPerBlock: 28, Group1Blocks: 28, Group1DataCodewordsPerBlock: 22, Group2Blocks: 6, Group2DataCodewordsPerBlock: 23, TotalDataCodewords: 754},
		{Version: 27, ECCodewordsPerBlock: 30, Group1Blocks: 8, Group1DataCodewordsPerBlock: 23, Group2Blocks: 26, Group2DataCodewordsPerBlock: 24, TotalDataCodewords: 808},
		{Version: 28, ECCodewordsPerBlock: 30, Group1Blocks: 4, Group1DataCodewordsPerBlock: 24, Group2Blocks: 31, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 871},
		{Version: 29, ECCodewordsPerBlock: 30, Group1Blocks: 1, Group1DataCodewordsPerBlock: 23, Group2Blocks: 37, Group2DataCodewordsPerBlock: 24, TotalDataCodewords: 911},
		{Version: 30, ECCodewordsPerBlock: 30, Group1Blocks: 15, Group1DataCodewordsPerBlock: 24, Group2Blocks: 25, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 985},
		{Version: 31, ECCodewordsPerBlock: 30, Group1Blocks: 42, Group1DataCodewordsPerBlock: 24, Group2Blocks: 1, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1033},
		{Version: 32, ECCodewordsPerBlock: 30, Group1Blocks: 10, Group1DataCodewordsPerBlock: 24, Group2Blocks: 35, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1115},
		{Version: 33, ECCodewordsPerBlock: 30, Group1Blocks: 29, Group1DataCodewordsPerBlock: 24, Group2Blocks: 19, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1171},
		{Version: 34, ECCodewordsPerBlock: 30, Group1Blocks: 44, Group1DataCodewordsPerBlock: 24, Group2Blocks: 7, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1231},
		{Version: 35, ECCodewordsPerBlock: 30, Group1Blocks: 39, Group1DataCodewordsPerBlock: 24, Group2Blocks: 14, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1286},
		{Version: 36, ECCodewordsPerBlock: 30, Group1Blocks: 46, Group1DataCodewordsPerBlock: 24, Group2Blocks: 10, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1354},
		{Version: 37, ECCodewordsPerBlock: 30, Group1Blocks: 49, Group1DataCodewordsPerBlock: 24, Group2Blocks: 10, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1426},
		{Version: 38, ECCodewordsPerBlock: 30, Group1Blocks: 48, Group1DataCodewordsPerBlock: 24, Group2Blocks: 14, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1502},
		{Version: 39, ECCodewordsPerBlock: 30, Group1Blocks: 43, Group1DataCodewordsPerBlock: 24, Group2Blocks: 22, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1582},
		{Version: 40, ECCodewordsPerBlock: 30, Group1Blocks: 34, Group1DataCodewordsPerBlock: 24, Group2Blocks: 34, Group2DataCodewordsPerBlock: 25, TotalDataCodewords: 1666},
	}
	if version >= 1 && version <= 40 {
		return &QRErrorCorrectionQ[version-1], nil
	}
	return nil, errors.New("invalid version")
}

func GetNumRepresentation(char rune) (int, error) {
	var CharacterMap = map[rune]int{
		'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'A': 10, 'B': 11, 'C': 12, 'D': 13, 'E': 14, 'F': 15, 'G': 16, 'H': 17, 'I': 18, 'J': 19,
		'K': 20, 'L': 21, 'M': 22, 'N': 23, 'O': 24, 'P': 25, 'Q': 26, 'R': 27, 'S': 28, 'T': 29,
		'U': 30, 'V': 31, 'W': 32, 'X': 33, 'Y': 34, 'Z': 35, ' ': 36, '$': 37, '%': 38,
		'*': 39, '+': 40, '-': 41, '.': 42, '/': 43, ':': 44,
	}
	char = rune(string(char)[0])
	if value, exists := CharacterMap[char]; exists {
		return value, nil
	}
	return -1, fmt.Errorf("character %c not found in the map", char)
}
