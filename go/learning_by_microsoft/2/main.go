package main

import (
	"fmt"

	ac "github.com/ddddddO/work/go/learning_by_microsoft/2/account"
)

func main() {
	const (
		firstName = "tarou"
		lastName  = "tanaka"
	)
	account := &ac.Account{}
	account.ChangeName(firstName, lastName)
	employee := &ac.Employee{
		Account: account,
	}
	fmt.Println(employee)

	employee.ChangeName("hanako", "yamada")
	fmt.Println(employee)

	add := 10000
	employee.AddCredits(add)

	minus := 500
	employee.RemoveCredits(minus)

	fmt.Println(employee.CheckCredits())
}
