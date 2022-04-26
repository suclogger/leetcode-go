package main

func projectionArea(grid [][]int) int {
	z := 0
	if len(grid) == 0 {
		return z
	}
	x_max := make([]int, len(grid[0]))
	for _, j := range grid {
		tmpz := 0
		tmpx := 0
		for ii, jj := range j {
			if jj > 0 {
				tmpz += 1
			}
			if x_max[ii] < jj {
				x_max[ii] = jj
			}
			if jj > tmpx {
				tmpx = jj
			}
		}
		z += tmpz
		z += tmpx
	}
	for _, t := range x_max {
		z += t
	}
	return z
}
