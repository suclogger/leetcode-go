package main

import "math"

func isBalanced(root *TreeNode) bool {
	return balance(root) != -1
}

func balance(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := balance(root.Left)
	if l == -1 {
		return -1
	}
	r := balance(root.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(float64(l-r)) > 1 {
		return -1
	} else {
		return int(math.Max(float64(l), float64(r))) + 1
	}
}
