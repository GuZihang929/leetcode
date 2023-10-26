package main

import "fmt"

func main() {

	fmt.Println(MaoPao([]int{4, 3, 12, 5, 546, 123, 53, 15, 61, 56, 7, 2}))
}

func MaoPao(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}
