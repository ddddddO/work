package bank

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// 送金
func (a *Account) Transfer(dest *Account, amount float64) error {
	if err := a.Withdraw(amount); err != nil {
		return err
	}
	if err := dest.Deposit(amount); err != nil {
		return err
	}

	return nil
}

type Statementer interface {
	Statement() string
}

func Statement(Statementer Statementer) string {
	return Statementer.Statement()
}

type JSONAccount struct {
	*Account
}

func (ja *JSONAccount) Statement() string {
	d, err := json.Marshal(ja)
	if err != nil {
		panic(err)
	}

	return string(d)
}
