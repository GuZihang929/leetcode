package main

import "math"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := increasingTriplet(nums)
	println(result)
}

func increasingTriplet(nums []int) bool {
	a := math.MaxInt32
	b := math.MaxInt32

	for _, e := range nums {
		if e <= a {
			a = e
		} else if e <= b {
			b = e
		} else {
			return true
		}
	}
	return false
}
