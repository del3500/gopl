package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 99}
	fmt.Println(max(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(max())
	fmt.Println(max(7))
	fmt.Println(max(x...))
}

func max(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums == nil {
		return 0
	}
	highest := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > highest {
			highest = nums[i]
		}
	}
	return highest
}
