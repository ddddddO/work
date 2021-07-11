package main

import (
	"fmt"
)

func main() {
	var target float64 = 25

	const limit = 10
	var before, after float64

	before = target
	after = 1
	for i := 0; (i < 10) && (before != after); i++ {
		after := newton(before)
		fmt.Printf("A guess for square root is  %d\n", after)
		before = after
	}
	fmt.Printf("Square root is: %d\n", after)
}

// sqroot n+1 = sqroot n − (sqroot n * sqroot n − x) / (2 * sqroot n)
/*
	sqroot=1かつx=25
	sqroot n+1 = 1 − (1 * 1 − 25) / (2 * 1)
			   = 1 - (-24 / 2)
			   = 13

	sqroot=1かつx=13
	sqroot n+1 = 1 − (1 * 1 − 13) / (2 * 1)
			   = 1 - (-12 / 2)
			   = 7?
*/
func newton(x float64) float64 {
	const sqroot = 1
	return sqroot - (sqroot*sqroot-x)/(2*sqroot)
}
