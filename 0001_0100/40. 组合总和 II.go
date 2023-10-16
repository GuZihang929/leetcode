package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	Arr40 = candidates
	res := &[][]int{}
	dfs40([]int{}, res, 0, target)
	return *res
}

var Arr40 []int

func dfs40(temp []int, res *[][]int, index int, target int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(temp))
			copy(b, temp)
			*res = append(*res, b)
		}
		return
	}

	for i := index; i < len(Arr40); i++ {
		if i > index && Arr40[i] == Arr40[i-1] {
			continue
		}

		if target >= Arr40[i] {
			temp = append(temp, Arr40[i])
			dfs40(temp, res, i+1, target-Arr40[i])
			temp = temp[:len(temp)-1]
		}
	}

}
