package main

import (
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
}

func (u User) Name() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

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
	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
	}

	fmt.Println(u.Name())

	pu := &User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
	}

	fmt.Println(pu.Name())

	a := Account{
		Name:    "Weerasak",
		Balance: 1000.00,
	}
	fmt.Printf("Account: %s\n", a.Name)
	fmt.Println(a.Balance)
	a.Deposit(1000)
	fmt.Println(a.Balance)
	a.WithDraw(500)
	fmt.Println(a.Balance)

	pa := &Account{
		Name:    "Kanokorn",
		Balance: 10000.00,
	}
	fmt.Printf("Account: %s\n", pa.Name)
	fmt.Println(pa.Balance)
	pa.Deposit(1000)
	fmt.Println(pa.Balance)
	pa.WithDraw(500)
	fmt.Println(pa.Balance)
}
