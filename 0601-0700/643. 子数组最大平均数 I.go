package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{0, 1, 1, 3, 3}, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	var res float64
	temp := 0
	for i := 0; i < k; i++ {
		temp += nums[i]
	}
	res = float64(temp) / float64(k)
	for i := 1; i < len(nums)-k+1; i++ {
		temp -= nums[i-1]
		temp += nums[k+i-1]
		if res < float64(temp)/float64(k) {
			res = float64(temp) / float64(k)
		}
	}
	return res
}
