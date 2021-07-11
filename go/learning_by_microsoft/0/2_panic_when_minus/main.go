package main

import (
	"fmt"
)

func main() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		// before code
		// if val < 0 {
		// 	panic("minus entry!")
		// }
		// if val == 0 {
		// 	fmt.Print("0 is neither negative nor positive")
		// 	continue
		// }

		// fmt.Println("You entered:", val)

		// solution code
		switch {
		case val < 0:
			panic("You entered a negative number!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}
