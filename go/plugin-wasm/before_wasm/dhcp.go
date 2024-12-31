package main

import (
	"fmt"

	"github.com/ddddddO/work/go/plugin-wasm/contract"
)

type DHCP struct{}

func (d DHCP) GetStr() string {
	return "xxxxxx"
}

func (d DHCP) Name(x string) string {
	return fmt.Sprintf("protocol name: %s", x)
}

func (d DHCP) Port() uint64 {
	return 68
}

func main() {
	contract.Register(DHCP{})
}
