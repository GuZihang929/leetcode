package main

import "fmt"

func main() {
	canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {}})
}

func canVisitAllRooms(rooms [][]int) bool {
	dfs841(0, rooms)
	fmt.Println(rooms)

	for _, room := range rooms {
		if len(room) == 0 || room[0] != -1 {
			return false
		}
	}
	return true
}

func dfs841(index int, rooms [][]int) {
	if len(rooms[index]) == 0 {
		rooms[index] = append(rooms[index], -1)
		return
	}

	if rooms[index][0] == -1 {
		return
	}
	temp := rooms[index]
	rooms[index] = []int{-1}

	for _, v := range temp {
		dfs841(v, rooms)
	}

}
