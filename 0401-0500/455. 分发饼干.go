package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(findContentChildren([]int{0}, []int{1, 1}))
}

func findContentChildren(g []int, s []int) int {
	res := 0

	sort.Ints(g)
	sort.Ints(s)

	for i, j := 0, 0; i < len(g) && j < len(s); i++ {

		if g[i] <= s[j] {
			res++
			j++
		} else {
			i--
			j++
		}

	}

	return res
}
