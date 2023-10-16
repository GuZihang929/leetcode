package main

func main() {

}

func isSubsequence(s string, t string) bool {
	ps, pt := 0, 0
	semp := []byte(s)
	temp := []byte(t)

	for ps < len(temp) && ps < len(semp) && pt < len(temp) {
		if temp[pt] == semp[ps] {
			pt++
			ps++
		} else {
			pt++
		}
	}
	if ps == len(semp) {
		return true
	}
	return false
}
