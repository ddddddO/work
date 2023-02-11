package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start stages")

	s := &stages{}
	s.start()

	fmt.Println("end stages")
}

type stages struct{}

func (s *stages) start() {
	s.fourth(s.third(s.second(s.first())))
}

func (s *stages) first() <-chan int {
	dest := make(chan int)

	go func() {
		defer close(dest)
		for i := 0; i < 10; i++ {
			dest <- i
		}
	}()

	return dest
}

// NOTE: pipeline & worker の動作確認
func (s *stages) second(src <-chan int) <-chan int {
	dest := make(chan int)

	go func() {
		defer close(dest)

		// Dispatcher
		workerNum := 10
		wg := &sync.WaitGroup{}
		for i := 0; i < workerNum; i++ {
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				for {
					select {
					case n, ok := <-src:
						if !ok {
							return
						}
						time.Sleep(1 * time.Second)
						dest <- n + 1
					}
				}
			}(wg)
		}
		wg.Wait()
	}()

	return dest
}

func (s *stages) third(src <-chan int) <-chan int {
	dest := make(chan int)

	go func() {
		defer close(dest)
		for {
			select {
			case n, ok := <-src:
				if !ok {
					return
				}
				dest <- n * n
			}
		}
	}()

	return dest
}

func (s *stages) fourth(src <-chan int) {
	for {
		select {
		case n, ok := <-src:
			if !ok {
				return
			}
			fmt.Println(n)
		}
	}
}
