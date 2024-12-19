package IntPow

import "math"

func Sqr(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
