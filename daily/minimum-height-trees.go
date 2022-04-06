package main

func findMinHeightTrees(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		x := e[0]
		y := e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	parent := make([]int, n)

	dfs := func(x int) (y int) {
		visited := make([]bool, n)
		visited[x] = true
		q := []int{x}
		for len(q) > 0 {
			y, q = q[0], q[1:]
			for _, z := range g[y] {
				if !visited[z] {
					visited[z] = true
					parent[z] = y
					q = append(q, z)
				}
			}
		}
		return
	}
	x := dfs(0)
	y := dfs(x)

	var path []int
	parent[x] = -1
	for y != -1 {
		path = append(path, y)
		y = parent[y]
	}

	if len(path)%2 == 0 {
		return []int{path[len(path)/2-1], path[len(path)/2]}
	} else {
		return []int{path[len(path)/2]}
	}
}
