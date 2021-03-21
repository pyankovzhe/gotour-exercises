package main

import (
	"fmt"
	"math"
)

const Eps = 1e-6

func Sqrt(x float64) float64 {
	z := x / 2
	prev := 0.0

	for {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("prev: %v, z: %v\n", prev, z)

		if math.Abs(prev-z) < Eps {
			return prev
		} else {
			prev = z
		}
	}
}

func main() {
	fmt.Printf("result of math.Sqrt: %v\n", math.Sqrt(2))
	fmt.Printf("result of custom Sqrt func: %v\n", Sqrt(2))
}
