package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/ddddddO/work/go/plugin-go-package/host/contract"
)

func main() {
	log.Println("plugin-go-package")

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type CompiledPlugin struct {
	Sentence func(name string) string
	Name     func() string
	Port     func() int
	Test     func(p contract.Protocol) contract.Protocol
}

const pluginDir = "../plugin"

func run() error {
	compiledPlugins := []CompiledPlugin{}

	if err := filepath.Walk(pluginDir, func(pluginPath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ".so" {
			return nil
		}

		p, err := plugin.Open(pluginPath)
		if err != nil {
			return err
		}

		s, err := p.Lookup("Sentence")
		if err != nil {
			return err
		}
		sentenceFunc, ok := s.(func(string) string)
		if !ok {
			return fmt.Errorf("required Sentence function")
		}

		n, err := p.Lookup("Name")
		if err != nil {
			return err
		}
		nameFunc, ok := n.(func() string)
		if !ok {
			return fmt.Errorf("required Name function")
		}

		port, err := p.Lookup("Port")
		if err != nil {
			return err
		}
		portFunc, ok := port.(func() int)
		if !ok {
			return fmt.Errorf("required Port function")
		}

		test, err := p.Lookup("Test")
		if err != nil {
			return err
		}
		testFunc, ok := test.(func(contract.Protocol) contract.Protocol)
		if !ok {
			return fmt.Errorf("required Test function")
		}

		compiledPlugins = append(compiledPlugins, CompiledPlugin{
			Sentence: sentenceFunc,
			Name:     nameFunc,
			Port:     portFunc,
			Test:     testFunc,
		})

		return nil
	}); err != nil {
		return err
	}

	for _, proto := range compiledPlugins {
		fmt.Println("---------------------------------")
		fmt.Printf("Sentence: %s\n", proto.Sentence(proto.Name()))
		fmt.Printf("\tName: %s\n", proto.Name())
		fmt.Printf("\tPort: %d\n", proto.Port())

		p := contract.Protocol{
			Xxx: proto.Name(),
			Www: proto.Port(),
		}
		testP := proto.Test(p)
		fmt.Printf("\tTest.contract.Protocol.Xxx: %s\n", testP.Xxx)
		fmt.Printf("\tTest.contract.Protocol.Www: %d\n", testP.Www)
	}
	// Output:
	// 2025/01/01 12:52:30 plugin-go-package
	// ---------------------------------
	// Sentence: Protocol description: dhcp
	// 		Name: dhcp
	// 		Port: 68
	// 		Test.contract.Protocol.Xxx: dhcp
	// 		Test.contract.Protocol.Www: 22222
	// ---------------------------------
	// Sentence: Protocol description: imap
	// 		Name: imap
	// 		Port: 143
	// 		Test.contract.Protocol.Xxx: converted!
	// 		Test.contract.Protocol.Www: 11111

	return nil
}
