package main

import (
	"fmt"
)

func main() {
	fmt.Println(0x01)
	fmt.Println(0x06)
	fmt.Println(0x0f)
	fmt.Println(0xff)
	fmt.Println()

	fmt.Println(0x01 << 1)
	fmt.Println(0x01 << 10)
	fmt.Println(0x0006)
	fmt.Println(0x01<<10 | 0x0006)
	fmt.Println(0x01<<10 | 0x0007)
	fmt.Println(0x01<<10 | 0x0008)
	fmt.Println()

	fmt.Println(0x02 << 1)
	fmt.Println(0x02 << 10)
	// Output:
	// 1
	// 6
	// 15
	// 255
	//
	// 2
	// 1024
	// 6
	// 1030
	// 1031
	// 1032
	//
	// 4
	// 2048
}
