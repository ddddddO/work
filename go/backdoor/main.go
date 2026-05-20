package main

import (
	"fmt"
	"time"

	_ "github.com/ddddddO/work/go/backdoor/silent"
)

func main() {
	fmt.Println("hello")
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
