package main

import "fmt"

func main() {
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	fmt.Println(months)

	q2 := months[4:7]
	summer := months[2:6]
	fmt.Println(q2)
	fmt.Println(summer)

	for _, s := range summer {
		fmt.Println(s)
	}

	fmt.Println(summer[:20])
	summer = summer[:5]
	fmt.Println(summer)

	ns := []int{0, 1, 2, 3, 4, 5}
	reverse(ns)
	fmt.Println(ns)

	ns2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(equal(ns, ns2))

	sl := make([]int, 10)
	fmt.Println(sl)

	sl2 := make([]int, 0, 10)
	fmt.Println(sl2)

	var runes []rune
	for _, r := range "Hello, กรุงเทพ" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

func reverse(s []int) {
	// excercise
}

func equal(x, y []string) bool {
	// excercise
	return false
}
