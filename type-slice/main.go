package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	s2 := nums[1:]
	fmt.Println(s2, len(s2))
	fmt.Println(nums, len(nums))
	// slice
	//var nums []int
	// nums := []int{1, 2, 3, 4}
	// nums := make([]int, 0, 4)
	// fmt.Println(nums)

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)

	// fmt.Println(nums)

	// fmt.Println(Sum(nums))
	// fmt.Println(Product(nums))
}

func Sum(nums []int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	return total
}
func Product(nums []int) int {
	total := 1
	for _, n := range nums {
		total += n
	}
	return total
}
