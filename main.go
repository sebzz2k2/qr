package main

import (
	"fmt"
	"github.com/sebzz2k2/qr/encoding"
)

func main() {
	str := "HELLO WORLD"
	encodedStr := encoding.Encode(&str)
	fmt.Println(encodedStr)
}
