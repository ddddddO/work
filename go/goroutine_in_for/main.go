// NOTE: こちらの挙動について確認
// https://github.com/hnakamur/httpcapt/issues/1
// どうやら、interface越しに呼ばれるとfor文内で変数に束縛せずにgoroutineで処理できるよう
package main

import (
	"context"
	"fmt"
	"time"
)

type ifHandler interface {
	handle(ctx context.Context, result <-chan string)
}

func run(ifh ifHandler, ctx context.Context, result <-chan string) {
	ifh.handle(ctx, result)
}

type multiHandler struct {
	handlers []*handler
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
	time.Sleep(time.Duration(h.n) * time.Second)
	fmt.Println(h.n)
}

func main() {
	a := &handler{n: 1}
	b := &handler{n: 2}
	c := &handler{n: 3}
	hs := []*handler{a, b, c}
	mh := &multiHandler{handlers: hs}

	ctx := context.Background()
	result := make(chan string)

	run(mh, ctx, result)
	result <- "aaa"
	result <- "aaa"
	result <- "aaa"

	time.Sleep(4 * time.Second)
	fmt.Println("end!")
	// 1
	// 2
	// 3
	// end!
}
