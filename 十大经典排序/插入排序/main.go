package main

import "fmt"

func main() {
	fmt.Println(InsertionSort([]int{4, 3, 12, 5, 546, 123, 53, 15, 61, 56, 7, 2}))

}

func InsertionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		j := i
		for j >= 0 && arr[j+1] < arr[j] {
			arr[j+1], arr[j] = arr[j], arr[j+1]
			j--
		}
	}

	return arr
}
