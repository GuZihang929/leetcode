package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}

func removeDuplicates(a string) string {
	s := []byte(a)
	res := []byte("")
	res = append(res, a[0])
	s = s[1:]
	for _, b := range s {
		if len(res) == 0 {
			res = append(res, b)
		} else {
			if b == res[len(res)-1] {
				res = res[:len(res)-1]
			} else {
				res = append(res, b)
			}
		}
	}

	return string(res)
}
