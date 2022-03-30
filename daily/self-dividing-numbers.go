package main

import "strconv"

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		zichu := true
		for _, r := range strconv.Itoa(i) {
			if int(r-'0') == 0 || i%(int(r-'0')) != 0 {
				zichu = false
				break
			}
		}
		if zichu {
			res = append(res, i)
		}

	}
	return res
}
