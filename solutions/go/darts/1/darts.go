package darts

import "math"

func Score(x, y float64) int {
	vlen := math.Sqrt(x*x + y*y)
	if vlen <= 1 {
		return 10
	}
	if vlen <= 5 {
		return 5
	}
	if vlen <= 10 {
		return 1
	}
	return 0
}
