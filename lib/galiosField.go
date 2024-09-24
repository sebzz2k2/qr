package lib

func GFAdd(x int, y int) int {
	return x ^ y
}
func GFPow(exp int) int {
	value := 1
	for i := 0; i < exp; i++ {
		value *= 2
		if value >= 256 {
			value ^= 285
		}
	}
	return value
}

func PowAdd(x int, y int) int {
	return (x + y) % 255
}
