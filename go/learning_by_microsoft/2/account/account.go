package account

import "fmt"

type Account struct {
	firstName string // 名前
	lastName  string // 苗字
}

func (a *Account) ChangeName(firstName, lastName string) {
	a.firstName = firstName
	a.lastName = lastName
}

func (a *Account) String() string {
	return fmt.Sprintf("%s - %s", a.lastName, a.firstName)
}

type Employee struct {
	*Account
	creditNumber int
}

func (e *Employee) AddCredits(add int) {
	e.creditNumber += add
}

func (e *Employee) RemoveCredits(minus int) {
	e.creditNumber -= minus
}

func (e *Employee) CheckCredits() int {
	return e.creditNumber
}
