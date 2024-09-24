package encoding

import (
	"errors"
	"fmt"
	"github.com/sebzz2k2/qr/lib"
	"log"
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
	if val == 0 {
		return 0, errors.New("invalid version")
	}
	for i := 0; i < len(maxChar); i++ {
		if maxChar[i] >= val {
			return i + 1, nil
		}
	}
	return 0, errors.New("invalid version")
}

func GetModeIndicator() string {
	// only alphanumeric for now
	return fmt.Sprintf("%04s", strconv.FormatInt(int64(0b0010), 2))
}

func GetEncodedDataStr(strPtr *string) string {
	var pairValues [][]int
	s := *strPtr
	for i := 0; i < len(s); i += 2 {
		var pair string
		if i+1 < len(s) {
			pair = s[i : i+2]
		} else {
			pair = s[i : i+1]
		}
		var values []int
		for _, char := range pair {
			value, err := lib.GetNumRepresentation(char)
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
	return "", errors.New("invalid version")
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

func Encode(s *string, qrVer int, charCapacity int) string {
	encodedStr := GetEncodedDataStr(s)

	charCount, err := GetCharCountIndicator(qrVer, len(*s))
	if err != nil {
		log.Fatalf(err.Error())
	}

	encodedStr = GetModeIndicator() + charCount + encodedStr + GetTerminator(charCapacity-len(encodedStr))
	encodedStr = encodedStr + GetZeroes(8-len(encodedStr)%8)
	encodedStr = encodedStr + GenPadding(len(encodedStr), charCapacity)
	return encodedStr
}
