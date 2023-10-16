package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(thousandSeparator(0))
}

func thousandSeparator(n int) string {
	res := ""
	j := 0
	temp := strconv.Itoa(n)
	for i := len(temp) - 1; i > -1; i-- {
		if j == 3 {
			res = "." + res
			j = 0
		}
		j++
		res = string(temp[i]) + res
	}
	return res
}
