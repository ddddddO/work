package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	script := "curl -v https://google.com"
	cmd := exec.Command("/bin/sh", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Start error!")
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Wait error!")
	}
}
