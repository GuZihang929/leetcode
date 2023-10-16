package main

import "sort"

func main() {

}

func maxOperations(nums []int, k int) int {
	left, right := 0, len(nums)-1
	res := 0
	sort.Ints(nums)

	for left < right {
		if nums[left]+nums[right] == k {
			res++
			right--
			left++
		} else if nums[left]+nums[right] > k {
			right--
		} else {
			left++
		}
	}
	return res
}
