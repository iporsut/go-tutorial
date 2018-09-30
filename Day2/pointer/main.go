package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	var pn *int
	n := 10
	pn = &n

	fmt.Println(n)

	*pn = 20

	fmt.Println(*pn)

	pn2 := &n

	*pn2++

	fmt.Println(n, *pn, *pn2)

	pn3 := new(int)

	fmt.Println(*pn3)
	*pn3 = 20
	fmt.Println(*pn3)

	u := &User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
	}
	fmt.Println(u)

	p := u
	p.FirstName = "May"
	fmt.Println(u)
}
