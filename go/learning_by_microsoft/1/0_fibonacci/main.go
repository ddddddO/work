package main

import (
	"fmt"
)

func main() {
	var val int
	fmt.Scanf("%d", &val)

	var result []int
	if val < 2 {
		panic(result)
	}

	result = append([]int{1, 1})
	cnt := val - 2
	for i := 0; i < cnt; i++ {
		lastIdx := len(result) - 1
		sum := fib(result[lastIdx-1], result[lastIdx])
		result = append(result, sum)
	}

	fmt.Println(result)
}

// main内の処理を関数に入れるんだろうけど、、
func fib(a, b int) int {
	return a + b
}
