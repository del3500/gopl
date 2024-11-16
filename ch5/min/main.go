package main

import "fmt"

func main() {
	nums := []int{2, 2, 3, 4, 5}
	fmt.Println(min(nums...))
	fmt.Println(min())
	fmt.Println(min(5))
	fmt.Println(min(4, 5, 6))
}
func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	lowest := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < lowest {
			lowest = nums[i]
		}
	}
	return lowest
}
