package main

import "strconv"

func minBitFlips(start int, goal int) int {
	r1 := []rune(strconv.FormatInt(int64(start), 2))
	r2 := []rune(strconv.FormatInt(int64(goal), 2))
	res := 0
	if len(r1) > len(r2) {
		diff := len(r1) - len(r2)
		for i := 0; i < diff; i++ {
			if r1[i]-'0' == 1 {
				res++
			}
		}
		for i := 0; i < len(r2); i++ {
			if r1[diff+i] != r2[i] {
				res++
			}
		}
	} else {
		diff := len(r2) - len(r1)
		for i := 0; i < diff; i++ {
			if r2[i] == 1 {
				res++
			}
		}
		for i := 0; i < len(r1); i++ {
			if r2[diff+i] != r1[i] {
				res++
			}
		}

	}
	return res
}
