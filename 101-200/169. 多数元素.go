package main

import "sort"

func main() {

}

func majorityElement(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	sort.Ints(nums)

	slow, fast := 0, 1

	for fast-slow < len(nums)/2 {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			slow = fast
			fast++
		}
	}
	return nums[fast]

}
