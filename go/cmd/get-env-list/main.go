package main

import (
	"fmt"
	"os"
)

// DB_PASSWORD=xxxxxxxxxxx go run main.go
func main() {
	envs := os.Environ()
	fmt.Println(envs)
}
