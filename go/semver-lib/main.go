package main

import (
	"fmt"

	"github.com/Songmu/gitsemvers"
)

func main() {
	const path = "../../../../ddddddO/gtree"
	sv := &gitsemvers.Semvers{RepoPath: path}
	semvers := sv.VersionStrings()
	for _, v := range semvers {
		fmt.Println(v)
	}
}