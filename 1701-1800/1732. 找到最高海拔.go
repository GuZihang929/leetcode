package main

import "fmt"

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
}

func largestAltitude(gain []int) int {
	res := 0
	temp := 0
	for i := 0; i < len(gain); i++ {
		temp = gain[i] + temp
		fmt.Println(temp)
		if res < temp {
			res = temp
		}
	}
	return res
}
