package main

func FindKDistantIndices(nums []int, key int, k int) []int {
	var keyIdx []int
	for idx, num := range nums {
		if key == num {
			keyIdx = append(keyIdx, idx)
		}
	}
	lastVisit := 0
	var res []int
	for _, idx := range keyIdx {
		var from int
		if idx-k > lastVisit {
			from = idx - k
		} else {
			from = lastVisit
		}
		var to int
		if idx+k > len(nums)-1 {
			to = len(nums) - 1
		} else {
			to = idx + k
		}
		lastVisit = to + 1

		for val := from; val <= to; val++ {
			res = append(res, val)
			println(val)
		}
	}
	return res
}

//func main() {
//	FindKDistantIndices([]int{2,2,2,2,2}, 2, 2)
//}
