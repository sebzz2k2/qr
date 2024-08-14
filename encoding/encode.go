package encoding

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetQrVersion(val int) (int, error) {
	maxChar := [40]int{
		16, 29, 47, 67, 87, 108, 125, 157, 189, 221,
		259, 296, 352, 376, 426, 470, 531, 574, 644, 702,
		742, 823, 890, 963, 1041, 1094, 1172, 1263, 1322, 1429,
		1499, 1618, 1700, 1787, 1867, 1966, 2071, 2181, 2298, 2420,
	}

	for i := 0; i < len(maxChar); i++ {
		if maxChar[i] >= val {
			return i + 1, nil
		}
	}
	return 0, errors.New("Invalid version")
}

func GetModeIndicator() string {
	// only alphanumeric for now
	return fmt.Sprintf("%04s", strconv.FormatInt(int64(0b0010), 2))
}

func getNumRepresentation(char rune) (int, error) {
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
func GetEncodedDataStr(s string) string {
	pairValues := [][]int{}
	for i := 0; i < len(s); i += 2 {
		var pair string
		if i+1 < len(s) {
			pair = s[i : i+2]
		} else {
			pair = s[i : i+1]
		}
		values := []int{}
		for _, char := range pair {
			value, err := getNumRepresentation(char)
			if err != nil {
				fmt.Println(err)
			}
			values = append(values, value)
		}
		pairValues = append(pairValues, values)
	}
	str := ""
	for i := 0; i < len(pairValues); i++ {
		if len(pairValues[i]) > 1 {
			interVal := pairValues[i][0]*45 + pairValues[i][1]
			str += fmt.Sprintf("%011s", strconv.FormatInt(int64(interVal), 2))
		} else {
			str += fmt.Sprintf("%06s", strconv.FormatInt(int64(pairValues[i][0]), 2))
		}
	}
	return str
}

func GetCharCountIndicator(version int, charLen int) (string, error) {
	if version >= 1 && version <= 9 {
		return fmt.Sprintf("%09s", strconv.FormatInt(int64(charLen), 2)), nil
	}
	if version >= 10 && version <= 26 {
		return fmt.Sprintf("%011s", strconv.FormatInt(int64(charLen), 2)), nil
	}
	if version >= 27 && version <= 40 {
		return fmt.Sprintf("%013s", strconv.FormatInt(int64(charLen), 2)), nil
	}
	return "", errors.New("Invalid version")
}

func GetCharCapacityBits(version int) (int, error) {
	maxChar := [40]int{
		13, 22, 34, 48, 62, 76, 88, 110, 132, 154, 180,
		206, 244, 261, 295, 325, 367, 397, 445, 485, 512,
		568, 614, 664, 718, 754, 808, 871, 911, 985, 1033,
		1115, 1171, 1231, 1286, 1354, 1426, 1502, 1582, 1666}

	if version >= 1 && version <= 40 {
		return maxChar[version-1] * 8, nil
	}
	return 0, errors.New("Invalid version")
}

func GetTerminator(diff int) string {
	if diff < 4 {
		return GetZeroes(diff)
	} else {
		return GetZeroes(4)
	}
}

func GetZeroes(count int) string {
	return strings.Repeat("0", count)
}

func GenPadding(length int, capacity int) string {
	diff := capacity - length
	q := diff / 8
	extraBits := [2]string{"11101100", "00010001"}
	pad := ""
	for i := 0; i < q; i++ {
		pad += extraBits[i%2]
	}
	return pad
}
