package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, -1, -5, -4, -5}))
}

func maxProfitII(prices []int) int {
	temp := make([]int, len(prices)-1)

	for i := 0; i < len(prices)-1; i++ {
		temp[i] = prices[i+1] - prices[i]
		if temp[i] < 0 {
			temp[i] = 0
		}
	}
	fmt.Println(temp)
	res := 0
	for _, i2 := range temp {
		res += i2
	}
	return res
}
