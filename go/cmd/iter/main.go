package main

import (
	"fmt"
	"iter"
)

func main() {
	s := []string{"hello", "world", "abcd"}
	for i, x := range backward(s) {
    fmt.Println(i, x)
	}
	// Output:
	// in backward: abcd
	// 2 abcd
	// in backward: world
	// 1 world
	// in backward: hello
	// 0 hello

	fmt.Println("----------------------------")

	for ss := range pull(backward(s)) {
		fmt.Println("XXXXX:", ss)
	}
	// Output:
	// in pull
	// in backward: abcd
	// 👺< 2 abcd true
	// XXXXX: abcd
	// in backward: world
	// 👺< 1 world true
	// XXXXX: world
	// in backward: hello
	// 👺< 0 hello true
	// XXXXX: hello
	// 👺< 0  false
}

// https://go.dev/wiki/RangefuncExperiment
func backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
			for i := len(s)-1; i >= 0; i-- {
				fmt.Println("in backward:", s[i])

				if !yield(i, s[i]) {
						return
				}
			}
	}
}

func pull[E any](seq iter.Seq2[int, E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		fmt.Println("in pull")

		p, stop := iter.Pull2(seq)
		defer stop()

		for {
			i, s, ok := p()
			fmt.Println("👺<", i, s, ok)
			if !ok || !yield(s) {
				return
			}
		}
	}
}