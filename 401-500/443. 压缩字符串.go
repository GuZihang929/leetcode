package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{
		'a',
	}))
}

func compress(chars []byte) int {

	if len(chars) == 1 {
		return 1
	}
	temp := chars[0]
	num := 1
	j := 0
	for i := 1; i < len(chars); i++ {
		if temp == chars[i] {
			num++
		} else {
			chars[j] = temp
			j++
			btye := re(num)

			for a := 0; a < len(btye); a++ {
				if len(btye) == 1 && btye[0] == '1' {

					break
				}

				chars[j] = btye[a]
				j++

			}
			temp = chars[i]
			num = 1
		}

		if i == len(chars)-1 {
			chars[j] = temp
			j++
			btye := re(num)
			for a := 0; a < len(btye); a++ {
				if len(btye) == 1 && btye[0] == '1' {

					break
				}

				chars[j] = btye[a]
				j++

			}
		}
	}

	return j
}

func re(a int) []byte {
	ok := a / 10
	temp := ""
	for ok != 0 {
		temp = strconv.Itoa(a%10) + temp
		a = a / 10
		ok = a / 10
	}
	temp = strconv.Itoa(a%10) + temp
	return []byte(temp)
}
