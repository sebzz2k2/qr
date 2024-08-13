package main

import (
	"fmt"
	"log"
	"strconv"
)

func getMode() string {
	// only alphanumeric for now
	return fmt.Sprintf("%04s", strconv.FormatInt(int64(0b0010), 2))
}
func getQrVersion(val int) int {
	maxChar := [40]int{
		16, 29, 47, 67, 87, 108, 125, 157, 189, 221,
		259, 296, 352, 376, 426, 470, 531, 574, 644, 702,
		742, 823, 890, 963, 1041, 1094, 1172, 1263, 1322, 1429,
		1499, 1618, 1700, 1787, 1867, 1966, 2071, 2181, 2298, 2420,
	}

	for i := 0; i < len(maxChar); i++ {
		if maxChar[i] >= val {
			return i + 1
		}
	}
	return 0
}
func main() {
	str := "HELLO WORLD"
	qrVer := getQrVersion(len(str))
	if qrVer == 0 {
		log.Fatalf(`Sting is not encodeable`)
	}
	fmt.Println(qrVer)
	j := getMode()
	fmt.Println(j)
}
