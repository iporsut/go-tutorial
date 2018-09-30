package user

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func (u *User) Name() string {
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
