package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(closeStrings("a", "aa"))
}

func closeStrings(word1 string, word2 string) bool {
	w1 := make(map[int]int)
	w2 := make(map[int]int)
	for _, i := range word1 {
		w1[int(i)]++
	}

	for _, i := range word2 {

		if _, ok := w1[int(i)]; !ok {
			return false
		}
		w2[int(i)]++
	}
	if len(w1) != len(w2) {
		return false
	}
	fmt.Println(w1)
	fmt.Println(w2)

	l1 := make([]int, len(w1))
	l2 := make([]int, len(w2))
	a := 0
	for _, i2 := range w1 {
		l1[a] = i2
		a++
	}

	b := 0
	for _, i2 := range w2 {
		l2[b] = i2
		b++
	}

	sort.Ints(l1)
	sort.Ints(l2)
	for i := 0; i < len(l1); i++ {
		if l2[i] != l1[i] {
			return false
		}
	}
	fmt.Println(l1)
	fmt.Println(l2)
	return true

}
