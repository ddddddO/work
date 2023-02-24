package main

import (
	"context"
	"fmt"
	"time"
)

type multiHandler struct {
	handlers []handler
}

func (mh *multiHandler) handle(ctx context.Context, result <-chan string) {
	for _, h := range mh.handlers {
		go h.handle(ctx, result) // TODO: この状態で、各hとなるなめに何が必要なのか？
	}
}

type handler struct {
	n int
}

func (h *handler) handle(_ context.Context, result <-chan string) {
	<-result
	fmt.Println(h.n)
}

func main() {
	a := handler{n: 10}
	b := handler{n: 20}
	c := handler{n: 40}
	hs := []handler{a, b, c}
	mh := multiHandler{handlers: hs}

	ctx := context.Background()
	result := make(chan string)

	mh.handle(ctx, result)
	time.Sleep(1 * time.Second)
	result <- "aaa"
	result <- "aaa"
	result <- "aaa"

	time.Sleep(3 * time.Second)
	fmt.Println("end!")
	// 40
	// 40
	// 40
	// end!
}
