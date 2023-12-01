package laba4

import (
	"math"
)

func Task1(xn, xk, deltax float64) []float64 {
	var answersA []float64 = make([]float64, 0)
	for x := xn; x <= xk; xn += deltax {
		y := math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), 1/7)
		answersA = append(answersA, y)
	}
	return answersA
}

func Task2(znX []float64) []float64 {
	var answersB []float64 = make([]float64, 0)
	for _, value := range znX {
		x := value
		y := math.Pow(math.Pow(math.Asin(x), 4)+math.Pow(math.Acos(x), 6), 1/7)
		answersB = append(answersB, y)
	}
	return answersB
}
