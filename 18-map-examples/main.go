package main

import (
	"fmt"
	"sort"
)

func main() {
	var ages map[string]int     // ages == nil
	ages = make(map[string]int) // empty map
	ages["alice"] = 31
	ages["charlie"] = 34

	fmt.Println(ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	if age, ok := ages["bob"]; !ok {
		fmt.Println("Bob age:", age)
	}

	delete(ages, "alice")

	fmt.Println(ages)
}

func equal(x, y map[string]int) bool {
	// exercise
	return false
}
