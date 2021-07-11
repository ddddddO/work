package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	return x
}

func main() {
	start := time.Now()

	cnt := 15
	ch := make(chan string, cnt-1)
	for i := 1; i < cnt; i++ {
		go func(i int) {
			n := fib(float64(i))
			ch <- fmt.Sprintf("Fib(%v): %v\n", i, n)
		}(i)

	}

	for i := 1; i < cnt; i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
