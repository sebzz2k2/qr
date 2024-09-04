package main

import (
	"fmt"
	"github.com/sebzz2k2/qr/encoding"
	"github.com/sebzz2k2/qr/errorCorrection"
	"strconv"

	"log"
)

func GetCodeBlocks(str string) []string {
	codeBlock := []string{}
	for i := 0; i < len(str); i += 8 {
		codeBlock = append(codeBlock, string(str[i:i+8]))
	}
	return codeBlock
}

func getDecimal(data string) int {
	decimalValue, _ := strconv.ParseInt(data, 2, 64)
	intVal := int(decimalValue)
	return intVal
}

// GetGBC G -> group B -> block C -> codeword
func GetGBC(data string, ec *errorCorrection.QRCodeData) [][][]int {
	codeblocks := GetCodeBlocks(data)

	//g1Arr and g2Arr are initialized with the number of blocks and codewords per block as defined in ec.
	g1Arr := make([][]int, ec.Group1Blocks)
	for i := range g1Arr {
		g1Arr[i] = make([]int, ec.Group1DataCodewordsPerBlock)
	}

	g2Arr := make([][]int, ec.Group2Blocks)
	for i := range g2Arr {
		g2Arr[i] = make([]int, ec.Group2DataCodewordsPerBlock)
	}

	//	Iterate through codeblocks.
	//	For each block, check whether it belongs to g1Arr or g2Arr based on the current blockIndex.
	//	Compute the row and column within the array where the block should be placed.
	//	If blockIndex exceeds the total capacity of g1Arr, it starts filling g2Arr.
	blockIndex := 0
	for _, block := range codeblocks {
		if blockIndex < len(g1Arr)*ec.Group1DataCodewordsPerBlock {
			// Determine the target block and position in g1Arr
			row := blockIndex / ec.Group1DataCodewordsPerBlock
			col := blockIndex % ec.Group1DataCodewordsPerBlock
			blockInt := getDecimal(block)
			g1Arr[row][col] = blockInt
		} else {
			// Shift the index for g2Arr
			g2Index := blockIndex - len(g1Arr)*ec.Group1DataCodewordsPerBlock
			row := g2Index / ec.Group2DataCodewordsPerBlock
			col := g2Index % ec.Group2DataCodewordsPerBlock
			blockInt := getDecimal(block)
			g2Arr[row][col] = blockInt
		}
		blockIndex++
	}
	return [][][]int{g1Arr, g2Arr}
}

func main() {
	//str := "HELLO WORLD"
	str := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG AND THEN RUNS AWAY QUICKLY OUT OF SIGHT"
	qrVer, err := encoding.GetQrVersion(len(str))
	if err != nil {
		log.Fatalf(err.Error())
	}
	ecVals, err := errorCorrection.GetErrCorrVals(qrVer)
	if err != nil {
		log.Fatalf(err.Error())
	}
	encodedStr := encoding.Encode(&str, qrVer, ecVals.TotalDataCodewords*8)

	gbc := GetGBC(encodedStr, ecVals)
	fmt.Println(gbc)
}
