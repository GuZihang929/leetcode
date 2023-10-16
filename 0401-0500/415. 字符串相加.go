package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(addStrings("77", "456"))
}

func addStrings(num1 string, num2 string) string {

	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	res := ""
	sss := 0

	for i, i2 := range num2 {
		n2, err := strconv.Atoi(string(i2))
		if err != nil {
			return ""
		}
		n1, err := strconv.Atoi(string(num1[len(num1)-i-1]))
		if err != nil {
			return ""
		}
		temp := (n1 + n2 + sss) % 10
		if n1+n2+sss > 9 {
			sss = 1
		} else {
			sss = 0
		}
		res = strconv.Itoa(temp) + res
	}
	if sss == 0 {

		res = num1[:len(num1)-len(num2)] + res
	} else {

		n1, err := strconv.Atoi(string(num1[len(num1)-len(num2)]))
		if err != nil {
			return ""
		}
		res = strconv.Itoa(n1) + res
		res = num1[:len(num1)-len(num2)] + res
	}
	return res
}
