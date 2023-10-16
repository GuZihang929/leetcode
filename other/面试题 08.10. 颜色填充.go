package main

func main() {
	floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	direction := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	if image[sr][sc] != newColor {
		dfs08(image[sr][sc], image, sr, sc, newColor, direction)

	}
	return image
}

func dfs08(val int, image [][]int, sr, sc int, newColor int, dir [][]int) {

	if val == image[sr][sc] {
		image[sr][sc] = newColor
		for i := 0; i < 4; i++ {
			if sr+dir[i][0] >= 0 && sr+dir[i][0] < len(image) && sc+dir[i][1] >= 0 && sc+dir[i][1] < len(image[0]) {
				dfs08(val, image, sr+dir[i][0], sc+dir[i][1], newColor, dir)

			}
		}
	}

}
