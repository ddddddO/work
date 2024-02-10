package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"log"
	"strings"
)

func main() {
	s := []string{"hello", "world", "abcd"}
	for i, x := range backward(s) {
		fmt.Println(i, x)
	}
	// Output:
	// in backward
	// b: abcd
	// 2 abcd
	// b: world
	// 1 world
	// b: hello
	// 0 hello

	fmt.Println("----------------------------")

	for ss := range pull(backward(s)) {
		fmt.Println("XXXXX:", ss)
	}
	// Output:
	// in pull
	// in backward
	// b: abcd
	// 👺< 2 abcd true
	// XXXXX: abcd
	// b: world
	// 👺< 1 world true
	// XXXXX: world
	// b: hello
	// 👺< 0 hello true
	// XXXXX: hello
	// 👺< 0  false

	fmt.Println("----------------------------")

	// if err := pipelineInPanic(); err != nil {
	// 	log.Fatal(err)
	// }

	// Output:
	// START: pipeline
	// END: Hello + world: assembled in stageA: assembled in stageB: assembled in stageC
	// END:  + iterator!!: assembled in stageA: assembled in stageB: assembled in stageC
	// END: aaaa + bbbb: assembled in stageA: assembled in stageB: assembled in stageC
	// END: cccc + : assembled in stageA: assembled in stageB: assembled in stageC
	// 👺< recovered in stageA
	// 👺< recovered in stageB
	// panic: error!! from stageB [recovered]
	// 				panic: error!! from stageB
	//
	// goroutine 19 [running]:
	// main.pipeline.stageB.func4.1()
	// 				/home/ochi/github.com/ddddddO/work/go/cmd/iter/main.go:153 +0x79
	// panic({0x498f40?, 0xc0000961d0?})
	// 				/usr/local/go/src/runtime/panic.go:770 +0x132
	// main.pipeline.stageB.func4(0xc0000aa180)
	// 				/home/ochi/github.com/ddddddO/work/go/cmd/iter/main.go:165 +0x10e
	// iter.Pull[...].func1()
	// 				/usr/local/go/src/iter/iter.go:75 +0xf5
	// created by iter.Pull[...] in goroutine 1
	// 				/usr/local/go/src/iter/iter.go:63 +0x105
	// exit status 2

	fmt.Println("----------------------------")

	if err := pipeline(); err != nil {
		log.Fatal(err)
	}
	// Output:
	// START: pipeline
	// END: Hello + world: assembled in stageA: assembled in stageB: assembled in stageC
	// END:  + iterator!!: assembled in stageA: assembled in stageB: assembled in stageC
	// END: aaaa + bbbb: assembled in stageA: assembled in stageB: assembled in stageC
	// 2024/02/10 12:47:31 error!! from stageB
	// exit status 1
}

// https://go.dev/wiki/RangefuncExperiment
func backward[E any](s []E) func(func(int, E) bool) {
	return func(yield func(int, E) bool) {
		fmt.Println("in backward")

		for i := len(s) - 1; i >= 0; i-- {
			fmt.Println("b:", s[i])

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

func pipelineInPanic() (err error) {
	defer func() {
		e, ok := recover().(error)
		if ok {
			err = e
		}
	}()

	src :=
		`Hello
world

iterator!!
aaaa
bbbb
cccc`

	fmt.Println("START: pipelineInPanic")
	for s := range stageC(stageB(stageA(strings.NewReader(src)))) {
		fmt.Println("END:", s)
	}

	return
}

func stageA(r io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		defer func() {
			fmt.Println("👺< recovered in stageA")
			e, ok := recover().(error)
			if ok {
				panic(e)
			}
		}()

		sc := bufio.NewScanner(r)
		i := 0
		s := ""
		for sc.Scan() {
			i++
			s += sc.Text()
			if i%2 == 1 {
				s += " + "
				continue
			}

			s += ": assembled in stageA"
			if !yield(s) {
				return
			}
			s = ""
		}
		s += ": assembled in stageA"
		if !yield(s) {
			return
		}
		if err := sc.Err(); err != nil {
			panic(err)
		}
	}
}

func stageB(iteratorA iter.Seq[string]) iter.Seq[string] {
	return func(yield func(string) bool) {
		defer func() {
			fmt.Println("👺< recovered in stageB")
			e, ok := recover().(error)
			if ok {
				panic(e)
			}
		}()

		p, stop := iter.Pull(iteratorA)
		defer stop()

		i := 0
		for {
			// パイプラインの真ん中のステージで一定以上のiteratorの処理でpanicさせる
			// errを呼び出し元でキャッチする検証
			if i > 3 {
				panic(errors.New("error!! from stageB"))
			}

			s, ok := p()
			if !ok {
				return
			}
			ss := s + ": assembled in stageB"
			if !yield(ss) {
				return
			}
			i++
		}
	}
}

func stageC(iteratorB func(func(string) bool)) func(func(string) bool) {
	return func(yield func(string) bool) {
		defer func() {
			fmt.Println("👺< recovered in stageC")
			e, ok := recover().(error)
			if ok {
				panic(e)
			}
		}()

		p, stop := iter.Pull(iteratorB)
		defer stop()

		for {
			s, ok := p()
			if !ok {
				return
			}
			s += ": assembled in stageC"
			if !yield(s) {
				return
			}
		}
	}
}

func pipeline() error {
	src :=
		`Hello
world

iterator!!
aaaa
bbbb
cccc`

	fmt.Println("START: pipeline")
	for s, err := range stageC_(stageB_(stageA_(strings.NewReader(src)))) {
		if err != nil {
			return err
		}
		fmt.Println("END:", s)
	}

	return nil
}

func stageA_(r io.Reader) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		sc := bufio.NewScanner(r)
		i := 0
		s := ""
		for sc.Scan() {
			i++
			s += sc.Text()
			if i%2 == 1 {
				s += " + "
				continue
			}

			s += ": assembled in stageA"
			if !yield(s, nil) {
				return
			}
			s = ""
		}
		s += ": assembled in stageA"
		if !yield(s, nil) {
			return
		}
		if err := sc.Err(); err != nil {
			yield("", err)
			return
		}
	}
}

func stageB_(iteratorA iter.Seq2[string, error]) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		p, stop := iter.Pull2(iteratorA)
		defer stop()

		i := 0
		for {
			// パイプラインの真ん中のステージで一定以上のiteratorの処理でerrorを発生させる
			if i > 2 {
				yield("", errors.New("error!! from stageB"))
				return
			}

			s, err, ok := p()
			if !ok {
				return
			}
			if err != nil {
				yield("", err)
				return
			}
			ss := s + ": assembled in stageB"
			if !yield(ss, nil) {
				return
			}
			i++
		}
	}
}

func stageC_(iteratorB func(func(string, error) bool)) func(func(string, error) bool) {
	return func(yield func(string, error) bool) {
		p, stop := iter.Pull2(iteratorB)
		defer stop()

		for {
			s, err, ok := p()
			if !ok {
				return
			}
			if err != nil {
				yield("", err)
				return
			}
			s += ": assembled in stageC"
			if !yield(s, nil) {
				return
			}
		}
	}
}
