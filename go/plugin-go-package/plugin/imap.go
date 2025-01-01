package main

import (
	"github.com/ddddddO/work/go/plugin-go-package/host/contract"
)

func Sentence(name string) string {
	return "Protocol description: " + name
}

func Name() string {
	return "imap"
}

func Port() int {
	return 143
}

func Test(p contract.Protocol) contract.Protocol {
	p.Xxx = "converted!"
	p.Www = 11111
	return p
}
