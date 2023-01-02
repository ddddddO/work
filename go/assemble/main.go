package main

import "fmt"

func main() {
	a := 9
	b := 8
	c := add(a, b)
	fmt.Printf("Assembly!: %d\n", c)
}

func add(a, b int) int {
	return a + b
}
