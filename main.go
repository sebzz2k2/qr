package main

import (
	"fmt"
	"github.com/sebzz2k2/qr/encoding"
	"log"
)

func GetCodeBlocks(str string) []string {
	codeBlock := []string{}
	for i := 0; i < len(str); i += 8 {
		codeBlock = append(codeBlock, string(str[i:i+8]))
	}
	return codeBlock
}

/*
group->block->codewords
*/
func main() {
	//str := "HELLO WORLD"
	str := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG AND THEN RUNS AWAY QUICKLY OUT OF SIGHT"
	qrVer, err := encoding.GetQrVersion(len(str))
	if err != nil {
		log.Fatalf(err.Error())
	}
	encodedStr := encoding.Encode(&str, qrVer)
	blocks := GetCodeBlocks(encodedStr)
	fmt.Println(encodedStr, blocks)
}
