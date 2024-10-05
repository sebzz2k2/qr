package main

import (
	"errors"
	"fmt"
	"github.com/sebzz2k2/qr/errorCorrection"
	"github.com/sebzz2k2/qr/lib"
	"strconv"
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

type GenPolyTerm struct {
	Alpha int
	X     int
}

func getGeneratorPolynomial(ecCodeWords int) ([]GenPolyTerm, error) {
	if ecCodeWords < 7 {
		return nil, errors.New("input should not be less than 7")
	}
	polynomials := [][]GenPolyTerm{
		{
			{Alpha: 0, X: 7},
			{Alpha: 87, X: 6},
			{Alpha: 229, X: 5},
			{Alpha: 146, X: 4},
			{Alpha: 149, X: 3},
			{Alpha: 238, X: 2},
			{Alpha: 102, X: 1},
			{Alpha: 21, X: 0},
		}, {
			{Alpha: 0, X: 8},
			{Alpha: 175, X: 7},
			{Alpha: 238, X: 6},
			{Alpha: 208, X: 5},
			{Alpha: 249, X: 4},
			{Alpha: 215, X: 3},
			{Alpha: 252, X: 2},
			{Alpha: 196, X: 1},
			{Alpha: 28, X: 0},
		}, {
			{Alpha: 0, X: 9},
			{Alpha: 95, X: 8},
			{Alpha: 246, X: 7},
			{Alpha: 137, X: 6},
			{Alpha: 231, X: 5},
			{Alpha: 235, X: 4},
			{Alpha: 149, X: 3},
			{Alpha: 11, X: 2},
			{Alpha: 123, X: 1},
			{Alpha: 36, X: 0},
		}, {
			{Alpha: 0, X: 10},
			{Alpha: 251, X: 9},
			{Alpha: 67, X: 8},
			{Alpha: 46, X: 7},
			{Alpha: 61, X: 6},
			{Alpha: 118, X: 5},
			{Alpha: 70, X: 4},
			{Alpha: 64, X: 3},
			{Alpha: 94, X: 2},
			{Alpha: 32, X: 1},
			{Alpha: 45, X: 0},
		}, {
			{Alpha: 0, X: 11},
			{Alpha: 220, X: 10},
			{Alpha: 192, X: 9},
			{Alpha: 91, X: 8},
			{Alpha: 194, X: 7},
			{Alpha: 172, X: 6},
			{Alpha: 177, X: 5},
			{Alpha: 209, X: 4},
			{Alpha: 116, X: 3},
			{Alpha: 227, X: 2},
			{Alpha: 10, X: 1},
			{Alpha: 55, X: 0},
		}, {
			{Alpha: 0, X: 12},
			{Alpha: 102, X: 11},
			{Alpha: 43, X: 10},
			{Alpha: 98, X: 9},
			{Alpha: 121, X: 8},
			{Alpha: 187, X: 7},
			{Alpha: 113, X: 6},
			{Alpha: 198, X: 5},
			{Alpha: 143, X: 4},
			{Alpha: 131, X: 3},
			{Alpha: 87, X: 2},
			{Alpha: 157, X: 1},
			{Alpha: 66, X: 0},
		}, {
			{Alpha: 0, X: 13},
			{Alpha: 74, X: 12},
			{Alpha: 152, X: 11},
			{Alpha: 176, X: 10},
			{Alpha: 100, X: 9},
			{Alpha: 86, X: 8},
			{Alpha: 100, X: 7},
			{Alpha: 106, X: 6},
			{Alpha: 104, X: 5},
			{Alpha: 130, X: 4},
			{Alpha: 218, X: 3},
			{Alpha: 206, X: 2},
			{Alpha: 140, X: 1},
			{Alpha: 78, X: 0},
		}, {
			{Alpha: 0, X: 14},
			{Alpha: 199, X: 13},
			{Alpha: 249, X: 12},
			{Alpha: 155, X: 11},
			{Alpha: 48, X: 10},
			{Alpha: 190, X: 9},
			{Alpha: 124, X: 8},
			{Alpha: 218, X: 7},
			{Alpha: 137, X: 6},
			{Alpha: 216, X: 5},
			{Alpha: 87, X: 4},
			{Alpha: 207, X: 3},
			{Alpha: 59, X: 2},
			{Alpha: 22, X: 1},
			{Alpha: 91, X: 0},
		}, {
			{Alpha: 0, X: 15},
			{Alpha: 8, X: 14},
			{Alpha: 183, X: 13},
			{Alpha: 61, X: 12},
			{Alpha: 91, X: 11},
			{Alpha: 202, X: 10},
			{Alpha: 37, X: 9},
			{Alpha: 51, X: 8},
			{Alpha: 58, X: 7},
			{Alpha: 58, X: 6},
			{Alpha: 237, X: 5},
			{Alpha: 140, X: 4},
			{Alpha: 124, X: 3},
			{Alpha: 5, X: 2},
			{Alpha: 99, X: 1},
			{Alpha: 105, X: 0},
		}, {
			{Alpha: 0, X: 16},
			{Alpha: 120, X: 15},
			{Alpha: 104, X: 14},
			{Alpha: 107, X: 13},
			{Alpha: 109, X: 12},
			{Alpha: 102, X: 11},
			{Alpha: 161, X: 10},
			{Alpha: 76, X: 9},
			{Alpha: 3, X: 8},
			{Alpha: 91, X: 7},
			{Alpha: 191, X: 6},
			{Alpha: 147, X: 5},
			{Alpha: 169, X: 4},
			{Alpha: 182, X: 3},
			{Alpha: 194, X: 2},
			{Alpha: 225, X: 1},
			{Alpha: 120, X: 0},
		}, {
			{Alpha: 0, X: 17},
			{Alpha: 43, X: 16},
			{Alpha: 139, X: 15},
			{Alpha: 206, X: 14},
			{Alpha: 78, X: 13},
			{Alpha: 43, X: 12},
			{Alpha: 239, X: 11},
			{Alpha: 123, X: 10},
			{Alpha: 206, X: 9},
			{Alpha: 214, X: 8},
			{Alpha: 147, X: 7},
			{Alpha: 24, X: 6},
			{Alpha: 99, X: 5},
			{Alpha: 150, X: 4},
			{Alpha: 39, X: 3},
			{Alpha: 243, X: 2},
			{Alpha: 163, X: 1},
			{Alpha: 136, X: 0},
		}, {
			{Alpha: 0, X: 18},
			{Alpha: 215, X: 17},
			{Alpha: 234, X: 16},
			{Alpha: 158, X: 15},
			{Alpha: 94, X: 14},
			{Alpha: 184, X: 13},
			{Alpha: 97, X: 12},
			{Alpha: 118, X: 11},
			{Alpha: 170, X: 10},
			{Alpha: 79, X: 9},
			{Alpha: 187, X: 8},
			{Alpha: 152, X: 7},
			{Alpha: 148, X: 6},
			{Alpha: 252, X: 5},
			{Alpha: 179, X: 4},
			{Alpha: 5, X: 3},
			{Alpha: 98, X: 2},
			{Alpha: 96, X: 1},
			{Alpha: 153, X: 0},
		}, {
			{Alpha: 0, X: 19},
			{Alpha: 67, X: 18},
			{Alpha: 3, X: 17},
			{Alpha: 105, X: 16},
			{Alpha: 153, X: 15},
			{Alpha: 52, X: 14},
			{Alpha: 90, X: 13},
			{Alpha: 83, X: 12},
			{Alpha: 17, X: 11},
			{Alpha: 150, X: 10},
			{Alpha: 159, X: 9},
			{Alpha: 44, X: 8},
			{Alpha: 128, X: 7},
			{Alpha: 153, X: 6},
			{Alpha: 133, X: 5},
			{Alpha: 252, X: 4},
			{Alpha: 222, X: 3},
			{Alpha: 138, X: 2},
			{Alpha: 220, X: 1},
			{Alpha: 171, X: 0},
		}, {
			{Alpha: 0, X: 20},
			{Alpha: 17, X: 19},
			{Alpha: 60, X: 18},
			{Alpha: 79, X: 17},
			{Alpha: 50, X: 16},
			{Alpha: 61, X: 15},
			{Alpha: 163, X: 14},
			{Alpha: 26, X: 13},
			{Alpha: 187, X: 12},
			{Alpha: 202, X: 11},
			{Alpha: 180, X: 10},
			{Alpha: 221, X: 9},
			{Alpha: 225, X: 8},
			{Alpha: 83, X: 7},
			{Alpha: 239, X: 6},
			{Alpha: 156, X: 5},
			{Alpha: 164, X: 4},
			{Alpha: 212, X: 3},
			{Alpha: 212, X: 2},
			{Alpha: 188, X: 1},
			{Alpha: 190, X: 0},
		},
		{
			{Alpha: 0, X: 21}, {Alpha: 240, X: 20}, {Alpha: 233, X: 19}, {Alpha: 104, X: 18},
			{Alpha: 247, X: 17}, {Alpha: 181, X: 16}, {Alpha: 140, X: 15}, {Alpha: 67, X: 14},
			{Alpha: 98, X: 13}, {Alpha: 85, X: 12}, {Alpha: 200, X: 11}, {Alpha: 210, X: 10},
			{Alpha: 115, X: 9}, {Alpha: 148, X: 8}, {Alpha: 137, X: 7}, {Alpha: 230, X: 6},
			{Alpha: 36, X: 5}, {Alpha: 122, X: 4}, {Alpha: 254, X: 3}, {Alpha: 148, X: 2},
			{Alpha: 175, X: 1}, {Alpha: 210, X: 0},
		},
		{
			{Alpha: 0, X: 22}, {Alpha: 210, X: 21}, {Alpha: 171, X: 20}, {Alpha: 247, X: 19},
			{Alpha: 242, X: 18}, {Alpha: 93, X: 17}, {Alpha: 230, X: 16}, {Alpha: 14, X: 15},
			{Alpha: 109, X: 14}, {Alpha: 221, X: 13}, {Alpha: 53, X: 12}, {Alpha: 200, X: 11},
			{Alpha: 74, X: 10}, {Alpha: 8, X: 9}, {Alpha: 172, X: 8}, {Alpha: 98, X: 7},
			{Alpha: 80, X: 6}, {Alpha: 219, X: 5}, {Alpha: 134, X: 4}, {Alpha: 160, X: 3},
			{Alpha: 105, X: 2}, {Alpha: 165, X: 1}, {Alpha: 231, X: 0},
		},
		{
			{Alpha: 0, X: 23}, {Alpha: 171, X: 22}, {Alpha: 102, X: 21}, {Alpha: 146, X: 20},
			{Alpha: 91, X: 19}, {Alpha: 49, X: 18}, {Alpha: 103, X: 17}, {Alpha: 65, X: 16},
			{Alpha: 17, X: 15}, {Alpha: 193, X: 14}, {Alpha: 150, X: 13}, {Alpha: 14, X: 12},
			{Alpha: 25, X: 11}, {Alpha: 183, X: 10}, {Alpha: 248, X: 9}, {Alpha: 94, X: 8},
			{Alpha: 164, X: 7}, {Alpha: 224, X: 6}, {Alpha: 192, X: 5}, {Alpha: 1, X: 4},
			{Alpha: 78, X: 3}, {Alpha: 56, X: 2}, {Alpha: 147, X: 1}, {Alpha: 253, X: 0},
		},
		{
			{Alpha: 0, X: 24}, {Alpha: 229, X: 23}, {Alpha: 121, X: 22}, {Alpha: 135, X: 21},
			{Alpha: 48, X: 20}, {Alpha: 211, X: 19}, {Alpha: 117, X: 18}, {Alpha: 251, X: 17},
			{Alpha: 126, X: 16}, {Alpha: 159, X: 15}, {Alpha: 180, X: 14}, {Alpha: 169, X: 13},
			{Alpha: 152, X: 12}, {Alpha: 192, X: 11}, {Alpha: 226, X: 10}, {Alpha: 228, X: 9},
			{Alpha: 218, X: 8}, {Alpha: 111, X: 7}, {Alpha: 0, X: 6}, {Alpha: 117, X: 5},
			{Alpha: 232, X: 4}, {Alpha: 87, X: 3}, {Alpha: 96, X: 2}, {Alpha: 227, X: 1},
			{Alpha: 21, X: 0},
		},
		{
			{Alpha: 0, X: 25}, {Alpha: 231, X: 24}, {Alpha: 181, X: 23}, {Alpha: 156, X: 22},
			{Alpha: 39, X: 21}, {Alpha: 170, X: 20}, {Alpha: 26, X: 19}, {Alpha: 12, X: 18},
			{Alpha: 59, X: 17}, {Alpha: 15, X: 16}, {Alpha: 148, X: 15}, {Alpha: 201, X: 14},
			{Alpha: 54, X: 13}, {Alpha: 66, X: 12}, {Alpha: 237, X: 11}, {Alpha: 208, X: 10},
			{Alpha: 99, X: 9}, {Alpha: 167, X: 8}, {Alpha: 144, X: 7}, {Alpha: 182, X: 6},
			{Alpha: 95, X: 5}, {Alpha: 243, X: 4}, {Alpha: 129, X: 3}, {Alpha: 178, X: 2},
			{Alpha: 252, X: 1}, {Alpha: 45, X: 0},
		},
		{
			{Alpha: 0, X: 26}, {Alpha: 173, X: 25}, {Alpha: 125, X: 24}, {Alpha: 158, X: 23},
			{Alpha: 2, X: 22}, {Alpha: 103, X: 21}, {Alpha: 182, X: 20}, {Alpha: 118, X: 19},
			{Alpha: 17, X: 18}, {Alpha: 145, X: 17}, {Alpha: 201, X: 16}, {Alpha: 111, X: 15},
			{Alpha: 28, X: 14}, {Alpha: 165, X: 13}, {Alpha: 53, X: 12}, {Alpha: 161, X: 11},
			{Alpha: 21, X: 10}, {Alpha: 245, X: 9}, {Alpha: 142, X: 8}, {Alpha: 13, X: 7},
			{Alpha: 102, X: 6}, {Alpha: 48, X: 5}, {Alpha: 227, X: 4}, {Alpha: 153, X: 3},
			{Alpha: 145, X: 2}, {Alpha: 218, X: 1}, {Alpha: 70, X: 0},
		},
		{
			{Alpha: 0, X: 27}, {Alpha: 79, X: 26}, {Alpha: 228, X: 25}, {Alpha: 8, X: 24},
			{Alpha: 165, X: 23}, {Alpha: 227, X: 22}, {Alpha: 21, X: 21}, {Alpha: 180, X: 20},
			{Alpha: 29, X: 19}, {Alpha: 9, X: 18}, {Alpha: 237, X: 17}, {Alpha: 70, X: 16},
			{Alpha: 99, X: 15}, {Alpha: 45, X: 14}, {Alpha: 58, X: 13}, {Alpha: 138, X: 12},
			{Alpha: 135, X: 11}, {Alpha: 73, X: 10}, {Alpha: 126, X: 9}, {Alpha: 172, X: 8},
			{Alpha: 94, X: 7}, {Alpha: 216, X: 6}, {Alpha: 193, X: 5}, {Alpha: 157, X: 4},
			{Alpha: 26, X: 3}, {Alpha: 17, X: 2}, {Alpha: 149, X: 1}, {Alpha: 96, X: 0},
		},
		{
			{Alpha: 0, X: 28}, {Alpha: 168, X: 27}, {Alpha: 223, X: 26}, {Alpha: 200, X: 25},
			{Alpha: 104, X: 24}, {Alpha: 224, X: 23}, {Alpha: 234, X: 22}, {Alpha: 108, X: 21},
			{Alpha: 180, X: 20}, {Alpha: 110, X: 19}, {Alpha: 190, X: 18}, {Alpha: 195, X: 17},
			{Alpha: 147, X: 16}, {Alpha: 205, X: 15}, {Alpha: 27, X: 14}, {Alpha: 232, X: 13},
			{Alpha: 201, X: 12}, {Alpha: 21, X: 11}, {Alpha: 43, X: 10}, {Alpha: 245, X: 9},
			{Alpha: 87, X: 8}, {Alpha: 42, X: 7}, {Alpha: 195, X: 6}, {Alpha: 212, X: 5},
			{Alpha: 119, X: 4}, {Alpha: 242, X: 3}, {Alpha: 37, X: 2}, {Alpha: 9, X: 1},
			{Alpha: 123, X: 0},
		},
		{
			{Alpha: 0, X: 29}, {Alpha: 156, X: 28}, {Alpha: 45, X: 27}, {Alpha: 183, X: 26},
			{Alpha: 29, X: 25}, {Alpha: 151, X: 24}, {Alpha: 219, X: 23}, {Alpha: 54, X: 22},
			{Alpha: 96, X: 21}, {Alpha: 249, X: 20}, {Alpha: 24, X: 19}, {Alpha: 136, X: 18},
			{Alpha: 5, X: 17}, {Alpha: 241, X: 16}, {Alpha: 175, X: 15}, {Alpha: 189, X: 14},
			{Alpha: 28, X: 13}, {Alpha: 75, X: 12}, {Alpha: 234, X: 11}, {Alpha: 150, X: 10},
			{Alpha: 148, X: 9}, {Alpha: 23, X: 8}, {Alpha: 9, X: 7}, {Alpha: 202, X: 6},
			{Alpha: 162, X: 5}, {Alpha: 68, X: 4}, {Alpha: 250, X: 3}, {Alpha: 140, X: 2},
			{Alpha: 24, X: 1}, {Alpha: 151, X: 0},
		},
		{
			{Alpha: 0, X: 30}, {Alpha: 41, X: 29}, {Alpha: 173, X: 28}, {Alpha: 145, X: 27},
			{Alpha: 152, X: 26}, {Alpha: 216, X: 25}, {Alpha: 31, X: 24}, {Alpha: 179, X: 23},
			{Alpha: 182, X: 22}, {Alpha: 50, X: 21}, {Alpha: 48, X: 20}, {Alpha: 110, X: 19},
			{Alpha: 86, X: 18}, {Alpha: 239, X: 17}, {Alpha: 96, X: 16}, {Alpha: 222, X: 15},
			{Alpha: 125, X: 14}, {Alpha: 42, X: 13}, {Alpha: 173, X: 12}, {Alpha: 226, X: 11},
			{Alpha: 193, X: 10}, {Alpha: 224, X: 9}, {Alpha: 130, X: 8}, {Alpha: 156, X: 7},
			{Alpha: 37, X: 6}, {Alpha: 251, X: 5}, {Alpha: 216, X: 4}, {Alpha: 238, X: 3},
			{Alpha: 40, X: 2}, {Alpha: 192, X: 1}, {Alpha: 180, X: 0},
		},
	}
	return polynomials[ecCodeWords-7], nil
}

type MsgPolyTerm struct {
	CoEff int
	X     int
}

func getMsgPolynomial(arr []int) []MsgPolyTerm {
	arrLen := len(arr)
	var msgPoly []MsgPolyTerm
	for i := 0; i < arrLen; i++ {
		msgPoly = append(msgPoly, MsgPolyTerm{
			CoEff: arr[i],
			X:     arrLen - (1 + i),
		})
	}
	return msgPoly
}

func main() {
	//str := "HELLO WORLD"
	//str := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG AND THEN RUNS AWAY QUICKLY OUT OF SIGHT"
	//qrVer, err := encoding.GetQrVersion(len(str))
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//ecVals, err := errorCorrection.GetErrCorrVals(qrVer)
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}
	//encodedStr := encoding.Encode(&str, qrVer, ecVals.TotalDataCodewords*8)
	//
	//gbc := GetGBC(encodedStr, ecVals)
	//fmt.Println(gbc)
	//
	//for _, groups := range gbc {
	//	for i, blocks := range groups {
	//		msgPoly := getMsgPolynomial(blocks)
	//		fmt.Println(msgPoly, i)
	//	}
	//}

	q := lib.LongDivide([]lib.Term{{3, 2}, {1, 1}, {-1, 0}}, []lib.Term{{1, 1}, {1, 0}})
	fmt.Println(q)
}
