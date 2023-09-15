package main

import (
	"math"
)

const DELTA = 0.0000001
const InitialZ = 100.0

func Sqrt(x float64) (z float64) {
	z = InitialZ

	step := func() float64 {
		return z - (z*z-x)/(2*z)
	}

	for zz := step(); math.Abs(zz-z) > DELTA; {
		z = zz
		zz = step()
	}
	return
}

func main() {
}
