package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{4, 3, 12, 5, 124, 465, 123, 53, 675, 12325, 15, 61, 56, 234, 2}))

}

func QuickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	temp := arr[0]
	var left, right []int
	for i := 1; i < len(arr); i++ {
		if arr[i] < temp {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}

	}

	l := QuickSort(left)
	r := QuickSort(right)
	l = append(l, temp)
	return append(l, r...)

}
