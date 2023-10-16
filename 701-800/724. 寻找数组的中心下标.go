package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 1, 1, 1}))
}

func pivotIndex(nums []int) int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	left[0] = nums[0]
	right[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i] + left[i-1]
		right[len(nums)-1-i] = nums[len(nums)-1-i] + right[len(nums)-i]
	}
	fmt.Println(left)
	fmt.Println(right)
	for i := 0; i < len(nums); i++ {
		if left[i] == right[i] {
			return i
		}
	}
	return -1
}
