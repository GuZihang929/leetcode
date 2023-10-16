package main

func main() {

}

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	p := 0
	temp := asteroids[0]
	for i := 1; i < len(asteroids); i++ {
		if (temp > 0 && asteroids[i] > 0) || (temp < 0 && asteroids[i] < 0) {
			temp = asteroids[i]
		} else if 0 == temp+asteroids[i] {

			temp = res[p]
		} else {
			if temp < -asteroids[i] {
				for j := len(res) - 1; j > -1; j-- {
					if res[j] < -asteroids[i] {
						p--
					} else if res[j]+asteroids[i] >= 0 {
						temp = res[p]
						break
					}
				}
			}
		}
	}
	return res[:p]

}
