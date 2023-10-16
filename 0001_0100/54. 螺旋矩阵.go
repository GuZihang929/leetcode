package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}}
	fmt.Println(spiralOrder(arr))
}

func spiralOrder(matrix [][]int) []int {
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := 0, 0
	//四个拐点
	x1, y1 := len(matrix[0])-1, 0
	x2, y2 := len(matrix[0])-1, len(matrix)-1
	x3, y3 := 0, len(matrix)-1
	x4, y4 := 0, 0
	ss := len(matrix[0]) * len(matrix)
	res := make([]int, ss)
	j := 0
	//mly := 0
	for i := 0; i < ss; i++ {
		res[i] = matrix[n][m]
		fmt.Println(res[i])
		n += dir[j][0]
		m += dir[j][1]
		if m > x1 && n == y1 {
			m--
			n++
			j++
			x4++
			y4++

		}
		if m == x2 && n > y2 {
			n--
			m--
			x1--
			y1++
			j++
		}
		if m < x3 && n == y3 {
			m++
			n--
			j++

			x2--
			y2--

		}
		fmt.Println("  =====", m, n)

		fmt.Println("-------", x4, y4)
		if m < x4 && n < y4 {
			m++
			n++
			j = 0
			//mly = 1
			fmt.Println("=-=-=-=")
			x3++
			y3--

		}
	}
	return res
}
