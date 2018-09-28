package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	u := User{
		"Weerasak",
		"Chongnguluam",
		33,
	}

	fmt.Printf("%+v\n", u)

	u = User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Age:       33,
	}

	fmt.Printf("%+v\n", u)
}
