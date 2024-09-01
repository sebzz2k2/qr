package main

import (
	"github.com/sebzz2k2/qr/encoding"
	"github.com/sebzz2k2/qr/errorCorrection"

	"log"
)

func GetCodeBlocks(str string) []string {
	codeBlock := []string{}
	for i := 0; i < len(str); i += 8 {
		codeBlock = append(codeBlock, string(str[i:i+8]))
	}
	return codeBlock
}

// GetGBC G -> group B -> block C -> codeword
func GetGBC(data string, ec *errorCorrection.QRCodeData) [][][]string {
	codeblocks := GetCodeBlocks(data)

	//g1Arr and g2Arr are initialized with the number of blocks and codewords per block as defined in ec.
	g1Arr := make([][]string, ec.Group1Blocks)
	for i := range g1Arr {
		g1Arr[i] = make([]string, ec.Group1DataCodewordsPerBlock)
	}

	g2Arr := make([][]string, ec.Group2Blocks)
	for i := range g2Arr {
		g2Arr[i] = make([]string, ec.Group2DataCodewordsPerBlock)
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
			g1Arr[row][col] = block
		} else {
			// Shift the index for g2Arr
			g2Index := blockIndex - len(g1Arr)*ec.Group1DataCodewordsPerBlock
			row := g2Index / ec.Group2DataCodewordsPerBlock
			col := g2Index % ec.Group2DataCodewordsPerBlock
			g2Arr[row][col] = block
		}
		blockIndex++
	}
	return [][][]string{g1Arr, g2Arr}
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

	GetGBC(encodedStr, ecVals)
	//fmt.Println( blocks)
}
