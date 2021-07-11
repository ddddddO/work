package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}

func TestTransfer(t *testing.T) {
	src := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 1000,
	}

	dest := Account{
		Customer: Customer{
			Name:    "Alice",
			Address: "Japan, Tokyo",
			Phone:   "(043) xxx xxxx",
		},
		Number:  1002,
		Balance: 3000,
	}

	amountJohnToAlice := float64(500)
	src.Transfer(&dest, amountJohnToAlice)

	if src.Balance != 500 {
		t.Error("unexpected")
	}
	if dest.Balance != 3500 {
		t.Error("unexpected")
	}

}
