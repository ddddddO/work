// NOTE: こちらの挙動について確認
// https://github.com/hnakamur/httpcapt/issues/1
// structのフィールドがポインタか値かで変わるよう
package main

import (
	"fmt"
	"time"
)

type multiHandler struct {
	handlers []handler
}

func (mh *multiHandler) handle() {
	for _, h := range mh.handlers {
		go h.do()
	}
}

type multiPointaHandler struct {
	handlers []*handler
}

func (mph *multiPointaHandler) handle() {
	for _, h := range mph.handlers {
		go h.do()
	}
}

type handler struct {
	n int
}

func (h *handler) do() {
	fmt.Println(h.n)
}

func main() {
	a := handler{n: 1}
	b := handler{n: 2}
	c := handler{n: 3}

	mh := &multiHandler{handlers: []handler{a, b, c}}
	mh.handle()
	time.Sleep(1 * time.Second)

	fmt.Println("-------")

	mph := &multiPointaHandler{handlers: []*handler{&a, &b, &c}}
	mph.handle()
	time.Sleep(1 * time.Second)

	// output:
	// 3
	// 3
	// 3
	// -------
	// 3
	// 1
	// 2
}
