package main

import "fmt"

type Runner interface {
	Run()
}

type User struct {
	Name string
}

func (u *User) Run() {
	fmt.Printf("%s run 3 kilos\n", u.Name)
}

type Num int

func (n *Num) Run() {
	*n += 1
}

type Account struct {
	Name    string
	Balance float64
}

func (a Account) String() string {
	return fmt.Sprintf("%s account has balance: %.2f", a.Name, a.Balance)
}

func main() {
	var r Runner

	// r = User{
	// 	Name: "Por",
	// }

	r = &User{
		Name: "Por",
	}

	r.Run()

	var n Num
	r = &n

	r.Run()

	fmt.Println(n)

	// u := r.(*User)

	// n = r.(*Num)

	// u, ok := r.(*User)

	// n, ok = r.(*Num)

	do(21)
	do("hello")
	do(true)

	a := Account{
		Name:    "Por",
		Balance: 100.00,
	}
	fmt.Println(a)
	pa := &a
	fmt.Println(pa)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
