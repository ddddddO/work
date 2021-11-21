package main

import (
	"fmt"
)

type bit uint

const (
	zero bit = iota
	one
)

func main() {
	var in string
	fmt.Scan(&in)

	fmt.Println(in)
	for _, chr := range in {
		bits := DecimalToBitSlice(uint8(chr))
		fmt.Printf("%+v ", bits)
	}
	fmt.Println()
}

// 10進数 -> [8]bit
func DecimalToBitSlice(src uint8) [8]bit {
	dst := [8]bit{}
	for i := 7; i >= 0 || src >= 2; i-- {
		divided, remainder := calculate(src)
		dst[i] = toBit(remainder)
		src = divided
	}

	return dst
}

// 10進数から2進数へ計算するためのfunc
func calculate(num uint8) (uint8, uint8) {
	divided := num / 2
	remainder := num % 2
	return divided, remainder
}

func toBit(src uint8) bit {
	switch src {
	case 0:
		return zero
	case 1:
		return one
	default:
		panic("not bit")
	}
}
