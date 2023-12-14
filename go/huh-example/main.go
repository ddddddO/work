package main

import (
	"log"

	"github.com/charmbracelet/huh"
)

func main() {
	confirm := false
	err := huh.NewConfirm().
		Title("Are you sure?").
		Affirmative("Yes!").
		Negative("No.").
		Value(&confirm).
		Run()
	if err != nil {
		log.Fatal(err)
	}
}
