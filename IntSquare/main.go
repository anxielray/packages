package IntPow

import "math"

func sqr(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
