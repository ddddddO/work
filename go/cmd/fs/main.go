package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	normal()
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

	fmt.Println()

	filesystemToMarkdown()
	// Output:
	//- .
	// 	- find_pipe_programmable-gtree
	// 					- main.go
	// 	- go-list_pipe_programmable-gtree
	// 					- main.go
	// 	- like_cli
	// 					- adapter
	// 									- executor.go
	// 									- indentation.go
	// 					- main.go
	// 	- noexist
	// 					- xxx
	// 	- programmable
	// 					- main.go
}

func normal() {
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
}

func filesystemToMarkdown() {
	const indent = "	"
	fileSystem := os.DirFS("../../../../gtree/example")
	startPoint := "."
	err := fs.WalkDir(fileSystem, startPoint, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		names := strings.Split(path, "/")

		line := fmt.Sprintf("%s- %s", strings.Repeat(indent, strings.Count(path, "/")+1), names[len(names)-1])
		if names[len(names)-1] == startPoint {
			line = fmt.Sprintf("%s- %s", strings.Repeat(indent, strings.Count(path, "/")), names[len(names)-1])
		}

		fmt.Println(line)
		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
