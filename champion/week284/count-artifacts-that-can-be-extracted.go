package main

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	// 初始化矿场
	field := make([][]int, n)
	for i := range field {
		field[i] = make([]int, n)
	}

	// 标识挖过了
	for _, d := range dig {
		field[d[0]][d[1]] = 1
	}

	res := 0

	for _, art := range artifacts {
		// 从当前位置往后做DFS，判断是否超出范围即可
		if dfs(art[0], art[1], art[2], art[3], field) {
			res++
		}
	}

	return res
}

// dfs
func dfs(xStart int, yStart int, xMax int, yMax int, field [][]int) bool {
	for i := xStart; i <= xMax; i++ {
		for j := yStart; j <= yMax; j++ {
			if field[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
