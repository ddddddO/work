package main

import (
	"fmt"
)

func main() {
	hello := "Hello, world!"

	fmt.Println(Black(hello))
	fmt.Println(Red(hello))
	fmt.Println(Green(hello))
	fmt.Println(Yellow(hello))
	fmt.Println(Blue(hello))
	fmt.Println(Magenta(hello))
	fmt.Println(Cyan(hello))
	fmt.Println(White(hello))
}

// https://github.com/uber-go/zap/blob/master/internal/color/color.go
type color uint8

const (
	black color = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func Black(s string) string {
	return colorize(black, s)
}

func Red(s string) string {
	return colorize(red, s)
}

func Green(s string) string {
	return colorize(green, s)
}

func Yellow(s string) string {
	return colorize(yellow, s)
}

func Blue(s string) string {
	return colorize(blue, s)
}

func Magenta(s string) string {
	return colorize(magenta, s)
}

func Cyan(s string) string {
	return colorize(cyan, s)
}

func White(s string) string {
	return colorize(white, s)
}

const format = "\x1b[%dm%s\x1b[0m"

func colorize(color color, s string) string {
	return fmt.Sprintf(format, color, s)
}
