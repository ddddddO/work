package main

import (
	"fmt"
)

type DHCP struct{}

//export Name
func (*DHCP) Name(x string) string {
	return fmt.Sprintf("protocol name: %s\n", x)
}

//export Port
func (*DHCP) Port() uint64 {
	return 68
}

func main() {}
