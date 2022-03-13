package main

func maximumTop(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if k%2 == 1 {
			return -1
		} else {
			return nums[0]
		}
	}
	if k == 0 {
		return nums[0]
	}
	//    1  2  3-   1  1  1  1
	// 5  1- 2- 3- 3+ 3-
	// 6  1- 2- 3- 3+ 3- 3+
	curMax := 0
	for i := 0; i <= k && i < len(nums); i++ {
		if nums[i] > curMax && ((k-i)%2 == 0 || k-i > 1) {
			curMax = nums[i]
		}
	}
	return curMax
}

func main() {
	println(maximumTop([]int{94, 23, 58, 92, 3, 63, 68, 43, 8, 97, 54, 11, 63, 55, 73, 38, 22, 80, 45, 43, 25, 34, 67, 76, 80, 9, 30, 27, 50, 7, 57, 63, 63, 27, 46, 1, 50, 55, 99, 92, 73, 9, 57, 25, 59, 54, 100, 56, 64, 94, 99},
		79))
}
