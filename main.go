package main

import (
	"fmt"
	"github.com/sebzz2k2/qr/encoding"
	"log"
)

func main() {
	str := "HELLO WORLD"
	qrVer, err := encoding.GetQrVersion(len(str))
	if err != nil {
		log.Fatalf(err.Error())
	}
	charCount, err := encoding.GetCharCountIndicator(qrVer, len(str))
	if err != nil {
		log.Fatalf(err.Error())
	}
	encodedStr := encoding.GetEncodedDataStr(str)
	charCapacity, err := encoding.GetCharCapacityBits(qrVer)
	if err != nil {
		log.Fatalf(err.Error())
	}
	encodedStr = encoding.GetModeIndicator() + charCount + encodedStr + encoding.GetTerminator(charCapacity-len(encodedStr))
	encodedStr = encodedStr + encoding.GetZeroes(8-len(encodedStr)%8)
	encodedStr = encodedStr + encoding.GenPadding(len(encodedStr), charCapacity)
	fmt.Println(encodedStr)
}
