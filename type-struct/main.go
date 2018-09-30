package main

import (
	"fmt"
	"sort"
	"strings"
)

type User struct {
	FirstName     string
	LastName      string
	Age           int
	FavoriteFoods []string
}

func main() {
	u := User{
		FirstName: "Weerasak",
		LastName:  "Chongnguluam",
		Age:       33,
	}

	fmt.Printf("%#v\n", u)
	fmt.Println("FirstName", u.FirstName)
	fmt.Println("LastName", u.LastName)
	fmt.Println("Age", u.Age)

	u2 := u
	u2.FirstName = "Kanokon"
	u2.Age = 27

	fmt.Printf("%#v\n", u)
	fmt.Println("FirstName", u.FirstName)
	fmt.Println("LastName", u.LastName)
	fmt.Println("Age", u.Age)

	var users []User
	users = append(users, User{
		FirstName:     "Weerasak",
		LastName:      "Chongnguluam",
		Age:           33,
		FavoriteFoods: []string{"Noodle", "Sushi"},
	})

	users = append(users, User{
		FirstName:     "Weerasak",
		LastName:      "ABCDEF",
		Age:           27,
		FavoriteFoods: []string{"Somtam", "Chicken grill"},
	})

	sort.Slice(users, func(i, j int) bool {
		if users[i].FirstName != users[j].FirstName {
			return users[i].FirstName < users[j].FirstName
		}

		return users[i].LastName < users[j].LastName
	})

	for _, u := range users {
		Display(u)
		fmt.Println()
	}
}

func Display(u User) {
	fmt.Printf("Name: %s %s\n", u.FirstName, u.LastName)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Favorite Foods: %s\n", strings.Join(u.FavoriteFoods, ", "))
	/*
		Name: Weerasak Chongnguluam
		Age: 33
		Favorite Foods: Noodle, Shushi
	*/
}
