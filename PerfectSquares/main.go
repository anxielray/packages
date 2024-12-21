package PerfectSquares

import "github.com/anxielray/packages/IntPow"

func PrevAft(n int) (int, int) {
	sqrt := IntPow.Sqrt(n)
	upperBound := IntPow.Sqr(sqrt + 1)
	lowerBound := IntPow.Sqr(sqrt - 1)
	return lowerBound, upperBound
}
