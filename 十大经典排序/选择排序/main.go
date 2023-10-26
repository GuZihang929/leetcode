package main

import "fmt"

func main() {
	fmt.Println(SelectionSort([]int{4, 3, 12, 5, 546, 123, 53, 15, 61, 56, 7, 2}))

}

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		max := 0
		for j := 1; j < len(arr)-i; j++ {
			if arr[max] < arr[j] {
				max = j
			}
		}
		arr[len(arr)-i-1], arr[max] = arr[max], arr[len(arr)-i-1]
	}
	return arr
}
