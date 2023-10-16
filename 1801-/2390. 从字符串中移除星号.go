package main

func main() {

}

func removeStars(s string) string {

	a := make([]byte, len(s))
	p := 0
	temp := []byte(s)

	for i := 0; i < len(s); i++ {
		if temp[i] == '*' {
			if p == 0 {
				continue
			}
			p--
			continue
		}
		a[p] = temp[i]
		p++
	}
	return string(a[:p])
}
