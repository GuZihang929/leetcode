package main

import (
	"fmt"
)

func main() {
	var n int
	var s, t string
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Scan(&t)

	res := 0
	add1, add2 := false, false
	smap := make(map[byte]int)
	tmap := make(map[byte]int)
	for i := 0; i < n; i++ {
		if s[i] == t[i] {
			res++
			continue
		}
		smap[s[i]] = i
		tmap[t[i]] = i
		if v, ok := smap[t[i]]; ok {
			add1 = true
			if t[v] == s[i] {
				add2 = true
			}
		}
	}
	if add1 {
		res++
	}

	if add2 {
		res++
	}
	fmt.Println(res)
}
