package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

// echo "Hello" | go run main.go > tmp.png
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	sc := bufio.NewScanner(os.Stdin)
	ret := ""
	for sc.Scan() {
		ret += fmt.Sprintf("%s\n", sc.Text())
	}
	if err := sc.Err(); err != nil {
		return err
	}

	b, err := qrcode.Encode(ret, qrcode.Highest, 256)
	if err != nil {
		return err
	}

	_, err = os.Stdout.Write(b)
	if err != nil {
		return err
	}

	return nil
}
