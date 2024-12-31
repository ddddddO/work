package main

import (
	"fmt"

	"github.com/ddddddO/work/go/plugin-wasm/contract"
)

type IMAP struct{}

func (i IMAP) GetStr() string {
	return "wwwwwwww"
}

func (i IMAP) Name(x string) string {
	return fmt.Sprintf("protocol name: %s", "IMAP")
}

func (i IMAP) Port() uint64 {
	return 143
}

func main() {
	contract.Register(IMAP{})
}
