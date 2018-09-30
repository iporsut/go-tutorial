package main

import "fmt"

func main() {
	// nums := [4]int{1, 2, 3, 4}
	nums := [4]int{2: 3, 3: 4}
	fmt.Println(len(nums))
	fmt.Println(nums[3])

	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	for _, v := range nums {
		fmt.Println(v)
	}
}

func Sum(nums [4]int) int

func Product(nums [4]int) int
