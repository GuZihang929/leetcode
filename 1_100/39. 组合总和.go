package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}

//给你一个 无重复元素 的整数数组candidates 和一个目标整数target，找出candidates中可以使数字和为目标数target 的 所有不同组合 ，
//并以列表形式返回。你可以按 任意顺序 返回这些组合。

//candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。

//对于给定的输入，保证和为target 的不同组合数少于 150 个。

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	sort.Ints(candidates)
	Arr = candidates
	res := &[][]int{}
	ResMap = make(map[string][]int)
	dfs39([]int{}, res, target, 0)

	return *res
}

var Arr []int
var ResMap map[string][]int

func dfs39(temp []int, res *[][]int, target int, index int) {

	if target <= 0 {

		if target == 0 {
			b := make([]int, len(temp))
			sort.Ints(temp)
			copy(b, temp)
			s := ""
			for _, i := range temp {
				s += string(i)
			}

			if _, ok := ResMap[s]; !ok {
				ResMap[s] = b
				*res = append(*res, b)
			}
			//*res = append(*res, temp)
			return
		}
	}

	for i := index; i < len(Arr); i++ {
		if target < Arr[i] {
			break
		}
		temp = append(temp, Arr[i])
		dfs39(temp, res, target-Arr[i], i)
		temp = temp[:len(temp)-1]

	}

}
