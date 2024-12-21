package IntPow

import "math"

func Exponential(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Sqr(x int) int {
	return x * x
}

func Sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}
