package main

import "fmt"

type Account struct {
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) WithDraw(amount float64) {
	a.Balance -= amount
}

func main() {
	a := &Account{
		Name:    "Por",
		Balance: 100.00,
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			go a.WithDraw(100)
		} else {
			go a.Deposit(100)
		}
	}

	fmt.Println(a.Balance)
}
