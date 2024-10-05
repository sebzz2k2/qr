package lib

func PowDivide(exp1 int, exp2 int) int {
	return exp1 - exp2
}

func PowMultiply(exp1 int, exp2 int) int {
	return exp1 + exp2
}

type Term struct {
	CoEff int
	Pow   int
}

func TermMultiply(term1 Term, term2 Term) Term {
	return Term{
		CoEff: term1.CoEff * term2.CoEff,
		Pow:   PowMultiply(term1.Pow, term2.Pow),
	}
}

func TermDivide(term1 Term, term2 Term) Term {
	return Term{
		CoEff: term1.CoEff / term2.CoEff,
		Pow:   PowDivide(term1.Pow, term2.Pow),
	}
}

func TermSub(term1 Term, term2 Term) Term {
	return Term{
		CoEff: term1.CoEff - term2.CoEff,
		Pow:   term1.Pow,
	}
}
func LongDivide(dividend []Term, divisor []Term) []Term {
	var quotient []Term
	for len(dividend) > 0 && len(dividend) >= len(divisor) {
		quotientTerm := TermDivide(dividend[0], divisor[0])
		quotient = append(quotient, quotientTerm)

		// quotientTerm * divisor
		var qTermDivisor []Term
		for _, divisorTerm := range divisor {
			qTermDivisor = append(qTermDivisor, TermMultiply(divisorTerm, quotientTerm))
		}

		var newDividend []Term
		for k := 0; k < len(dividend); k++ {
			if k >= len(qTermDivisor) {
				newDividend = append(newDividend, dividend[k])
				continue
			}
			if dividend[k].CoEff != qTermDivisor[k].CoEff {
				newDividend = append(newDividend, TermSub(dividend[k], qTermDivisor[k]))
			}
		}
		dividend = newDividend
	}
	return quotient
}
