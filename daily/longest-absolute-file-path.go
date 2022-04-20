package main

func lengthLongestPath(input string) int {
	input = string(append([]byte(input), '\n'))
	stack := make([]string, 0)
	l := 0
	is_file := false
	cur := ""
	ans := 0
	d := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' || input[i] == '"' {

			if is_file {
				if ans < l+len(cur)+d {
					ans = l + len(cur) + d
				}
				is_file = false
			} else {
				stack = append(stack, cur)
				l += len(cur)

			}
			cur = ""
			d = 0
		} else if input[i] == '\t' {
			d += 1
		} else {
			for len(stack) != d && len(stack) > 0 {
				l -= len(stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			cur += string(input[i])
			if input[i] == '.' {
				is_file = true
			}
		}

	}
	return ans

}
