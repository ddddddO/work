package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 64)
	fmt.Println(b, len(b))
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 64

	bb := make([]byte, 8)
	fmt.Println(bb, len(bb))
	// [0 0 0 0 0 0 0 0] 8

	bb[1] = 0x0001
	fmt.Println(bb)
	// [0 1 0 0 0 0 0 0]

	bb[3] = 9
	fmt.Println(bb)
	// [0 1 0 9 0 0 0 0]

	bb[5] = 0x00ff
	fmt.Println(bb)
	// [0 1 0 9 0 255 0 0]

	// bb[7] = 256 // ./main.go:28:8: constant 256 overflows byte

	opCode := 0x08<<10 | 0x0006
	fmt.Println(opCode)
	// 8198

	fmt.Println(byte(opCode))
	// 6
}
