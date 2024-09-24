package lib

import (
	"fmt"
	"math"
)

func PowDivide(exp1 int, exp2 int) int {
	return int(math.Abs(float64(exp1 - exp2)))
}

type Term struct {
	CoEff int
	Pow   int
}

func LongDivide(dividend []Term, divisor []Term) {
	i := 0
	var q []Term
	for i < len(dividend) && i < len(divisor) {
		if divisor[i].CoEff != 0 && dividend[i].CoEff != 0 {
			q = append(q, Term{
				CoEff: dividend[i].CoEff / divisor[i].CoEff,
				Pow:   PowDivide(dividend[i].Pow, divisor[i].Pow),
			})
		}
		i++

	}
	fmt.Println(q)
}

//
//import (
//"fmt"
//"math"
//)
//
//type Term struct {
//	CoEff int
//	Pow   int
//}
//
//func PowDivide(exp1 int, exp2 int) int {
//	return int(math.Abs(float64(exp1 - exp2)))
//}
//
//func longDivide(dividend []Term, divisor []Term) {
//	fmt.Println(dividend, divisor)
//
//}
//
//func main() {
//	dividend := []Term{{3, 2}, {1, 1}, {1, 0}}
//	divisor := []Term{{1, 1}, {1, 0}}
//
//	longDivide(dividend, divisor)
//}
