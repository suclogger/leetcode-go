package main

import "strings"

func toGoatLatin(sentence string) string {
	sentence += " "
	arr := make([]string, 0)
	cur := ""
	fx := ""
	for _, b := range sentence {
		if len(cur) == 0 && len(fx) == 0 {
			if b != 'a' &&
				b != 'e' &&
				b != 'i' &&
				b != 'o' &&
				b != 'u' &&
				b != 'A' &&
				b != 'E' &&
				b != 'I' &&
				b != 'O' &&
				b != 'U' {
				fx += string(b)
			} else {
				cur += string(b)
			}
			fx += "ma"
			for i := 0; i <= len(arr); i++ {
				fx += "a"
			}
		} else if b == ' ' {
			if len(cur) > 0 || len(fx) > 0 {
				cur += fx
				arr = append(arr, cur)
				cur = ""
				fx = ""
			}
		} else {
			cur += string(b)
		}

	}
	return strings.Join(arr, " ")

}
