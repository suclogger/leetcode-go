package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	if len(answerKey) == 0 {
		return 0
	}
	runes := []rune(answerKey)
	t_count := maxWithKey(runes, k, 'T')
	f_count := maxWithKey(runes, k, 'F')
	if t_count > f_count {
		return t_count
	}
	return f_count
}

func maxWithKey(answerKey []rune, k int, key rune) int {
	left := 0
	right := 0
	sum := 0
	max := 0
	for {
		if right == len(answerKey) {
			break
		}
		if answerKey[right] != key {
			sum++
		}

		for {
			if sum <= k {
				break
			}
			if answerKey[left] != key {
				left++
				sum--
			} else {
				left++
			}
		}

		if right-left+1 > max {
			max = right - left + 1
		}
		right++
	}
	return max
}
