package main

import "fmt"

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}

func wiggleMaxLength(nums []int) int {
	res := 0
	if len(nums) == 1 || (len(nums) == 2 && nums[0] != nums[1]) {
		return len(nums)
	}
	temp := make([]int, len(nums))
	temp[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		temp[i+1] = nums[i] - nums[i+1]
	}
	temp = append(temp, 0)
	fmt.Println(temp)
	for i := 1; i < len(temp)-1; i++ {
		if temp[i] == 0 {
			continue
		}

		if temp[i] < temp[i+1] && temp[i] < temp[i-1] {
			fmt.Println("////", temp[i])
			res++
		} else if temp[i] > temp[i+1] && temp[i] > temp[i-1] {
			fmt.Println("=-=-=", temp[i])
			res++
		} else {
			fmt.Println(temp[i])
			temp = append(temp[:i], temp[i+1:]...)
		}
	}
	fmt.Println(temp)

	return res + 1
}
