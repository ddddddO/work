package main

import (
	"github.com/ddddddO/work/go/plugin-go-package/host/contract"
)

func Sentence(name string) string {
	return "Protocol description: " + name
}

func Name() string {
	return "dhcp"
}

func Port() int {
	return 68
}

func Test(p contract.Protocol) contract.Protocol {
	p.Www = 22222
	return p
}
