package main

import (
	"fmt"
	// "os"
	"time"
	"runtime"

	// "github.com/ddddddO/gtree"
)

func main() {
	fmt.Println("show")

	// root := gtree.NewRoot("root")
	// if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
	// 	os.Exit(1)
	// }

	go func() {
		fmt.Println("r")
	}()

	go func() {
		fmt.Println("rr")
	}()

	go func() {
		fmt.Println("rrr")
	}()

	stacks := make([]byte, 4096*2)
	_ = runtime.Stack(stacks, true)

	time.Sleep(1 * time.Second)
	fmt.Println(string(stacks))
}