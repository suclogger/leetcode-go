package daily

import "math"

func imageSmoother(img [][]int) [][]int {
	if len(img) == 0 || len(img[0]) == 0 {
		return img
	}
	m := len(img)
	n := len(img[0])
	var res [][]int

	for i := 0; i < m; i++ {
		var r []int
		for j := 0; j < n; j++ {
			s := img[i][j]
			c := 1
			if i > 0 {
				s += img[i-1][j]
				c += 1
			}
			if j > 0 {
				s += img[i][j-1]
				c += 1
			}
			if i > 0 && j > 0 {
				s += img[i-1][j-1]
				c += 1
			}
			if i > 0 && j < n-1 {
				s += img[i-1][j+1]
				c += 1
			}
			if i < m-1 {
				s += img[i+1][j]
				c += 1
			}
			if j < n-1 {
				s += img[i][j+1]
				c += 1
			}
			if i < m-1 && j < n-1 {
				s += img[i+1][j+1]
				c += 1
			}
			if i < m-1 && j > 0 {
				s += img[i+1][j-1]
				c += 1
			}
			r = append(r, int(math.Floor(float64(s)/float64(c))))
		}
		res = append(res, r)
	}
	return res
}
