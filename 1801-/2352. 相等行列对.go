package main

import "strconv"

func main() {

}

func equalPairs(grid [][]int) int {
	res := make(map[string]int)
	temp := make([]string, len(grid))
	for i := 0; i < len(grid); i++ {
		r := ""
		for j := 0; j < len(grid); j++ {
			r += strconv.Itoa(grid[i][j])
			if j == len(grid)-1 {
				res[r]++
			}
			temp[j] += strconv.Itoa(grid[i][j])
		}
	}
	nums := 0
	for _, s := range temp {
		if n, ok := res[s]; ok {
			nums += n + 1
		}
	}
	return nums
}
