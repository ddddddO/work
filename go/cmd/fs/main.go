package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fileSystem := os.DirFS("../../../../gtree/example")
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	// Output:
	// .
	// find_pipe_programmable-gtree
	// find_pipe_programmable-gtree/main.go
	// go-list_pipe_programmable-gtree
	// go-list_pipe_programmable-gtree/main.go
	// like_cli
	// like_cli/adapter
	// like_cli/adapter/executor.go
	// like_cli/adapter/indentation.go
	// like_cli/main.go
	// programmable
	// programmable/main.go
}
